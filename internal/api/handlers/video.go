package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yanisssssse/vidego/internal/ffmpeg"
	"github.com/Yanisssssse/vidego/internal/storage"
	"github.com/google/uuid"
)

type VideoHandlers struct {
	storage storage.Storage
}

func NewVideoHandlers(storage storage.Storage) *VideoHandlers {
	return &VideoHandlers{storage: storage}
}

func (h *VideoHandlers) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 30)

	f, header, err := r.FormFile("file")
	if f == nil {
		http.Error(w, "Missing file", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	videoBytes, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	p, err := h.probeFromBytes(videoBytes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read video metadata: %v", err), http.StatusBadRequest)
		return
	}

	videoID := uuid.New().String()
	videoKey := filepath.Join("video", videoID)

	err = h.storage.Write(r.Context(), videoKey, videoBytes)
	if err != nil {
		http.Error(w, "Failed to save video", http.StatusInternalServerError)
		return
	}

	metadataBytes, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Failed to encode metadata", http.StatusInternalServerError)
		return
	}

	metadataKey := filepath.Join("metadata", videoID+"_metadata.json")
	err = h.storage.Write(r.Context(), metadataKey, metadataBytes)
	if err != nil {
		http.Error(w, "Failed to save metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"success","id":"%s","filename":"%s"}`, videoID, header.Filename)
}

func (h *VideoHandlers) probeFromBytes(data []byte) (*ffmpeg.ProbeResult, error) {
	tempFile, err := os.CreateTemp("", "vidego-probe-*.tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	_, err = tempFile.Write(data)
	if err != nil {
		tempFile.Close()
		return nil, fmt.Errorf("failed to write to temp file: %w", err)
	}
	tempFile.Close()

	p, err := ffmpeg.ProbeVideo(tempPath)
	if err != nil {
		return nil, fmt.Errorf("ffprobe failed: %w", err)
	}

	return p, nil
}
