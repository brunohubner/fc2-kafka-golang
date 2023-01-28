# Apache Kafka with Golang

#### A simple example of Consumer and Producer Kafka implemented using Golang
#

- Kafka C++ lib
```bash
sudo apt install librdkafka-dev
```

- Create topic
```bash
kafka-topics --create --bootstrap-server=localhost:9092 --topic=test --partitions=3
```

- Monitoring Consumer logs
```bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test
```
