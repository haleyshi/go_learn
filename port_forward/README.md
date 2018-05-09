```bash
# Shell 1
go run main.go --target 127.0.0.1:8080 --port 1337

# Shell 2
go run main.go --target 127.0.0.1:1337 --port 1338

# Shell 3
cd src/go_learn/restful/simple_web_server/
go run simple_web_server.go

# Browser
http://localhost:1338
```