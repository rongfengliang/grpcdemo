# grpc with restful
build with grpc-gateway  && xservie also is a optional

## install packages
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
## build proto
```bash
  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  userlogin.proto
```
## Implement   service in gRPC 
```bash
for server just see app/server/main.go
for client(grpc call) see app/client/main.go
```
## Generate reverse-proxy
```bash
  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  userlogin.proto
```
## Generate swagger definitions
```bash
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  userlogin.proto
```
## write proxy code 
```bash
for proxy just see proxy/main.go
```
## request restful api(all the api call with post)
```bash 
localhost:8089/v1/api/list
localhost:8089/v1/api/echo
```
## requet with grpc 
```bash
cd app
go run server/main.go

cd app
go run client/main.go

```