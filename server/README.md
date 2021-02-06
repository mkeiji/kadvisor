# Server

## Dependencies
[resource](https://golang.org/doc/modules/managing-dependencies)

### Discovering available updates
- List all of the modules that are dependencies of your current module, along with the latest version available for each:
```bash
$ go list -m -u all
```

- Display the latest version available for a specific module:
```bash
$ go list -m -u example.com/theirmodule
```

### Getting a specific dependency version

- To get a specific numbered version, append the module path with an @ sign followed by the version you want:
```bash
$ go get example.com/theirmodule@v1.3.4
```

- To get the latest version, append the module path with @latest:
```bash
$ go get example.com/theirmodule@latest
```


## Unit Tests

### Unit Tests Dependencies

1. [ginkgo](https://github.com/onsi/ginkgo)
2. [gomega](https://github.com/onsi/gomega)
3. [gomock](https://github.com/golang/mock)

## Install frameworks

- ginkgo & gomega:

```bash
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...
```

- gomock w/ mockgen:

```bash
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
```

## Generate package test suit with ginkgo

- Navigate to `package/path` and run:

```bash
ginkgo bootstrap
```

## Generate test file

- Navigate to `package/path` and run:

```bash
ginkgo generate ${nameOfTheStruct}
```

## Running tests

- To run the suite in the current directory, simply run:

```bash
ginkgo #or go test
```

- To run the suites in other directories, simply run:

```bash
ginkgo /path/to/package /path/to/other/package ...
```

- more details for [running tests](https://onsi.github.io/ginkgo/#running-tests)

### Run all test suites

```bash
ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress
```
