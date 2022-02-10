package tester

import (
	"clickhouse-examples/internal/clickmail"
	"clickhouse-examples/internal/clicknative"
	"clickhouse-examples/internal/config"
	"clickhouse-examples/internal/event"
	"log"
	"time"
)

func RunClickhouseNative() {
	log.Println("Test clickhouse native lib")

	config := config.GetConfig()

	clicknative, err := clicknative.NewClickhouse(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := clicknative.PrepareDatabase(); err != nil {
		log.Fatal(err)
	}

	var writedRecords uint64 = 0
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		e := event.NewEvent()
		if err := clicknative.Write(e); err != nil {
			log.Fatal(err)
		} else {
			writedRecords++
		}
	}

	log.Println("Writed ", writedRecords, " time: ", time.Since(startTime))
}

func RunClickhouseMailLib() {
	log.Println("Test clickhouse mail lib")

	config := config.GetConfig()

	clickmail, err := clickmail.NewClickhouse(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := clickmail.PrepareDatabase(); err != nil {
		log.Fatal(err)
	}

	// e := event.NewEvent()
}
