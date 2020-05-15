module github.com/rfinochi/golang-workshop-todo

go 1.13

require (
	cloud.google.com/go v0.57.0 // indirect
	cloud.google.com/go/datastore v1.1.0
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rfinochi/golang-workshop-todo/docs v0.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.5.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5 // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.3.3
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/tools v0.0.0-20200513122804-866d71a3170a // indirect
	google.golang.org/api v0.24.0
)

replace github.com/rfinochi/golang-workshop-todo/docs v0.0.0 => ./docs
