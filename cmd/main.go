package main

import (
	"log"

	"clickhouse-bench/internal/config"
	"clickhouse-bench/internal/tester"
)

func main() {
	log.Println("clickhouse bench")

	config.InitConfig()

	tester.RunClickhouseNative()
}
