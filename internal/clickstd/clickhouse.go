package clickstd

import (
	"clickhouse-bench/internal/config"
	"clickhouse-bench/internal/sqlstatement"
	"database/sql"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

const (
	MAX_CONNS = 50
)

type Clickhouse struct {
	db *sql.DB
}

func NewClickhouse() (*Clickhouse, error) {
	c := &Clickhouse{}
	db, err := sql.Open("clickhouse", config.Config.ClickhouseStdUrl)
	if err != nil {
		return nil, err
	}
	c.db = db
	c.db.SetMaxIdleConns(MAX_CONNS)

	if err := c.db.Ping(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Clickhouse) PrepareDatabase() error {
	if _, err := c.db.Exec(fmt.Sprintf(sqlstatement.SQL_DROP_TABLE, config.Config.Table)); err != nil {
		return err
	}
	if _, err := c.db.Exec(fmt.Sprintf(sqlstatement.SQL_CREATE_TABLE, config.Config.Table)); err != nil {
		return err
	}
	return nil
}
