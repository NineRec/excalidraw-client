package main

import (
	"embed"
	"fmt"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	m := menu.NewMenu()
	m.Append(menu.AppMenu())

	about := m.AddSubmenu("File")
	about.AddText("Save To...", nil, saveToDisck)
	about.AddText("World", nil, nil)

	if runtime.GOOS == "darwin" {
		// on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
		m.Append(menu.EditMenu())
		m.Append(menu.WindowMenu())
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "excalidraw",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Menu: m,
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "Excalidraw Client",
				Message: "@2023 CodePanic",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func saveToDisck(cb *menu.CallbackData) {
	fmt.Println("Export file")
}
