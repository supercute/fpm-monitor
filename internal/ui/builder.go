package ui

import (
	"fmt"
	termUI "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/supercute/fpm-monitor/internal/monitor"
)

type WidgetBuilder struct {
	Locale Locale
}

type Widgets struct {
	ProcessChart          *widgets.Plot
	PoolGauge             *widgets.Gauge
	RPSGauge              *widgets.Gauge
	AvgDurationParagraph  *widgets.Paragraph
	SlowRequestsParagraph *widgets.Paragraph
	InfoTable             *widgets.Table
	HelpParagraph         *widgets.Paragraph
}

func NewWidgetBuilder(lang string) *WidgetBuilder {
	locale, exists := Locales[lang]
	if !exists {
		locale = Locales["en"]
	}

	return &WidgetBuilder{
		Locale: locale,
	}
}

func (w *WidgetBuilder) CreateWidgets() *Widgets {
	return &Widgets{
		ProcessChart:          w.createProcessChart(),
		PoolGauge:             w.createPoolGauge(),
		RPSGauge:              w.createRPSGauge(),
		AvgDurationParagraph:  w.createAvgDurationParagraph(),
		SlowRequestsParagraph: w.createSlowRequestsParagraph(),
		InfoTable:             w.createInfoTable(),
		HelpParagraph:         w.createHelpParagraph(),
	}
}

func (w *WidgetBuilder) createProcessChart() *widgets.Plot {
	processChart := widgets.NewPlot()
	processChart.Title = w.Locale.ProcessesTitle
	processChart.LineColors[0] = termUI.ColorRed
	processChart.AxesColor = termUI.ColorWhite
	return processChart
}

func (w *WidgetBuilder) createPoolGauge() *widgets.Gauge {
	poolGauge := widgets.NewGauge()
	poolGauge.Title = w.Locale.PoolLoadTitle
	poolGauge.BarColor = termUI.ColorBlue
	return poolGauge
}

func (w *WidgetBuilder) createRPSGauge() *widgets.Gauge {
	rpsGauge := widgets.NewGauge()
	rpsGauge.Title = w.Locale.RPSTitle
	rpsGauge.BarColor = termUI.ColorGreen
	return rpsGauge
}

func (w *WidgetBuilder) createAvgDurationParagraph() *widgets.Paragraph {
	avgDurationParagraph := widgets.NewParagraph()
	avgDurationParagraph.Title = w.Locale.AvgDurationTitle
	return avgDurationParagraph
}

func (w *WidgetBuilder) createSlowRequestsParagraph() *widgets.Paragraph {
	slowRequestsParagraph := widgets.NewParagraph()
	slowRequestsParagraph.Title = w.Locale.SlowRequestsTitle
	return slowRequestsParagraph
}

func (w *WidgetBuilder) createInfoTable() *widgets.Table {
	infoTable := widgets.NewTable()
	infoTable.Title = "Pool Info"
	infoTable.TextStyle = termUI.NewStyle(termUI.ColorWhite)
	infoTable.RowSeparator = false
	return infoTable
}

func (w *WidgetBuilder) createHelpParagraph() *widgets.Paragraph {
	helpParagraph := widgets.NewParagraph()
	helpParagraph.Text = w.Locale.Exit
	return helpParagraph
}

func (w *WidgetBuilder) InitializeEmptyData(widgets *Widgets) {
	widgets.InfoTable.Rows = [][]string{
		{w.Locale.TotalProcesses, "Loading..."},
		{w.Locale.QueueLength, "Loading..."},
		{w.Locale.MaxChildren, "Loading..."},
		{w.Locale.Uptime, "Loading..."},
		{w.Locale.AcceptedConnections, "Loading..."},
		{"Pool", "Loading..."},
	}

	widgets.PoolGauge.Percent = 0
	widgets.PoolGauge.Label = "Loading..."

	widgets.RPSGauge.Percent = 0
	widgets.RPSGauge.Label = "Loading..."

	widgets.AvgDurationParagraph.Text = "Loading..."
	widgets.SlowRequestsParagraph.Text = "Loading..."
}

func (w *WidgetBuilder) SetLayout(widgets *Widgets) {
	widgets.ProcessChart.SetRect(0, 0, 60, 15)
	widgets.PoolGauge.SetRect(60, 0, 120, 8)
	widgets.RPSGauge.SetRect(60, 8, 120, 15)
	widgets.AvgDurationParagraph.SetRect(0, 15, 30, 25)
	widgets.SlowRequestsParagraph.SetRect(30, 15, 60, 25)
	widgets.InfoTable.SetRect(60, 15, 120, 25)
	widgets.HelpParagraph.SetRect(0, 25, 120, 28)
}

func (w *WidgetBuilder) SetResponsiveLayout(widgets *Widgets, width, height int) {
	widgets.ProcessChart.SetRect(0, 0, width*3/5, height*3/5)
	widgets.PoolGauge.SetRect(width*3/5, 0, width, height*2/5)
	widgets.RPSGauge.SetRect(width*3/5, height*2/5, width, height*3/5)
	widgets.AvgDurationParagraph.SetRect(0, height*3/5, width/4, height*4/5)
	widgets.SlowRequestsParagraph.SetRect(width/4, height*3/5, width/2, height*4/5)
	widgets.InfoTable.SetRect(width/2, height*3/5, width, height*4/5)
	widgets.HelpParagraph.SetRect(0, height*4/5, width, height)
}

func (w *WidgetBuilder) UpdateWidgets(widgets *Widgets, mon *monitor.Monitor, status *monitor.FPMStatus) {
	w.updateProcessChart(widgets.ProcessChart, mon)
	w.updatePoolGauge(widgets.PoolGauge, status)
	w.updateRPSGauge(widgets.RPSGauge, mon)
	w.updateAvgDuration(widgets.AvgDurationParagraph, mon)
	w.updateSlowRequests(widgets.SlowRequestsParagraph, status)
	w.updateInfoTable(widgets.InfoTable, mon, status)
}

func (w *WidgetBuilder) updateProcessChart(processChart *widgets.Plot, mon *monitor.Monitor) {
	if len(mon.History) > 1 {
		plotData := make([][]float64, 2)
		plotData[0] = make([]float64, len(mon.History))
		plotData[1] = make([]float64, len(mon.History))

		for i, s := range mon.History {
			plotData[0][i] = float64(s.ActiveProcesses)
			plotData[1][i] = float64(s.IdleProcesses)
		}

		processChart.Data = plotData
	}
}

func (w *WidgetBuilder) updatePoolGauge(poolGauge *widgets.Gauge, status *monitor.FPMStatus) {
	if status.TotalProcesses > 0 {
		poolLoad := int((float64(status.ActiveProcesses) / float64(status.TotalProcesses)) * 100)
		poolGauge.Percent = poolLoad
		poolGauge.Label = fmt.Sprintf("%d/%d (%d%%)", status.ActiveProcesses, status.TotalProcesses, poolLoad)
	}
}

func (w *WidgetBuilder) updateRPSGauge(rpsGauge *widgets.Gauge, mon *monitor.Monitor) {
	rps := mon.CalculateRPS()
	rpsValue := int(rps)
	if rpsValue > 100 {
		rpsValue = 100
	}
	rpsGauge.Percent = rpsValue
	rpsGauge.Label = fmt.Sprintf("%.1f req/s", rps)
}

func (w *WidgetBuilder) updateAvgDuration(avgDurationParagraph *widgets.Paragraph, mon *monitor.Monitor) {
	avgDuration := mon.GetAvgRequestDuration()
	avgDurationParagraph.Text = fmt.Sprintf("%.2f ms", avgDuration)
}

func (w *WidgetBuilder) updateSlowRequests(slowRequestsParagraph *widgets.Paragraph, status *monitor.FPMStatus) {
	slowRequestsParagraph.Text = fmt.Sprintf("%d %s", status.SlowRequests, w.Locale.SlowRequestsText)
}

func (w *WidgetBuilder) updateInfoTable(infoTable *widgets.Table, mon *monitor.Monitor, status *monitor.FPMStatus) {
	infoTable.Rows = [][]string{
		{w.Locale.TotalProcesses, fmt.Sprintf("%d", status.TotalProcesses)},
		{w.Locale.QueueLength, fmt.Sprintf("%d", status.ListenQueueLen)},
		{w.Locale.MaxChildren, fmt.Sprintf("%d", status.MaxChildrenReached)},
		{w.Locale.Uptime, mon.FormatUptime(status.StartSince)},
		{w.Locale.AcceptedConnections, fmt.Sprintf("%d", status.AcceptedConn)},
		{"Pool", status.Pool},
	}
}

func (w *WidgetBuilder) Render(widgets *Widgets) {
	termUI.Render(
		widgets.ProcessChart,
		widgets.PoolGauge,
		widgets.RPSGauge,
		widgets.AvgDurationParagraph,
		widgets.SlowRequestsParagraph,
		widgets.InfoTable,
		widgets.HelpParagraph,
	)
}

func (w *WidgetBuilder) HandleEvents(widgets *Widgets) {
	uiEvents := termUI.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(termUI.Resize)
			w.SetResponsiveLayout(widgets, payload.Width, payload.Height)
			termUI.Clear()
			w.Render(widgets)
		}
	}
}
