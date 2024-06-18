# 脚手架生成server代码
[cwgo命令参考](https://www.cloudwego.io/docs/cwgo/tutorials/auto-completion/)
```shell
mkdir -p demo/demo_thrift
cd demo/demo_thrift
cwgo server --type RPC --module github.com/Ccmuyu/biz-demo/gomall/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift


  demo_thrift [main] ⚡  go work use .
☁  demo_thrift [main] ⚡  go run . 
```

# docker compose
```shell
docker compose up -d
```