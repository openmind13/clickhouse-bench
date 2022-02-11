# Clickhouse examples

## Cmd's

```bash
service clickhouse-server start

service clickhouse-server stop
```

## Benchmarks

### Clickhouse-native lib

#### Синхронная вставка на локальной машине

- 1 поток-воркер: ~ 600 inserts/sec
- 2 потока-воркера: ~ 1000 inserts/sec
- 3 потока-воркера: ~ 1400 inserts/sec
- 4 потока-воркера: ~ 1400 inserts/sec
- 5 потоков-воркеров: ~ 1400 inserts/sec
- 10 потоков-воркеров: ~ 1450 inserts/sec

#### Асинхронная вставка на локальной машине

- 1 поток-воркер: ~ 1800 inserts/sec
- 2 потока-воркера: ~  inserts/sec
- 3 потока-воркера: ~ 4650 inserts/sec
- 4 потока-воркера: ~ 6000 inserts/sec