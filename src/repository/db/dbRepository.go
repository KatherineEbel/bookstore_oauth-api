package db

import (
	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
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
	return nil, errors.NewDatabaseError()
}
