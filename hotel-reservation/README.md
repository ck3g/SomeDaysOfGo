# Hotel reservation backend

## Project outline

- users -> book room from an hotel
- admings -> going to check reservatikon/bookings
- Authentication and authorization -> JWT tockens
- hotels -> CRUD API -> JSON
- rooms -> CRUD API -> JSON
- scripts -> database management -> seeding, migrations

## Resources

### MongoDB driver

Documentation

```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing MongoDB client

```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber

Documentation

```
https://docs.gofiber.io/
```

Installing gofiber

```
go get github.com/gofiber/fiber/v2
```

## Docker

### Installing MongoDB as a Docker container

```
docker run --name mongodb -d -p 27017:27017 mongo:latest
```
