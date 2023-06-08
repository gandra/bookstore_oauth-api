package db

import (
	"github.com/gandra/bookstore/oauthapi/src/clients/cassandra"
	"github.com/gandra/bookstore/oauthapi/src/domain/access_token"
	"github.com/gandra/bookstore/oauthapi/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("db conn not implemented")
}
