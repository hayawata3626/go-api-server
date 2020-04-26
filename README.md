# go-api-server

## make `.env` file on project root directory

```
DB_USER=
DB_PASSWORD=
DB_NAME=
DATABASE_URL=tcp:localhost:3306*$DB_USER/$DB_USER/$DB_PASSWORD
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
