package storage

type Storage interface {
	SetData(key string, data any) error
	GetData(key string) (any, error)
}
