package models

import "github.com/google/uuid"

type Status string

type Video struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Filename   string  `json:"filename"`
	Path       string  `json:"path"`
	Duration   float64 `json:"duration"`
	Size       int64   `json:"size"`
	Format     string  `json:"format"`
	Resolution string  `json:"resolution"`
	Status     Status  `json:"status"`
}

const (
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusTranscoded Status = "transcoded"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

func NewVideo(title, filename, path, format, resolution string, duration float64, size int64, status Status) *Video {
	return &Video{
		ID:         uuid.New().String(),
		Title:      title,
		Filename:   filename,
		Path:       path,
		Duration:   duration,
		Size:       size,
		Format:     format,
		Resolution: resolution,
		Status:     status,
	}
}
