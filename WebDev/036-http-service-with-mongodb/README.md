# HTTP Service with MongoDB

## Usage

### Read user

```
curl http://localhost:8080/user/503
```

### Create user

```
curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Gender":"male","Age":32,"Id":"503"}' http://localhost:8080/user
```

### Delete user

```
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/503
```