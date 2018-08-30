go-kafakesque
--------------

[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![Build status](https://circleci.com/gh/packetloop/go-kafkaesque.svg?style=shield&circle-token=ebd68735d49a76441b9272111ba0b12d472ee4d9)](https://circleci.com/gh/packetloop/go-kafkaesque)

[godocs]: https://godoc.org/github.com/packetloop/go-kafkaesque

A Go binding for [Kafka Admin Service](https://github.com/packetloop/kafka-admin-service)
Since I couldn't manage to find one, hence, write a
new one. One of the intention of having this package is
to allow me to easily write a Terraform provider.

## Usage:

Import package
```bash
go get github.com/packetloop/go-kafkaesque
```

For package dependency management, we use dep:
```bash
go get -u github.com/golang/dep/cmd/dep
```

If new package is required, pls run below command
after go get. For more information about dep, please
visit this URL https://github.com/golang/dep.
```bash
dep ensure
```

Run test:
```bash
make test
```

To maintain codebase quality with static checks and analysis:
```bash
make run
```

Examples:
```go
package main

import (
	"fmt"

	gokafkaesqueue "github.com/packetloop/go-kafkaesque"
)
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
