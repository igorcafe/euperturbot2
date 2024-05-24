set -eu
go test -coverpkg=./... -coverprofile=./coverage.out ./... &> /dev/null
go tool cover -func ./coverage.out