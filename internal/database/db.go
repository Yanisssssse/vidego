package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
	mu   sync.RWMutex
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

func (db *Database) Connect(dsn string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("Failed to open database : %w", err)
	}

	if err := conn.Ping(); err != nil {
		return fmt.Errorf("erreur de connexion: %w", err)
	}

	db.conn = conn
	return nil
}

func (db *Database) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.conn != nil {
		return db.conn.Close()
	}
	return nil
}

func (db *Database) GetConnection() *sql.DB {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.conn
}
