package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
)

const DRIVER_NAME = "clickhouse"

var (
	clickhouseURL = flag.String("url", "", "setting up a clickhouse url")
	database      = flag.String("database", "", "setting up a database name")
	username      = flag.String("username", "", "setting up a clickhouse username")
	password      = flag.String("password", "", "setting up a clichouse password")
)

const (
	SQL_CREATE_TABLE_STATEMENT = `
			CREATE TABLE test (
				Id UInt32,
				Type String,
				Time DateTime('Europe/Moscow')
			)
			Engine = MergeTree
			ORDER BY (Id)
	`

	SQL_DROP_TABLE_STATEMENT = "DROP TABLE IF EXISTS test"
)

type ClientEvent struct {
	Id   uint32
	Type string
	Time time.Time
}

func main() {
	log.Println("clickhouse examples")

	flag.Parse()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{*clickhouseURL},
		Auth: clickhouse.Auth{
			Database: *database,
			Username: *username,
			Password: *password,
		},
		DialTimeout: 5 * time.Second,
		Debug:       true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx := context.Background()

	if err := conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	if err := conn.Exec(ctx, SQL_DROP_TABLE_STATEMENT); err != nil {
		log.Fatal(err)
	}

	if err := conn.Exec(ctx, SQL_CREATE_TABLE_STATEMENT); err != nil {
		log.Fatal(err)
	}

	id := uuid.New().ID()
	startEvent := ClientEvent{
		Id:   id,
		Type: "Start",
		Time: time.Now(),
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO test")
	if err != nil {
		log.Fatal(err)
	}
	if err := batch.AppendStruct(&startEvent); err != nil {
		log.Fatal(err)
	}
	if err := batch.Send(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	endEvent := ClientEvent{
		Id:   id,
		Type: "Stop",
		Time: time.Now(),
	}

	batch, err = conn.PrepareBatch(ctx, "INSERT INTO test")
	if err != nil {
		log.Fatal(err)
	}
	if err := batch.AppendStruct(&endEvent); err != nil {
		log.Fatal(err)
	}
	if err := batch.Send(); err != nil {
		log.Fatal(err)
	}
}
