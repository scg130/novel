module novel

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/codahale/hdrhistogram v0.0.0-00010101000000-000000000000 // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/scg130/tools v0.10.9
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
