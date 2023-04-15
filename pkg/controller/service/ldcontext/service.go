package ldcontext

import (
	"errors"
	"ld-context-provider/pkg/internal/common"

	"github.com/google/uuid"
)

type Storer interface {
	Save(any) (any, error)
}

type Service struct {
	Store Storer
}

func NewService(store Storer) *Service {
	return &Service{store}
}

func (s *Service) CreateLDContext(arg *LDContextArg) common.Error {
	entity := LDContextEntity{
		Id:           uuid.NewString(),
		LDContextArg: arg,
	}
	id, err := s.Store.Save(entity)
	if err != nil {
		return common.NewError(SaveLDContextErrorCode, err)
	}
	if id == nil {
		return common.NewError(SaveResultErrorCode, errors.New("result id is empty"))
	}

	return nil
}
