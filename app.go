package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
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

func (a *App) GetData(data ytData) {

	client := youtube.Client{}

	video, err := client.GetVideo(data.Url)
	if err != nil {
		fmt.Println("Failed to get the URL", err)
		return
	}

	if data.StartHH == "" {
		data.StartHH = "00"
	}
	if data.StartMM == "" {
		data.StartMM = "00"
	}
	if data.StartSS == "" {
		data.StartSS = "00"
	}
	if data.Name == "" {
		data.Name = correctFilename(video.Title)
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

	downloadAndTrim(data)
	fmt.Println(data)
}

func downloadAndTrim(data ytData) {
	fmt.Println(data)
	client := youtube.Client{}

	video, err := client.GetVideo(data.Url)
	if err != nil {
		fmt.Println("Failed to get the URL", err)
		return
	}

	//AUDIO FORMAT ONLY
	var bestAudio *youtube.Format
	maxBitrate := 0
	for _, f := range video.Formats {
		if f.AudioChannels > 0 && f.QualityLabel == "" {
			if f.Bitrate > maxBitrate {
				bestAudio = &f
				maxBitrate = f.Bitrate
			}
		}
	}

	if bestAudio == nil {
		fmt.Println("No suitable audio only format found")
	}

	stream, _, err := client.GetStream(video, bestAudio)
	if err != nil {
		fmt.Println("Failed to get Audio Stream:", err)
		return
	}

	filename := correctFilename(video.Title) + ".webm"
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed creating file:", err)
		return
	}
	defer outFile.Close()

	fmt.Println("Downloading audio...")
	_, err = io.Copy(outFile, stream)
	if err != nil {
		fmt.Println("Failed to write audio:", err)
		return
	}
	fmt.Println("Audio downloaded")

	// trim audio with ffmpeg
	outputFile := correctFilename(data.Name) + ".webm"

	ss := data.StartHH + ":" + data.StartMM + ":" + data.StartSS
	to := data.EndHH + ":" + data.EndMM + ":" + data.EndSS

	err = trimAudio(filename, outputFile, ss, to)
	if err != nil {
		fmt.Println("Failed to trim the audio:", err)
		return
	}

	err = os.Remove(outFile.Name())
	if err != nil {
		fmt.Println("Failed to delete the full audio file, please delete it manually:", err)
		return
	}

	fmt.Printf("Audio file saved as %s\n", outputFile)
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
		output,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getVideoDuration(url string) (string, string, string, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return "", "", "", fmt.Errorf("Failed to get video: %w", err)
	}

	hh, mm, ss := formatDurationToHHMMSS(video.Duration)
	return hh, mm, ss, nil
}

func formatDurationToHHMMSS(d time.Duration) (hh, mm, ss string) {
	totalSecond := int(d.Seconds())
	h := totalSecond / 3600
	m := (totalSecond % 3600) / 60
	s := totalSecond % 60

	return fmt.Sprintf("%02d", h), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", s)
}
