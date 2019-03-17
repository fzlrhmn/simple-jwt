# User Authentication using Go-Kit

## Feature
1. Rate limiting (1 call/minute)
2. JWT Token
3. Configuration Based

## Prerequisite
- install [goose](https://github.com/pressly/goose)
- Run Migration by running `cd config/migration && goose postgres "user={postgre_username} password={postgre_password} dbname=simplejwt sslmode=disable" up`
- Run `make build` to build the app
- Run `./build/user`

## Step to test
1. Create a user
```shell
curl -X POST \
  http://localhost:8796/1.0/user \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username": "{username}",
	"password": "{password}"
}'
```
it will return response
```json
{
    "data": {
        "username": "{username}"
    },
    "meta": {
        "now": 1552844417203127000,
        "requestId": "f93df545-06b7-4723-a94a-efc895c4d86e"
    }
}
```

2. Signin a user
```shell
curl -X POST \
  http://localhost:8796/1.0/user/signin \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username": "{username}",
	"password": "{password}"
}'
```
it will return response something like below
```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InBhaWpvMSIsImV4cCI6MTU1NDA1NDE2OSwiaXNzIjoiRmFpc2FsIFJhaG1hbiJ9.coegp2JyAGRnzBYsmehn1mCkCr2alZMUWpyiNKD0Lvk"
    },
    "meta": {
        "now": 1552844569804990000,
        "requestId": "a8b34098-77b9-4be4-9651-1852a76c8de3"
    }
}
```