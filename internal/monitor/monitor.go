package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FPMStatus struct {
	Pool               string    `json:"pool"`
	ProcessManager     string    `json:"process manager"`
	StartTime          int64     `json:"start time"`
	StartSince         int       `json:"start since"`
	AcceptedConn       int       `json:"accepted conn"`
	ListenQueue        int       `json:"listen queue"`
	MaxListenQueue     int       `json:"max listen queue"`
	ListenQueueLen     int       `json:"listen queue len"`
	IdleProcesses      int       `json:"idle processes"`
	ActiveProcesses    int       `json:"active processes"`
	TotalProcesses     int       `json:"total processes"`
	MaxActiveProcesses int       `json:"max active processes"`
	MaxChildrenReached int       `json:"max children reached"`
	SlowRequests       int       `json:"slow requests"`
	Processes          []Process `json:"processes"`
}

type Process struct {
	PID               int     `json:"pid"`
	State             string  `json:"state"`
	StartTime         int64   `json:"start time"`
	StartSince        int     `json:"start since"`
	Requests          int     `json:"requests"`
	RequestDuration   int     `json:"request duration"`
	RequestMethod     string  `json:"request method"`
	RequestURI        string  `json:"request uri"`
	ContentLength     int     `json:"content length"`
	User              string  `json:"user"`
	Script            string  `json:"script"`
	LastRequestCPU    float64 `json:"last request cpu"`
	LastRequestMemory int     `json:"last request memory"`
}

type Monitor struct {
	statusURL  string
	History    []FPMStatus
	maxHistory int
}

func NewMonitor(statusURL string) *Monitor {
	return &Monitor{
		statusURL:  statusURL,
		History:    make([]FPMStatus, 0),
		maxHistory: 60,
	}
}

func (m *Monitor) FetchStatus() (*FPMStatus, error) {
	resp, err := http.Get(m.statusURL + "?json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status FPMStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func (m *Monitor) AddToHistory(status *FPMStatus) {
	m.History = append(m.History, *status)
	if len(m.History) > m.maxHistory {
		m.History = m.History[1:]
	}
}

func (m *Monitor) CalculateRPS() float64 {
	if len(m.History) < 2 {
		return 0
	}

	current := m.History[len(m.History)-1]
	previous := m.History[len(m.History)-2]

	connDiff := current.AcceptedConn - previous.AcceptedConn
	return float64(connDiff)
}

func (m *Monitor) GetAvgRequestDuration() float64 {
	if len(m.History) == 0 {
		return 0
	}

	status := m.History[len(m.History)-1]
	totalDuration := 0
	activeCount := 0

	for _, process := range status.Processes {
		if process.State == "Running" && process.RequestDuration > 0 {
			totalDuration += process.RequestDuration
			activeCount++
		}
	}

	if activeCount > 0 {
		return float64(totalDuration) / float64(activeCount) / 1000.0
	}
	return 0
}

func (m *Monitor) FormatUptime(seconds int) string {
	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, secs)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, secs)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, secs)
	}
	return fmt.Sprintf("%ds", secs)
}
