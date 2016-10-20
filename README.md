- Switch to cmd/monitor
- Run make
- Run ./monitor --brokers=[journal] --datamanager=[data-manager] --topic=[topic to listen for (ie..reservation)
example:
./monitor --brokers=http://localhost:8081 --datamanager=http://localhost:6001 --topic=reservation&

