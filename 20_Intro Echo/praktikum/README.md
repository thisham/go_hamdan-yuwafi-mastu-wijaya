# Echo Golang Static API

## Overview

This API was created using echo golang framework with static data. Featured with hot-reload
for development purposes.

## Tech-Stack

### Pre-Requiresities

- Golang 1.17.5

### Libraries and Framework

- Echo, Golang web framework ([documentation](https://echo.labstack.com))
- Air, Golang server hot-reloader ([documentation](https://github.com/cosmtrek/air))

## Getting Started

### Starting Echo Server

Echo server was run on port 8000 as default port. Can be changed on `main.go` for sure.

```bash
go run main.go
```

### Start Hot-Reloading

With hot-reloading, you needn't to start and stop golang echo manually. When changes detected,
golang server will be automatically restarted. You may need to add this package as global package
before using it into development server.

```bash
go install github.com/cosmtrek/air@latest
```

After installing this packages, you need to start the hot-reloading server.

```bash
air
```

## Contributing

Nilai 80+ lah mas :)
