package main

import (
	"log"

	"clickhouse-examples/internal/tester"
)

func main() {
	log.Println("clickhouse examples")

	tester.RunClickhouseMailLib()

	// tester.RunClickhouseNative()
}
