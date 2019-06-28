## Usage

- Make sure you have a MySQL instance running at `localhost:9306/gotest` and user `user:password`. Refer to the Database section below
- `go run gin-rest/`
- Server will run at 8080

## Routes
```
GET     /person         Get all Person records
GET     /person/:id     Get a Person record with given id
POST    /person         Create a new Person record
DELETE  /person/:id     Delete an existing Person record with given id
PUT     /person/:id     Update an existing Person record with given id and new values on request body
```

## Database

- 
