# Unit Test
    testing
      |__ identity
      |__ mocks
      |__ testcerts
      |__ interfaces.go
      |__ tools.go
      |__ var.go

## Run test
```bash
$ cd ./testing/identity && go test
```

## testcerts folder
Folder with fake certificates for mock up
- testcert.go (to load test certificates found in the testcerts folder)
```golang

import (
	"github.com/kmilodenisglez/cc-identity-go/testcerts"
)

// load certificate with attributes for abac
certBytes, _ := testcerts.Certificates[3].CertBytes()
```

## counterfeiter
Use counterfeiter to generate directives.

Counterfeiter allows you to simply generate test doubles for a given interface.

see https://github.com/maxbrunsfeld/counterfeiter

# Behavior-Driven Development (BDD)
We use [Ginkgo](https://github.com/onsi/ginkgo) which enables the BDD.
Paired with the [Gomega](https://github.com/onsi/gomega) matcher library.