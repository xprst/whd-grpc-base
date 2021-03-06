module github.com/xprst/whd-grpc-base

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20181014144952-4e0d7dc8888f
	github.com/go-sql-driver/mysql v1.4.0
	github.com/go-xorm/core v0.6.0
	github.com/go-xorm/xorm v0.7.1
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.1.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.16.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	go.uber.org/zap => github.com/uber-go/zap v1.9.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190122013713-64072686203f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190110200230-915654e7eabc
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190111180523-db91494dd46c
	google.golang.org/grpc => github.com/grpc/grpc-go v1.2.1-0.20190114234020-98a94b0cb0eb
)
