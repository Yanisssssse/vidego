package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	baseDir string
}

func NewLocalStorage(baseDir string) (*LocalStorage, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}

	return &LocalStorage{baseDir: baseDir}, nil
}

func (l *LocalStorage) Write(ctx context.Context, key string, data []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	fp := filepath.Join(l.baseDir, key)

	dirPath := filepath.Dir(fp)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(fp, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (l *LocalStorage) Read(ctx context.Context, key string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	fp := filepath.Join(l.baseDir, key)
	data, err := os.ReadFile(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", key)
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return data, nil
}

func (l *LocalStorage) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	fp := filepath.Join(l.baseDir, key)
	if err := os.Remove(fp); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (l *LocalStorage) Exists(ctx context.Context, key string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	}

	fp := filepath.Join(l.baseDir, key)
	_, err := os.Stat(fp)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("failed to check file existence: %w", err)
}

func (l *LocalStorage) GetServingURL(ctx context.Context, key string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	fp := filepath.Join(l.baseDir, key)
	if _, err := os.Stat(fp); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file not found: %s", key)
		}
		return "", fmt.Errorf("failed to check file: %w", err)
	}

	return fmt.Sprintf("/video/%s", key), nil
}
