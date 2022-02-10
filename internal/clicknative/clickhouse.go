package clicknative

import (
	"context"
	"fmt"
	"time"

	"clickhouse-examples/internal/config"
	"clickhouse-examples/internal/event"
	"clickhouse-examples/internal/sqlstatement"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

const (
	WRITE_TIMEOUT = 5 * time.Second
)

type Clickhouse struct {
	conn   driver.Conn
	ctx    context.Context
	config config.Config
}

func NewClickhouse(config config.Config) (*Clickhouse, error) {
	c := &Clickhouse{}
	c.config = config
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{c.config.ClickhouseURL},
		Auth: clickhouse.Auth{
			Database: c.config.Database,
			Username: c.config.Username,
			Password: c.config.Password,
		},
		MaxOpenConns:    500,
		MaxIdleConns:    500,
		ConnMaxLifetime: time.Hour,
		DialTimeout:     5 * time.Second,
		Debug:           true,
		// Compression: &clickhouse.Compression{
		// 	Method: clickhouse.CompressionLZ4,
		// },
	})
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.ctx = context.Background()

	return c, nil
}

func (c *Clickhouse) PrepareDatabase() error {
	if err := c.conn.Exec(c.ctx, fmt.Sprintf(sqlstatement.SQL_DROP_TABLE, c.config.Table)); err != nil {
		return err
	}
	if err := c.conn.Exec(c.ctx, fmt.Sprintf(sqlstatement.SQL_CREATE_TABLE, c.config.Table)); err != nil {
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
	batch, err := c.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", c.config.Table))
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
	if err := c.conn.AsyncInsert(ctx, fmt.Sprintf(`INSERT INTO %s VALUES (%d, %s, %s, %s)`, c.config.Table, e.Id, e.Type, e.Data, e.Time), true); err != nil {
		return err
	}
	return nil
}
