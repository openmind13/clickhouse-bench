package clickmail

import (
	"clickhouse-examples/internal/config"
	"clickhouse-examples/internal/sqlstatement"
	"database/sql"
	"fmt"

	_ "github.com/mailru/go-clickhouse"
)

const (
	MAX_CONNS = 50
)

type Clickhouse struct {
	config config.Config
	db     *sql.DB
}

func NewClickhouse(config config.Config) (*Clickhouse, error) {
	c := &Clickhouse{}
	c.config = config
	db, err := sql.Open("clickhouse", config.ClickmailURL)
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
	if _, err := c.db.Exec(fmt.Sprintf(sqlstatement.SQL_DROP_TABLE, c.config.Table)); err != nil {
		return err
	}
	if _, err := c.db.Exec(fmt.Sprintf(sqlstatement.SQL_CREATE_TABLE, c.config.Table)); err != nil {
		return err
	}
	return nil
}
