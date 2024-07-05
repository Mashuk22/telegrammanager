module github.com/Mashuk22/telegrammanager/userservice

go 1.22.0

require (
	github.com/Mashuk22/telegrammanager/pkg/userpb v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.65.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Mashuk22/telegrammanager/userservice/pkg/userpb => ./pkg/userpb
