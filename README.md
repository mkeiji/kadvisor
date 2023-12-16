# Kadvisor
[![Build](https://github.com/mkeiji/kadvisor/actions/workflows/ci-build.yml/badge.svg)](https://github.com/mkeiji/kadvisor/actions/workflows/ci-build.yml) [![Test](https://github.com/mkeiji/kadvisor/actions/workflows/ci-test.yml/badge.svg)](https://github.com/mkeiji/kadvisor/actions/workflows/ci-test.yml) [![Deployment](https://github.com/mkeiji/kadvisor/actions/workflows/cd-deployment.yml/badge.svg)](https://github.com/mkeiji/kadvisor/actions/workflows/cd-deployment.yml)

## Setup Dependencies:

1. [Nodejs && npm](https://nodejs.org/en/)
2. [Nx](https://nx.dev/)
3. [Go lang](https://golang.org/)
4. [Docker](https://www.docker.com/)
5. [Make](https://www.gnu.org/software/make/)

## Setup

- MAKE SURE you have the `.env` file in the root folder

### Using Make

```bash
make dependencies
```

### Manually

- Get `client` dependencies:
```bash
npm install -g nx
```

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

#### Environment Variables

| Option            | Description                               |
| ----------------- | ----------------------------------------- |
| APP_ENV           | server run mode (options: DEV, PROD, STG) |
| PORT              | server run port                           |
| DB_TYPE           | database (eg: mysql)                      |
| DB_HOST           | database address (eg: localhost:3306)     |
| DB_NAME           | database name                             |
| DB_USER           | database user                             |
| DB_PASS           | database password                         |
| DB_SQLITE         | sqlite ver if type is sqlite (eg: sqlite3)|
| DB_SQLITE_PATH    | sqlite path (if applicable)               |
