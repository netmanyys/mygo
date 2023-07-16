
❯ go mod init github.com/netmanyys/mygo/tree/main/projects/go-bookstore

go: modules disabled by GO111MODULE=off; see 'go help modules'
❯ export GO111MODULE="on"

go mod init github.com/netmanyys/mygo/tree/main/projects/go-bookstore
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"