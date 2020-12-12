package gateways

// IFZF .
type IFZF interface {
	Find(list interface{}, itemFunc func(int) string, previewFunc func(int, int, int) string) (int, error)
	Close()
	Sync() error
}
