# go-rest-api-mongodb
Rest Api using Golang and mongodb as a database

Go web server with Mongodb, Crypto and Jwt Authentication.

Structure
~~~
go-rest-api-mongodb
    /controller
      -controller.go
  	/config
      -config.go
    /crypto
      -crypto.go
    /dao
      -dao.go
    /helper
      -helper.go
    /middlewares
      -middleware.go
    /models
      -models.go
    -app.go
    -config.go
~~~

Dependancies
  - go get gopkg.in/mgo.v2
  - go get github.com/gorilla/mux
  - go get github.com/gorilla/handlers
  - go get golang.org/x/crypto/bcrypt
  - go get github.com/google/uuid
  - go get github.com/BurntSushi/toml
  - go get github.com/gorilla/context
  - go get github.com/go-chi/chi
  - go get github.com/dgrijalva/jwt-go