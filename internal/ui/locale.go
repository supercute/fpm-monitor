package ui

// Locale Structure
type Locale struct {
	ProcessesTitle      string
	ActiveProcesses     string
	IdleProcesses       string
	PoolLoadTitle       string
	RPSTitle            string
	AvgDurationTitle    string
	SlowRequestsTitle   string
	SlowRequestsText    string
	TotalProcesses      string
	QueueLength         string
	MaxChildren         string
	Uptime              string
	AcceptedConnections string
	Exit                string
}

var Locales = map[string]Locale{
	"en": {
		ProcessesTitle:      "Active/Idle Processes",
		ActiveProcesses:     "Active",
		IdleProcesses:       "Idle",
		PoolLoadTitle:       "Pool Load %",
		RPSTitle:            "Requests/sec",
		AvgDurationTitle:    "Avg Request Duration (ms)",
		SlowRequestsTitle:   "Slow Requests",
		SlowRequestsText:    "slow requests",
		TotalProcesses:      "Total Processes",
		QueueLength:         "Queue Length",
		MaxChildren:         "Max Children Reached",
		Uptime:              "Uptime",
		AcceptedConnections: "Accepted Connections",
		Exit:                "Press 'q' or Ctrl+C to quit",
	},
	"ru": {
		ProcessesTitle:      "Активные/Свободные Процессы",
		ActiveProcesses:     "Активные",
		IdleProcesses:       "Свободные",
		PoolLoadTitle:       "Загрузка Пула %",
		RPSTitle:            "Запросов/сек",
		AvgDurationTitle:    "Среднее Время Запроса (мс)",
		SlowRequestsTitle:   "Медленные Запросы",
		SlowRequestsText:    "медленных запросов",
		TotalProcesses:      "Всего Процессов",
		QueueLength:         "Длина Очереди",
		MaxChildren:         "Достигнут Макс. Лимит",
		Uptime:              "Время Работы",
		AcceptedConnections: "Принято Соединений",
		Exit:                "Нажмите 'q' или Ctrl+C для выхода",
	},
}
