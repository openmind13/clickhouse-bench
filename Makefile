
sync:
	go run cmd/main.go \
						--clickhouse_native_url=127.0.0.1:9000 \
						--database=test_metrics \
						--table=test \
						--use_async=false \
						--workers_count=3 \
						--working_time_seconds=10 \
						--engine=MergeTree

async:
	go run cmd/main.go \
						--clickhouse_native_url=127.0.0.1:9000 \
						--database=test_metrics \
						--table=test \
						--use_async=true \
						--workers_count=3 \
						--working_time_seconds=10 \
						--engine=MergeTree
