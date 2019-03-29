# echo-requests

Runs a server on a port than simply prints out the request & returns a 200 OK response

## usage

```
$ go install github.com/christianscott/echo-requests
$ echo-requests 5555 # then in another session: curl localhost:5555
GET /
User-Agent [curl/7.54.0]
Accept [*/*]
```
