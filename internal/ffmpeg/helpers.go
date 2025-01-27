package ffmpeg

import (
	"bytes"
	"os/exec"
)

func TranscodeToJPEG(b []byte) ([]byte, error) {
	cmd := exec.Command("ffmpeg", "-hide_banner", "-i", "-", "-f", "mjpeg", "-")
	cmd.Stdin = bytes.NewBuffer(b)
	return cmd.Output()
}

func TranscodeToJPG(b []byte) ([]byte, error) {
	cmd := exec.Command("ffmpeg", "-hide_banner", "-i", "-", "-f", "image2", "-")
	cmd.Stdin = bytes.NewBuffer(b)
	return cmd.Output()
}
