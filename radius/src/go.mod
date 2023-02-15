module fbc/cwf/radius

replace (
	fbc/lib/go/machine => ../lib/go/machine
	fbc/lib/go/radius => ../lib/go/radius
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	fbc/lib/go/machine v0.0.0-00010101000000-000000000000
	fbc/lib/go/radius v0.0.0-00010101000000-000000000000
	github.com/donovanhide/eventsource v0.0.0-20171031113327-3ed64d21fb0b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.1
	github.com/stretchr/testify v1.4.0
	go.opencensus.io v0.21.0
	go.uber.org/atomic v1.4.0
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	google.golang.org/grpc v1.21.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
