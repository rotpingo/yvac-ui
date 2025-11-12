package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type ytData struct {
	Url     string
	StartHH string
	StartMM string
	StartSS string
	EndHH   string
	EndMM   string
	EndSS   string
	Name    string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DownloadAndTrim(data ytData) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Validate URL
	if !strings.Contains(data.Url, "youtube.com") && !strings.Contains(data.Url, "youtu.be") {
		fmt.Println("Invalid YouTube URL")
		return
	}

	checkData(&data)

	tmpDir, err := os.MkdirTemp("", "ytclip-*")
	if err != nil {
		fmt.Println("Failed to create temp dir:", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// Temp files
	fullAudio := filepath.Join(tmpDir, "full.opus")
	finalAudio := correctFilename(data.Name) + ".opus"

	// Build yt-dlp command
	fmt.Println("Downloading autio with yt-dlp")
	ytCmd := exec.CommandContext(ctx, "yt-dlp",
		"--extract-audio",
		"--audio-format", "opus", // or "mp3" if you have libmp3lame
		"--audio-quality", "0", // Best quality
		"--no-playlist",
		"--output", fullAudio,
		"--user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"--referer", "https://www.youtube.com/",
		"--retries", "3",
		data.Url,
	)

	ytCmd.Stdout = os.Stdout
	ytCmd.Stderr = os.Stderr

	if err := ytCmd.Run(); err != nil {
		fmt.Println("yt-dlp failed:", err)
		return
	}

	// Verify file exists
	if _, err := os.Stat(fullAudio); os.IsNotExist(err) {
		fmt.Println("yt-dlp did not create the file:", fullAudio)
		return
	}

	// if err := ytCmd.Run(); err != nil {
	// 	fmt.Println("yt-dlp dit not create the file :", fullAudio)
	// 	return
	// }

	// Trim with FFmpeg
	ss := fmt.Sprintf("%s:%s:%s", data.StartHH, data.StartMM, data.StartSS)
	to := fmt.Sprintf("%s:%s:%s", data.EndHH, data.EndMM, data.EndSS)

	if err := trimAudio(fullAudio, finalAudio, ss, to); err != nil {
		fmt.Println("Trim failed:", err)
		return
	}

	// Cleanup
	fmt.Printf("Saved: %s\n", finalAudio)
}

func checkData(data *ytData) {

	if data.StartHH == "" {
		data.StartHH = "00"
	}
	if data.StartMM == "" {
		data.StartMM = "00"
	}
	if data.StartSS == "" {
		data.StartSS = "00"
	}

	if data.EndHH == "" || data.EndMM == "" || data.EndSS == "" {
		hh, mm, ss, err := getVideoDuration(data.Url)
		if err != nil {
			return
		}

		if data.EndHH == "" {
			if hh == "" {
				data.EndHH = "00"
			} else {
				data.EndHH = hh
			}
		}

		if data.EndMM == "" {
			if hh == "" {
				data.EndMM = "00"
			} else {
				data.EndMM = mm
			}
		}

		if data.EndSS == "" {
			if hh == "" {
				data.EndSS = "00"
			} else {
				data.EndSS = ss
			}
		}
	}
}

func correctFilename(name string) string {
	illegal := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, ch := range illegal {
		name = strings.ReplaceAll(name, ch, "_")
	}

	return name
}

func trimAudio(input, output, start, end string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", input,
		"-ss", start,
		"-to", end,
		"-c:a", "libopus",
		"-y",
		output,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getVideoDuration(url string) (hh, mm, ss string, err error) {

	out, err := exec.Command("yt-dlp", "--get-duration", url).Output()
	if err != nil {
		return "", "", "", err
	}
	dur := strings.TrimSpace(string(out))

	parts := strings.Split(dur, ":")

	switch len(parts) {
	case 2: // MM:SS
		return "00", parts[0], parts[1], nil
	case 3: // HH:MM:SS
		return parts[0], parts[1], parts[2], nil
	default:
		return "00", "00", "00", fmt.Errorf("unexpected duration format: %s", dur)
	}
}

// func formatDurationToHHMMSS(d time.Duration) (hh, mm, ss string) {
// 	totalSecond := int(d.Seconds())
// 	h := totalSecond / 3600
// 	m := (totalSecond % 3600) / 60
// 	s := totalSecond % 60
//
// 	return fmt.Sprintf("%02d", h), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", s)
// }
