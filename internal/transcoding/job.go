package transcoding

import "time"

type Job struct {
	ID           int
	VideoID      int
	Status       Status
	InputKey     string
	OutputKey    string
	CreatedAt    time.Time
	StartedAt    time.Time
	CompletedAt  time.Time
	ErrorMessage string
}

func NewEmptyJob() *Job {
	return &Job{
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}

func NewJob(videoID int, inputKey string, outputKey string) *Job {
	return &Job{
		VideoID:   videoID,
		Status:    StatusPending,
		InputKey:  inputKey,
		OutputKey: outputKey,
		CreatedAt: time.Now(),
	}
}
