export GOARCH=amd64
export GOOS=darwin
go build -o ./build/mac/secret-api
export GOOS=linux
go build -o ./build/linux/secret-api
export GOOS=windows 
go build -o ./build/win/secret-api.exe