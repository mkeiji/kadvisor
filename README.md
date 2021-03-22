# Kadvisor
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/mkeiji/kadvisor/Kadvisor-pipeline/master?label=ci%20%2F%20cd&style=flat-square)

## Setup Dependencies:

1. [Nodejs && npm](https://nodejs.org/en/)
2. [Go lang](https://golang.org/)
3. [Docker](https://www.docker.com/)
4. [Make](https://www.gnu.org/software/make/)

## Setup

- MAKE SURE you have the `.env` file in the root folder

### Using Make

```bash
make dependencies
```

### Manually

- Get `client` dependencies:

```bash
cd client/
```

```bash
npm install
```

- Get `go` dependencies by running the app for the first time:

```bash
go run main.go
```

## Development

### Run Server

```bash
make runServer
```

### Run Client

```bash
make runClient
```

### Run compiled app

```bash
make run
```

### Debug server (intelliJ only)

```bash
make debug
```

### Debug server with client (intelliJ only)

```bash
make runDebug
```

### Build App

```bash
make build
```

### Build Docker image

```bash
make dockerimg
```

### Spin up `test db` (auto deleted when closed)

```bash
make testdb
```

### Spin up `dev db`

```bash
make db
```

#### Just a table example

| Option            | Usage                             |
| ----------------- | --------------------------------- |
| -a=APP, --app=APP | Add some comments in this section |
