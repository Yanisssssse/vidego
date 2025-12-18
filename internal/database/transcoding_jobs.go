package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Yanisssssse/vidego/internal/transcoding"
)

func (db *Database) InitTranscodingJobTable() error {
	conn := db.GetConnection()
	if conn == nil {
		return fmt.Errorf("db conn is nil")
	}

	query := `
	CREATE TABLE IF NOT EXISTS transcoding_jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT ,
			video_id INTEGER,
			status TEXT,
			input_key TEXT,
			output_key TEXT,
			created_at DATETIME,
			started_at DATETIME,
			completed_at DATETIME,
			error_message TEXT
	)
	`

	_, err := conn.Exec(query)
	if err != nil {
		fmt.Errorf("error while creating transcoding jobs table: %v", err)
	}

	return nil
}

func (db *Database) NewTranscodingJob(videoId int, inputKey string, outputKey string) (int64, error) {
	conn := db.GetConnection()
	if conn == nil {
		return 0, fmt.Errorf("db conn is nil")
	}

	query := `
	INSERT INTO transcoding_jobs (video_id, status, input_key, output_key, created_at)
	VALUES (?, ?, ?, ?, ?)
	`

	result, err := conn.ExecContext(context.Background(), query, videoId, transcoding.StatusPending, inputKey, outputKey, time.Now())
	if err != nil {
		return 0, fmt.Errorf("error while creating transcoding job: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %v", err)
	}

	return id, nil
}

func (db *Database) GetNextPendingJob() (*transcoding.Job, error) {
	conn := db.GetConnection()
	if conn == nil {
		return nil, fmt.Errorf("db conn is nil")
	}

	query := `
	SELECT * FROM transcoding_jobs LIMIT 1
	`

	row := conn.QueryRowContext(context.Background(), query)
	var job transcoding.Job
	err := row.Scan(
		&job.ID,
		&job.VideoID,
		&job.Status,
		&job.InputKey,
		&job.OutputKey,
		&job.CreatedAt,
		&job.StartedAt,
		&job.CompletedAt,
		&job.ErrorMessage)
	if err != nil {
		return nil, fmt.Errorf("error getting next pending job: %v", err)
	}

	return &job, nil
}

func (db *Database) UpdateTranscodingJobStatus(id int64, status transcoding.Status) error {
	conn := db.GetConnection()
	if conn == nil {
		return fmt.Errorf("db conn is nil")
	}

	query := `
	UPDATE transcoding_jobs SET status = ? WHERE id = ?
	`

	result, err := conn.ExecContext(context.Background(), query, status, id)
	if err != nil {
		return fmt.Errorf("error updating transcoding job status: %v", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if n == 0 {
		return fmt.Errorf("error updating transcoding job status: no rows affected")
	}

	return nil
}
