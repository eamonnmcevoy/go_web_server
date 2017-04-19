package mock

type Hash struct{}

func (h *Hash) Generate(s string) (string, error) {
	return s, nil
}

func (h *Hash) Compare(hash string, s string) error {
	return nil
}
