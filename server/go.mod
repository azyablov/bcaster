module github.com/azyablov/bcaster/server

go 1.18

replace github.com/azyablov/bcaster/protos/bcaster => ../protos/bcaster

require (
	github.com/azyablov/bcaster/msgq v0.0.0-00010101000000-000000000000
	github.com/azyablov/bcaster/protos/bcaster v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/azyablov/bcaster/msgq => ../msgq
