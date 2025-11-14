package ffmpeg

import (
	"encoding/json"
	"os/exec"
)

type ProbeResult struct {
	Format struct {
		Filename   string `json:"filename"`
		Duration   string `json:"duration"`
		Size       string `json:"size"`
		FormatName string `json:"format_name"`
	} `json:"format"`

	Streams []struct {
		CodecName string `json:"codec_name"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
	} `json:"streams"`
}

func ProbeVideo(filePath string) (*ProbeResult, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		filePath,
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result ProbeResult
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
