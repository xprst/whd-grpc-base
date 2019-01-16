module github.com/xprst/whd-grpc-base

require (
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc
	google.golang.org/grpc v1.16.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
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
