# "Go Programming Language" Book Notes

## Run code

```
go run goFile.go
```

Example:

```
go run ./main.go
```

## New Project

Create new module

```
go mod init repo-or-project-name
```

## Install third party packages

From the root folder of the project (where is located the go.mod file):

```
go get package_repo_url
```

Example:

```
go get github.com/donvito/hellomod
```

## Install all required third party packages

From the root folder of the project (where is located the go.mod file):

```
go mod tidy
```

## Test

```
go test ./folder_with_tests
```

Example:

```
go test ./topics/28-basic-testing/
```

or execute tests located in all subdirectories:

```
go test ./...
```

## Test Coverage

```
go test -cover ./...
```

Example:

```
go test -cover ./...
```

### More coverage info

```
go test ./... -coverprofile=coverage
go tool cover -func=coverage
```

Example:

```
go test ./... -coverprofile=coverage
go tool cover -func=coverage
```

### HTML coverage info

```
go test ./... -coverprofile=coverage
go tool cover -html=coverage
```

Example:

```
go test ./... -coverprofile=coverage
go tool cover -html=coverage
```

## Test Profile

Install [graphviz](https://graphviz.org/)

```
go test ./... -cpuprofile=profile
go tool pprof profile
top
list functionName
web
pdf
quit
```

Example:

```
go test ./... -cpuprofile=profile
go tool pprof profile
top
list Fibonacci
web
pdf
quit
```

## Tooling

### Air

- Install Air (Live reload for Go apps)

```
go install github.com/cosmtrek/air@latest
air init
air
```

- Execute Air with custom .air.toml

```
air -c .air.custom.toml
```

Example:

```
air -c .air.webserver.toml
```

### Mockery

- Install Mockery

```
brew install mockery
```

or download from [https://github.com/vektra/mockery/releases](https://github.com/vektra/mockery/releases).

- Use Mockery (from the folder where is located the interface)

```
mockery --name=InterfaceName
```

### VSCode Extensions

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Golang Snippets](https://marketplace.visualstudio.com/items?itemName=honnamkuan.golang-snippets)
  -- Examples: inte, st, ne, ...
- [Paste JSON as Code](https://marketplace.visualstudio.com/items?itemName=quicktype.quicktype)
- [Generate Method Stubs for Golang Interfaces](https://marketplace.visualstudio.com/items?itemName=ricardoerikson.vscode-go-impl-methods)
- [Go Mock](https://marketplace.visualstudio.com/items?itemName=nawath.go-lazy-mock)
- [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)
- [Better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml)
- [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens)
