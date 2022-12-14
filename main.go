package main

// https://wails.io/docs/guides/application-development/#application-setup

import (
  "embed"
  "github.com/wailsapp/wails/v2"
  "github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend
var assets embed.FS

func main() {
  // Create an instance of the app structure
  app := LibroLoversApp()

  // Create application with options
  err := wails.Run(&options.App{
    Title:            "Libro Lovers",
    Width:            800,
    Height:           600,
    Assets:           assets,
    BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
    OnStartup:        app.startup,
    OnShutdown:       app.shutdown,
    // Menu:             app.menu(),
    Bind: []interface{}{
      app,
    },
  })

  if err != nil {
    println("Error:", err.Error())
  }
}
