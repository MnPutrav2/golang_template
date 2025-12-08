# GOLANG ARSITEKTUR TEMPLATE

![Logo Aplikasi](./asset/flowchart.png)

## Migrate

Add table
```bash
migrate create -ext sql -dir ./migrate -seq table_name
```

Migration
```bash
migrate -path ./migrate -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

Rollback
```bash
migrate -path ./migrate -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down
```