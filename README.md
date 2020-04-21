# go-api-server

## make `.env` file on project root directory

```
DATABASE_URL=tcp:${db-host}:3306*${db-name}/${db-user}/${db-password}
```

## Migration

```bash
$ source .env

# migration status check
$ goose status

# migration up
$ goose up

# migration down
$ goose down
```


## start server

```
go run main.go
```
