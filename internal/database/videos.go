package database

import (
	"context"
	"database/sql"
	"fmt"
)

type VideoRecord struct {
	ID            int
	Filename      string
	Title         string
	Duration      float64
	Format        string
	Resolution    string
	HLSStorageKey string
	OriginalKey   string
}

func InitVideoTable(db *Database) error {

	conn := db.GetConnection()
	if conn == nil {
		return fmt.Errorf("db conn is nil")
	}

	query := `
	CREATE TABLE IF NOT EXISTS videos (
	    	id INTEGER PRIMARY KEY AUTOINCREMENT,
	    	filename TEXT NOT NULL,
	    	title TEXT,
	    	duration REAL,
	    	format TEXT,
	    	resolution TEXT,
	    	hls_storage_key TEXT,
	    	original_key TEXT
	);
	`
	_, err := conn.Exec(query)
	if err != nil {
		fmt.Printf("Error while creating video table: %v\n", err)
	}
	return nil
}

func (db *Database) NewVideo(video *VideoRecord) (int, error) {
	conn := db.GetConnection()
	if conn == nil {
		return 0, fmt.Errorf("db conn is nil")
	}

	query := `
	INSERT INTO videos (filename, title, duration, format, resolution, hls_storage_key, original_key)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := conn.ExecContext(context.Background(), query,
		video.Filename, video.Title, video.Duration, video.Format,
		video.Resolution, video.HLSStorageKey, video.OriginalKey)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *Database) GetVideoByID(id int) (*VideoRecord, error) {
	conn := db.GetConnection()
	if conn == nil {
		return nil, fmt.Errorf("db conn is nil")
	}

	query := `
	SELECT id, filename, title, duration, format, resolution, hls_storage_key, original_key
	FROM videos WHERE id = ?
	`

	row := conn.QueryRowContext(context.Background(), query, id)

	var video VideoRecord
	err := row.Scan(&video.ID, &video.Filename, &video.Title, &video.Duration,
		&video.Format, &video.Resolution, &video.HLSStorageKey, &video.OriginalKey)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no video with id: %d", id)
	} else if err != nil {
		return nil, fmt.Errorf("error while searching video with id: %d, err: %v", id, err)
	}

	return &video, nil
}

func (db *Database) UpdateVideoHLSStorageKey(id int, hlsStorageKey string) error {
	conn := db.GetConnection()
	if conn == nil {
		return fmt.Errorf("db conn is nil")
	}

	query := `
	UPDATE videos SET hls_storage_key = ? WHERE id = ?
	`

	result, err := conn.ExecContext(context.Background(), query, hlsStorageKey, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no video with id: %d", id)
	}

	return nil
}
