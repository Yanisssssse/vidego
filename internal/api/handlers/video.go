package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yanisssssse/vidego/internal/ffmpeg"
	"github.com/joho/godotenv"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 30)

	err := godotenv.Load(".env")
	if err != nil {
		http.Error(w, "Cannot load env", http.StatusInternalServerError)
		return
	}

	f, header, err := r.FormFile("file")
	if f == nil {
		http.Error(w, "Missing file", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	uploadDir := os.Getenv("UPDIR")
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		http.Error(w, "Cannot create uplaods directory", http.StatusInternalServerError)
		return
	}

	dest, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, "Writing error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	_, err = dest.ReadFrom(f)
	if err != nil {
		http.Error(w, "Transfer error", http.StatusInternalServerError)
		return
	}

	p, err := ffmpeg.ProbeVideo(filepath.Join("data/uploads/", header.Filename))
	if err != nil {
		http.Error(w, "Cannot read video metadata", http.StatusBadRequest)
		return
	}
	fmt.Println(p)

	_, err = io.Copy(dest, f)
	if err != nil {
		http.Error(w, "Error while copying file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"success","filename":"%s"}`, header.Filename)
}
