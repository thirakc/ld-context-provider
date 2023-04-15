package ldcontext

import (
	"errors"
	"testing"

	mockldcontextstore "ld-context-provider/pkg/mock/store/ldcontext"

	"github.com/stretchr/testify/require"
)

func TestCreateLDContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := NewService(&mockldcontextstore.MockLDContextStore{InsertId: "mockId"})
		err := s.CreateLDContext(&LDContextArg{
			Url:         "https://w3id.org/citizenship/v1",
			DocumentUrl: "https://w3c-ccg.github.io/citizenship-vocab/contexts/citizenship-v1.jsonld",
			Content:     nil,
		})
		require.NoError(t, err)
	})

	t.Run("should fail when save error", func(t *testing.T) {
		s := NewService(&mockldcontextstore.MockLDContextStore{ErrSave: errors.New("save error")})
		err := s.CreateLDContext(&LDContextArg{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "save error")
	})

	t.Run("should fail when insertId is nil", func(t *testing.T) {
		s := NewService(&mockldcontextstore.MockLDContextStore{InsertId: nil})
		err := s.CreateLDContext(&LDContextArg{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "result id is empty")
	})
}
