package root

// Hash interface
type Hash interface {
	Generate(s string) (string, error)
	Compare(hash string, s string) error
}
