package storage

import "context"

type Storage interface {
	Write(ctx context.Context, key string, data []byte) error
	Read(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	GetServingURL(ctx context.Context, key string) (string, error)
}
