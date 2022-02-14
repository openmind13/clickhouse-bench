package config

import (
	"flag"
	"time"
)

var (
	useClickhouseNativeFlag = flag.Bool("use_native", true, "use clickhouse native")
	clickhouseNativeUrlFlag = flag.String("clickhouse_native_url", "", "clickhouse native url")
	clickhouseStdUrlFlag    = flag.String("clickhouse_std_url", "", "clickhouse std url")
	databaseFlag            = flag.String("database", "", "database name")
	tableFlag               = flag.String("table", "", "table name")
	isAsyncFlag             = flag.Bool("use_async", false, "is async ????")
	workersCountFlag        = flag.Int("workers_count", 1, "worker count (up to 10)")
	workingTimeSecondsFlag  = flag.Int("working_time_seconds", 1, "working time (in seconds)")
	databaseEngineFlag      = flag.String("engine", "MergeTree", "database engine")
)

var Config config

type config struct {
	UseClickhouseNative bool
	ClickhouseNativeUrl string
	ClickhouseStdUrl    string
	Database            string
	Table               string
	WorkersCount        int
	WorkingTime         time.Duration
	IsAsync             bool
	Engine              string
}

func InitConfig() {
	flag.Parse()

	Config = config{
		UseClickhouseNative: *useClickhouseNativeFlag,
		ClickhouseNativeUrl: *clickhouseNativeUrlFlag,
		ClickhouseStdUrl:    *clickhouseStdUrlFlag,
		Database:            *databaseFlag,
		Table:               *tableFlag,
		IsAsync:             *isAsyncFlag,
		WorkersCount:        *workersCountFlag,
		WorkingTime:         time.Duration(*workingTimeSecondsFlag) * time.Second,
		Engine:              *databaseEngineFlag,
	}
}
