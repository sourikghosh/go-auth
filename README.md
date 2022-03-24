# go-auth
go-auth is an auth service written in <em>golang</em>. It exposes few basic operations like register, login, and protected get profile. It is using <b>mysql</b> database
for storing the users. [gorm](https://gorm.io/) is used for interacting and making complex sql query. [JWT](https://jwt.io/) for authZ


## Installation Guide
prerequisite go, docker, docker-compose should be installed

`git clone git@github.com:sourikghosh/go-auth.git` clone the repository<br>
`cp .env.example .env` put all the envs inside the .env file <br>
`make up` for running the mysql database<br>
`go run cmd/main.go` to run the auth service

## API examples
- Register - stores new user if not already exist in db
  ```bash
  curl --request POST \
  --url http://localhost:6969/auth-svc/v1/register \
  --header 'Content-Type: application/json' \
  --data '{
	"userName":"sourikghosh",
	"fullName":"Sourik Ghosh",
	"password":"1someThingComplex<31"
  }'
  ```
- Login - user login compares the cred and returns access token
  ```bash
  curl --request POST \
  --url http://localhost:6969/auth-svc/v1/login \
  --header 'Content-Type: application/json' \
  --data '{
	"userName":"sourikghosh",	
	"password":"1someThingComplex<31"
  }'
  ```
- Profile - profile is a protected route that authZ the user and then returns basic user details
  ```bash
  curl --request GET \
  --url http://localhost:6969/auth-svc/v1/profile \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgxMzUxNDEsInVzZXJfaWQiOjF9.LOlYbA3NEcDIPwR8IUrP0-VgwnJe8LAEsC4h742qudI'
  ```

## Project Structure
```bash
├── Dockerfile
├── Makefile
├── cmd                       // entrypoint to the exe binary
│   └── main.go
├── deployments              // has everthing related to deployment
│   └── docker-compose.yml
├── go.mod
├── go.sum
├── implementation           // holds all the internal business logic.
│   └── auth
│       ├── entity.go
│       ├── repository.go
│       └── sevice.go
├── pkg                      // conatins 3rd party-packages 
│   ├── config
│   │   ├── configurations.go
│   │   └── constants.go
│   ├── endpoint.go
│   ├── errors.go
│   ├── hash.go
│   ├── hash_test.go
│   ├── jwt.go
│   ├── logger
│   │   └── logger.go
│   └── scripts
│       └── init_db.go
└── transport                // deals with the transport layer
    ├── endpoints
    │   ├── common.go
    │   ├── loginHnadler.go
    │   ├── register.go
    │   └── register_test.go
    ├── http
    │   └── routes.go
    └── requestResponse.go
```
