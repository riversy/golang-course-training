Math Package
==================

Please use the following commands to test the app. 

```bash
go test ./mymath
go test -bench ./mymath
go test -coverprofile .coverate.out ./mymath
go tool cover -html=.coverate.out
godoc -http :8080
```