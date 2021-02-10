Dog Package
===========

```
godoc -http :8080
go test -coverprofile .coverate.out
go tool cover -html=.coverate.out
go test -bench .
```