package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
		"github.com/wailsapp/wails/v2/pkg/runtime"

)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte
func startHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	server := &http.Server{
		Addr:           ":4000", // Change this port if needed
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Starting HTTP server on port 4000...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	
	go startHTTPServer()
	    AppMenu := menu.NewMenu()
    FileMenu := AppMenu.AddSubmenu("File")
    // FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
    // FileMenu.AddSeparator()
    FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
        runtime.Quit(app.ctx)
    })

    // if runtime.GOOS == "darwin" {
    // AppMenu.Append(menu.EditMenu())  // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
    // }

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "vanilla",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        true,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		Menu:              AppMenu,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Fullscreen,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "vanilla",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
