package db

import (
	"log"

	"github.com/KatherineEbel/bookstore_oauth-api/src/clients/cassandra"
	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

const (
	createTokenQuery      = `INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);`
	getTokenQuery         = `SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;`
	updateExpirationQuery = `UPDATE access_tokens SET expires=? WHERE access_token=?1;`
)

type IDBRepository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
	Create(accessToken.AccessToken) *errors.RestError
	UpdateExpirationTime(accessToken.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) UpdateExpirationTime(t accessToken.AccessToken) *errors.RestError {
	db := cassandra.GetSession()
	if err := db.Query(updateExpirationQuery, t.AccessToken, t.UserId, t.ClientId, t.Expires).Exec(); err != nil {
		return errors.NewDatabaseError()
	}
	return nil
}

func Repository() IDBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*accessToken.AccessToken, *errors.RestError) {
	session := cassandra.GetSession()
	var t accessToken.AccessToken
	err := session.Query(getTokenQuery, id).Scan(&t.AccessToken, &t.UserId, &t.ClientId, &t.Expires)
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &t, nil
}

func (r *dbRepository) Create(t accessToken.AccessToken) *errors.RestError {
	session := cassandra.GetSession()
	if err := session.Query(createTokenQuery, t.AccessToken, t.UserId, t.ClientId, t.Expires).Exec(); err != nil {
		log.Println(err)
		return errors.NewDatabaseError()
	}
	return nil
}
