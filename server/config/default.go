package config

const Default = `
db:
  host: 127.0.0.1
  port: "5431"
  user: postgres
  dbname: monitor
  password: postgres
  sslmode: disable
nats:
  host: nats://localhost:4221
  topic: check
  queue: monitors
threshold: 101
`
