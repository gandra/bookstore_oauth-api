package app

import (
	"github.com/gandra/bookstore/oauthapi/src/clients/cassandra"
	"github.com/gandra/bookstore/oauthapi/src/domain/access_token"
	"github.com/gandra/bookstore/oauthapi/src/http"
	"github.com/gandra/bookstore/oauthapi/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
