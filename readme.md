## Instructions
- install docker
- run `docker-compose up -d`
- run `go run main.go`
- now visit `http://localhost:8080/users` which will return empty array list.

NB: As per the task only have to handle request and response with db connection, so I have only added one function to return all users. Have not added other CRUD operation or seeder or high level migration files.