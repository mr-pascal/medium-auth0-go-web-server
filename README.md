medium-auth0-go-web-server
===========================

Small Go web server written with Go Fiber to demonstrate securing routes with a JWT middleware.

```sh
### Run
go run .
```

```sh
# Hit the server with authorization header
curl --header "Authorization: Bearer YOUR_TOKEN" http://localhost:8080/ping
```