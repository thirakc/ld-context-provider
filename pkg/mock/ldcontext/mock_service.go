package ldcontext

import (
	ldcontextsvc "ld-context-provider/pkg/controller/service/ldcontext"
	"ld-context-provider/pkg/internal/common"
)

type MockService struct {
	ErrCreateLDContext common.Error
}

func (s *MockService) CreateLDContext(*ldcontextsvc.LDContextArg) common.Error {
	if s.ErrCreateLDContext != nil {
		return s.ErrCreateLDContext
	}
	return nil
}
