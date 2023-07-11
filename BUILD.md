# Building seekr from source

## Dependencies
- go v1.20.3
- tsc

## Build steps

```sh
go generate ./...
tsc --project web
go build
```

## Testing

```sh
go generate ./...
tsc --project web
go test ./...
```

## tsc watch mode
Go generate only has to be executed once after each source code change in the api.
```sh
go generate ./...
```
```sh
tsc --project web --watch true
go run main.go
```
