package mock

// Hash mock
type Hash struct{}

// Generate mock
func (h *Hash) Generate(s string) (string, error) {
	return s, nil
}

// Compare mock
func (h *Hash) Compare(hash string, s string) error {
	return nil
}
