package ldcontext

type MockLDContextStore struct {
	InsertId any
	ErrSave  error
}

func (m *MockLDContextStore) Save(any) (any, error) {
	if m.ErrSave != nil {
		return nil, m.ErrSave
	}
	if m.InsertId == nil {
		return nil, nil
	}
	return m.InsertId, nil
}
