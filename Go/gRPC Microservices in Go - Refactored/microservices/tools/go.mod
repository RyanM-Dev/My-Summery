module github.com/huseyinbabal/microservices/tools

go 1.27.3

require (
	github.com/huseyinbabal/microservices-proto v0.0.0
	github.com/twmb/franz-go v1.22.2
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/klauspost/compress v1.20.3 // indirect
	github.com/pierrec/lz4/v4 v4.1.33 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.15.0 // indirect
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260928230214-8a89bd6388cc // indirect
	google.golang.org/grpc v1.85.0 // indirect
)

replace github.com/huseyinbabal/microservices-proto => ../proto
