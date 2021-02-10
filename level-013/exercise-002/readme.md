Word Count Package
==================

Please use the following commands to test the app. 

```bash
go test ./...
go test -bench ./...
go test -coverprofile .coverate.out
go tool cover -html=.coverate.out
godoc -http :8080
```