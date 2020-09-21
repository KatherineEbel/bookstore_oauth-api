package accessToken

import (
	"fmt"
	"strings"

	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

type IRepository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type IService interface {
	GetBId(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository IRepository
}

func Service(repo IRepository) IService {
	return &service{repository: repo}
}

func (s *service) GetBId(id string) (*AccessToken, *errors.RestError) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("invalid token id: %s", id))
	}
	return s.repository.GetById(id)
}
