package main

// https://wails.io/docs/guides/application-development/#binding-methods

import (
  "context"
  // "fmt"
)

// App struct
type App struct {
  ctx context.Context
}

// NewApp creates a new App application struct
func LibroLoversApp() *App {
  return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
  a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
}

