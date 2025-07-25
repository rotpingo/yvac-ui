package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type ytData struct {
	url     string
	startHH string
	startMM string
	startSS string
	endHH   string
	endMM   string
	endSS   string
	name    string
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

func (a *App) GetData(data *ytData) error {
	fmt.Println("Form files:", data.url, data.startHH, data.startMM, data.startSS, data.endHH, data.endMM, data.endSS, data.name)

	if data.startHH == "" {
		data.startHH = "00"
	}

	if data.startMM == "" {
		data.startMM = "00"
	}

	if data.startSS == "" {
		data.startSS = "00"
	}

	return nil
}
