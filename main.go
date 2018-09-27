package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.schibsted.io/Schibsted-fr-lab/graphql-poc/resolver"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	var err error
	var s []byte
	var schema *graphql.Schema
	var handler *relay.Handler

	if s, err = ioutil.ReadFile("schema.graphql"); err != nil {
		logrus.WithError(err).Fatal("Couldn't open schema")
	}

	schema = graphql.MustParseSchema(string(s), &resolver.MainResolver{})
	handler = &relay.Handler{Schema: schema}

	r := gin.Default()
	r.POST("/graphql", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	r.Run("127.0.0.1:8080")
}
