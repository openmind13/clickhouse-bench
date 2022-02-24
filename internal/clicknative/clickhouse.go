package clicknative

import (
	"context"
	"fmt"
	"time"

	"clickhouse-bench/internal/config"
	"clickhouse-bench/internal/event"
	"clickhouse-bench/internal/sqlstatement"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

const (
	WRITE_TIMEOUT = 5 * time.Second
)

type Clickhouse struct {
	conn driver.Conn
	ctx  context.Context
}

func NewClickhouse() (*Clickhouse, error) {
	c := &Clickhouse{}
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{config.Config.ClickhouseNativeUrl},
		Auth: clickhouse.Auth{
			Database: config.Config.Database,
			// Username: "default",
			// Password: "default",
		},
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: time.Hour,
		DialTimeout:     5 * time.Second,
		// Debug:           true,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.ctx = context.Background()

	return c, nil
}

func (c *Clickhouse) PrepareDatabase() error {
	if err := c.conn.Exec(c.ctx, fmt.Sprintf(sqlstatement.SQL_DROP_TABLE, config.Config.Table)); err != nil {
		return err
	}
	if err := c.conn.Exec(c.ctx, fmt.Sprintf(sqlstatement.SQL_CREATE_TABLE, config.Config.Table, config.Config.Engine)); err != nil {
		return err
	}
	return nil
}

func (c *Clickhouse) Write(e event.Event) error {
	ctx, cancel := context.WithTimeout(c.ctx, WRITE_TIMEOUT)
	go func() {
		select {
		case <-time.After(WRITE_TIMEOUT):
			cancel()
		case <-ctx.Done():
			return
		}
	}()
	batch, err := c.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", config.Config.Table))
	if err != nil {
		return err
	}
	if err := batch.AppendStruct(&e); err != nil {
		return err
	}
	if err := batch.Send(); err != nil {
		return err
	}
	return nil
}

func (c *Clickhouse) WriteAsync(e event.Event) error {
	ctx, cancel := context.WithTimeout(c.ctx, WRITE_TIMEOUT)
	go func() {
		select {
		case <-time.After(WRITE_TIMEOUT):
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	testStatement := fmt.Sprintf(`INSERT INTO test_metrics.test
		(
			ClientId,
			EventType,
			Data,
			Time,
			ClientConnectionTime,
			SessionDurationSeconds,
			StreamName,
			StreamPluginName,
			Provider,
			Platform,
			Merchant,
			Account,
			ProjectId,
			Fp,
			PlayerVersion,
			Ip,
			RealIp,
			Forward,
			BitRate,
			FramesDecodedRate,
			NackCountRate,
			PacketsLossPercent,
			Rtt,
			DelayStart,
			PreStart,
			Ua,
			BrowserName,
			BrowserVersion,
			OsName,
			OsPlatform,
			OsVersion,
			DeviceType,
			Geohash,
			AccuracyRadius,
			Latitude,
			Longitude,
			Timezone,
			CityNameEn,
			CountryNameEn,
			ContinentNameEn) VALUES (
				%d,
				'%s', 
				'%s', 
				toDateTime('%s', 'Etc/UTC'), 
				toDateTime('%s', 'Etc/UTC'), 
				%d, 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s',
			 	%f, 
				%f, 
				%f, 
				%f, 
				%f, 
				%f, 
				%f, 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				'%s', 
				%d, 
				%f, 
				%f, 
				'%s', 
				'%s', 
				'%s', 
				'%s')`,
		e.ClientId,
		e.EventType,
		e.Data,
		e.Time,
		e.ClientConnectionTime,
		e.SessionDurationSeconds,
		e.StreamName,
		e.StreamPluginName,
		e.Provider,
		e.Platform,
		e.Merchant,
		e.Account,
		e.ProjectId,
		e.Fp,
		e.PlayerVersion,
		e.Ip,
		e.RealIp,
		e.Forward,
		e.BitRate,
		e.FramesDecodedRate,
		e.NackCountRate,
		e.PacketsLossPercent,
		e.Rtt,
		e.DelayStart,
		e.PreStart,
		e.Ua,
		e.BrowserName,
		e.BrowserVersion,
		e.OsName,
		e.OsPlatform,
		e.OsVersion,
		e.DeviceType,
		e.Geohash,
		e.AccuracyRadius,
		e.Latitude,
		e.Longitude,
		e.Timezone,
		e.CityNameEn,
		e.CountryNameEn,
		e.ContinentNameEn,
	)

	_ = fmt.Sprintf(`INSERT INTO test_metrics.test
		(
			ClientId,
			EventType,
			Data,
			Time,
			ClientConnectionTime,
			SessionDurationSeconds,
			Name,
			Provider,
			Platform,
			Merchant,
			Account,
			ProjectId,
			Fp,
			PlayerVersion,
			Ip,
			RealIp,
			Forward,
			BitRate,
			FramesDecodedRate,
			NackCountRate,
			PacketsLossPercent,
			Rtt,
			DelayStart,
			PreStart,
			Ua,
			BrowserName,
			BrowserVersion,
			OsName,
			OsPlatform,
			OsVersion,
			DeviceType,
			Geohash,
			AccuracyRadius,
			Latitude,
			Longitude,
			Timezone,
			CityNameEn,
			CountryNameEn,
			ContinentNameEn) VALUES (%d, '%s', '%s', toDateTime('%s', 'Etc/UTC'), toDateTime('%s', 'Etc/UTC'), %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s',
			 %f, %f, %f, %f, %f, %f, %f, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %f, %f, '%s', '%s', '%s', '%s')`,
		e.ClientId,
		e.EventType,
		e.Data,
		e.Time,
		e.ClientConnectionTime,
		e.SessionDurationSeconds,
		e.StreamName,
		e.StreamPluginName,
		e.Provider,
		e.Platform,
		e.Merchant,
		e.Account,
		e.ProjectId,
		e.Fp,
		e.PlayerVersion,
		e.Ip,
		e.RealIp,
		e.Forward,
		e.BitRate,
		e.FramesDecodedRate,
		e.NackCountRate,
		e.PacketsLossPercent,
		e.Rtt,
		e.DelayStart,
		e.PreStart,
		e.Ua,
		e.BrowserName,
		e.BrowserVersion,
		e.OsName,
		e.OsPlatform,
		e.OsVersion,
		e.DeviceType,
		e.Geohash,
		e.AccuracyRadius,
		e.Latitude,
		e.Longitude,
		e.Timezone,
		e.CityNameEn,
		e.CountryNameEn,
		e.ContinentNameEn,
	)

	if err := c.conn.AsyncInsert(ctx, testStatement, false); err != nil {
		return err
	}

	return nil
}
