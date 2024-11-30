package alert

// represents the alert system.
type Alert struct {
	alertHandlers []AlertHandler
}

// adds a handler to the alert system.
func (a *Alert) AddAlertHandler(handler AlertHandler) {
	a.alertHandlers = append(a.alertHandlers, handler)
}

// checks the API statistics and triggers alerts.
// we use info to make modify the arguments more easily.
func (a *Alert) Check(info ApiStatInfo) {
	for _, handler := range a.alertHandlers {
		handler.Handle(info)
	}
}
