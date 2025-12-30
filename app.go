package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx context.Context
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello my nigga %s, It's show time!", name)
}

// read a file
func (a *App) ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

// save a text file
func (a *App) SaveTextFile(content, title string) error {
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Save text file",
		DefaultFilename: title + ".txt",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Text Files",
				Pattern:     "*.txt",
			},
		},
	})

	// User cancelled dialog
	if err != nil || path == "" {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}
