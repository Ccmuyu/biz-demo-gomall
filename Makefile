.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC --module github.com/Ccmuyu/biz-demo/gomall/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-proto:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/Ccmuyu/biz-demo/gomall/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-thrift-api
gen-thrift-api:
	@cd demo/thrift_api && cwgo server --type HTTP --idl ../../idl/hello_api.thrift --module github.com/Ccmuyu/biz-demo/gomall/thrift_api --service thrift_api
