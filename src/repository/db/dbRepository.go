package db

import (
	"log"

	"github.com/KatherineEbel/bookstore_oauth-api/src/clients/cassandra"
	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

const (
	getTokenQuery = "SELECT accessToken, userId, clientId, expires FROM accessTokens WHERE accessToken=?"
)

type IDBRepository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func Repository() IDBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*accessToken.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Println(err)
		return nil, errors.NewDatabaseError()
	}
	defer session.Close()
	var t accessToken.AccessToken
	err = session.Query(getTokenQuery, id).Scan(&t.Token, &t.UserId, &t.ClientId, &t.Expires)
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &t, nil
}
