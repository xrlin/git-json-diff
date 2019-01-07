SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o git-json-diff-osx cmd/git-json-diff/main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o git-json-diff cmd/git-json-diff/main.go

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o git-json-diff.exe cmd/git-json-diff/main.go