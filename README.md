# go_basics

## go structure

workspaces -> modules -> packages -> files

## Create a new workspace

```bash
mkdir workspace_name
cd workspace_name
```

## Create a new module

```bash
go mod init github.com/username/repo
```

## Run a go file

```bash
go run main.go
# or
go run .
```

## crete a package

```bash
mkdir package_name
touch package_name/package.go
```

### install a gin package will change go.mod file

```bash
go get -u github.com/gin-gonic/gin
```

## build a go file

```bash
go build .
```

## install a go file to bin will be available to run from anywhere

```bash
go install .
```

```

```
