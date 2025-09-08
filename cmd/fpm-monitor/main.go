package main

import (
	"flag"
	"log"
	"time"

	termUI "github.com/gizak/termui/v3"
	"github.com/supercute/fpm-monitor/internal/monitor"
	"github.com/supercute/fpm-monitor/internal/ui"
)

func main() {
	var lang = flag.String("lang", "en", "Language: en or ru")
	var statusURL = flag.String("url", "http://localhost/status", "PHP-FPM status URL")
	flag.Parse()

	if err := termUI.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termUI.Close()

	newMonitor := monitor.NewMonitor(*statusURL)
	widgetBuilder := ui.NewWidgetBuilder(*lang)
	widgets := widgetBuilder.CreateWidgets()
	widgetBuilder.SetLayout(widgets)
	widgetBuilder.InitializeEmptyData(widgets)
	widgetBuilder.Render(widgets)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			status, err := newMonitor.FetchStatus()
			if err != nil {
				log.Printf("Error fetching status: %v", err)
				continue
			}

			newMonitor.AddToHistory(status)
			widgetBuilder.UpdateWidgets(widgets, newMonitor, status)
			widgetBuilder.Render(widgets)
		}
	}()

	widgetBuilder.HandleEvents(widgets)
}
