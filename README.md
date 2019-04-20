# languages_go_api

## Update module files.

To add a new module, please use `go get` to download a new  module, to add the module do a `go test` to check and add module to 'go.sum' and 'go.mod'.

## Use Go module with vendor

To activate go vendors, you can use `go mod vendor`, this will create a vendor folder.

To build the application with vendor use the build command as follow: `go build -mod vendor`.