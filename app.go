package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

type Greeting struct {
	Name string `json:"name"`
}


// Greet returns a greeting for the given name
func (a *App) Greet(jsonStr string) string {
		var greeting Greeting

	// Unmarshal JSON string into the Greeting struct
	err := json.Unmarshal([]byte(jsonStr), &greeting)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return fmt.Sprintf("Hello %s, It's show time!", greeting.Name)
}

func (a *App) Copy(jsonStr string) string {
		var greeting Greeting

	// Unmarshal JSON string into the Greeting struct
	err := json.Unmarshal([]byte(jsonStr), &greeting)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return fmt.Sprintf("Hello %s, It's show time!", greeting.Name)
}

