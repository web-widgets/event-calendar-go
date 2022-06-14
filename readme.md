Go backend for the Scheduler
==================

### How to start

- start the backend

```bash
go build
./scheduler-go
```

### Configuration

- create config.yml with application parameters and SQLite config

```yaml
db:
  path: db.sqlite
  resetonstart: true
server:
  url: "http://localhost:3000"
  port: ":3000"
  cors:
    - "*"
```




