go build hello.go
go clean hello.go
go build
go run 0_hello_world # delete go.mod and run this

go vet # idea covers you
go format # idea covers you

https://go.dev/ref/mod#go-install
go install github.com/viggin543/awesomne-golang/0_hello_world@latest
go env | grep -i path
# > GOPATH="/Users/domrevigor/go"
file $GOPATH/bin/0_hello_world


# clone, compile, execute, delete binaries
go run github.com/viggin543/awesomne-golang/0_hello_world@latest

###################################### releasing binaries and libraries was never so easy !!   ##################################

# what would have happened if this was a private repo ?