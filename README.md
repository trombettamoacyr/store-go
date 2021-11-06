# store-go

This golang mvn application is a simple store with the actions add, alter and delete product.

### Dependencie:

- [Postgres](github.com/lib/pq) - Postgres driver

### Install dependencies
``` 
go build
``` 

### Create postgres container
``` 
make docker-start
``` 
Stop container
``` 
make docker-stop
``` 

### Run application
``` 
go run .
```

### Run migration
http://localhost:8080/migration
