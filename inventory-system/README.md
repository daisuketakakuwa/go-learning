# inventory-system

## Migrate

migrate -database "postgres://postgres:postgres@localhost:5435/postgres?sslmode=disable" -path db/migrations up

## How to start application

1. goa gen（do this every time you modify design）
```
goa gen inventory-system/api/svc/design  
```
2. goa example
```
goa example inventory-system/api/svc/design  
```
3. goa build (when modify sth, just do this and restart app)
```
go build ./cmd/inventory_system
```
4. start app
```
./inventory_system
```