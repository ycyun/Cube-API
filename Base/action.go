package Base

func (s Struct) GetID() string {
	return s.ID
}

func (s Struct) SetID(newID string) bool {
	s.ID = newID
	return true
}
