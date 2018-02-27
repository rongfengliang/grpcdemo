
## code 

```code
service UserLogin{
   rpc AppLogin(UserLoginRequest) returns(UserLOginResponse);
}
message UserLoginRequest{
   string name =1;
   string password=2;
}
message UserLOginResponse{
   int code=1;
   string message=2;
   string token=3;
}
```
## build with xservice
```bash
protoc -I . userlogin.proto --xservice_out=./code --go_out=./code 
```
## build with grpc
```bash
protoc -I . userlogin.proto --go_out=plugins=grpc:rpc
```
