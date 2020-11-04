# Server

## Unit Tests
### Unit Tests Dependencies
1. [ginkgo](https://github.com/onsi/ginkgo)
2. [gomega](https://github.com/onsi/gomega)
3. [gomock](https://github.com/golang/mock)

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
