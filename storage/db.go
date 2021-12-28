package storage

type Database interface {
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Close() error
}
