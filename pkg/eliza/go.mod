module github.com/junsazanami430u/examples-go/pkg/eliza

go 1.23.0

require (
	connectrpc.com/connect v1.18.1
	github.com/rs/cors v1.11.1
	github.com/stretchr/testify v1.10.0
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/junsazanami430u/jsasanami-eliza/pkg/eliza => github.com/junsazanami430u/examples-go/pkg/eliza v0.0.0-20250308074936-11b478fa78dc
