package tester

import (
	"clickhouse-bench/internal/clicknative"
	"clickhouse-bench/internal/config"
	"clickhouse-bench/internal/event"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	writedRecords uint64 = 0
)

func RunClickhouseNative() {
	log.Println("Test clickhouse native lib")

	clicknative, err := clicknative.NewClickhouse()
	if err != nil {
		log.Fatal(err)
	}

	if err := clicknative.PrepareDatabase(); err != nil {
		log.Fatal(err)
	}

	if config.Config.IsAsync {
		benchAsyncronousInsert(clicknative)
	} else {
		benchSyncronousInsert(clicknative)
	}

}

type worker struct {
	Number    int
	EventChan chan event.Event
	click     *clicknative.Clickhouse
}

func newWorker(number int, click clicknative.Clickhouse) worker {
	return worker{
		Number:    number,
		EventChan: make(chan event.Event, 1),
		click:     &click,
	}
}

func (w *worker) runSyncWriting() {
	for {
		e := <-w.EventChan
		if err := w.click.Write(e); err != nil {
			log.Fatal(err)
		} else {
			atomic.AddUint64(&writedRecords, 1)
		}
	}
}

func (w *worker) runAsyncWriting() {
	for {
		e := <-w.EventChan
		if err := w.click.WriteAsync(e); err != nil {
			log.Fatal(err)
		} else {
			atomic.AddUint64(&writedRecords, 1)
			// writedRecords++
		}
	}
}

func benchSyncronousInsert(clicknative *clicknative.Clickhouse) {
	stopChan := make(chan struct{}, 1)
	go func() {
		<-time.After(WORKING_TIME)
		stopChan <- struct{}{}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	workerPool := make(map[int]worker)
	for i := 0; i < config.Config.WorkersCount; i++ {
		w := newWorker(i, *clicknative)
		workerPool[i] = w
		go w.runSyncWriting()
	}

	log.Println("Running syncronous insert with ", config.Config.WorkersCount, " workers")
	for {
		select {
		case <-stopChan:
			log.Println("Writed ", writedRecords, " records")
			os.Exit(0)
		case sig := <-sigChan:
			log.Println("stopped by signal", sig)
			log.Println("Writed ", writedRecords, " records")
			os.Exit(0)
		default:
			e := event.NewEvent()
			select {
			case workerPool[0].EventChan <- e:
			case workerPool[1].EventChan <- e:
			case workerPool[2].EventChan <- e:
			case workerPool[3].EventChan <- e:
			case workerPool[4].EventChan <- e:
			case workerPool[5].EventChan <- e:
			case workerPool[6].EventChan <- e:
			case workerPool[7].EventChan <- e:
			case workerPool[8].EventChan <- e:
			case workerPool[9].EventChan <- e:
			}
		}
	}
}

func benchAsyncronousInsert(clicknative *clicknative.Clickhouse) {
	stopChan := make(chan struct{}, 1)
	go func() {
		<-time.After(WORKING_TIME)
		stopChan <- struct{}{}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	workerPool := make(map[int]worker)
	for i := 0; i < config.Config.WorkersCount; i++ {
		w := newWorker(i, *clicknative)
		workerPool[i] = w
		go w.runAsyncWriting()
	}

	log.Println("Running asyncronous insert with ", config.Config.WorkersCount, " workers")
	for {
		select {
		case <-stopChan:
			log.Println("Writed ", writedRecords, " records")
			os.Exit(0)
		case sig := <-sigChan:
			log.Println("stopped by signal", sig)
			log.Println("Writed ", writedRecords, " records")
			os.Exit(0)
		default:
			e := event.NewEvent()
			select {
			case workerPool[0].EventChan <- e:
			case workerPool[1].EventChan <- e:
			case workerPool[2].EventChan <- e:
			case workerPool[3].EventChan <- e:
			case workerPool[4].EventChan <- e:
			case workerPool[5].EventChan <- e:
			case workerPool[6].EventChan <- e:
			case workerPool[7].EventChan <- e:
			case workerPool[8].EventChan <- e:
			case workerPool[9].EventChan <- e:
			}
		}
	}
}
