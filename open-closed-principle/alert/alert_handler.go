package alert

// the interface for all alert handlers.
type AlertHandler interface {
	Handle(info ApiStatInfo)
}

// the alert rule configuration.
type AlertRule struct {
}

// the notification system.
type Notification struct {
}

// BaseHandler is a common structure for handlers.
type BaseHandler struct {
	alertRule    *AlertRule
	notification *Notification
}

func NewBaseHandler(alertRule *AlertRule, notification *Notification) *BaseHandler {
	return &BaseHandler{
		alertRule:    alertRule,
		notification: notification,
	}
}

// handles alerts related to TPS (Transactions Per Second).
type TpsAlertHandler struct {
	*BaseHandler
}

func NewTpsAlertHandler(alertRule *AlertRule, notification *Notification) *TpsAlertHandler {
	return &TpsAlertHandler{
		BaseHandler: NewBaseHandler(alertRule, notification),
	}
}

func (h *TpsAlertHandler) Handle(info ApiStatInfo) {
	// Handle logic for TPS alert
}

// handles alerts related to errors.
type ErrorAlertHandler struct {
	*BaseHandler
}

func NewErrorAlertHandler(alertRule *AlertRule, notification *Notification) *ErrorAlertHandler {
	return &ErrorAlertHandler{
		BaseHandler: NewBaseHandler(alertRule, notification),
	}
}

func (h *ErrorAlertHandler) Handle(info ApiStatInfo) {
	// Handle logic for error alert
}

// handles alerts related to timeouts.
type TimeoutAlertHandler struct {
	*BaseHandler
}

func NewTimeoutAlertHandler(alertRule *AlertRule, notification *Notification) *TimeoutAlertHandler {
	return &TimeoutAlertHandler{
		BaseHandler: NewBaseHandler(alertRule, notification),
	}
}

func (h *TimeoutAlertHandler) Handle(info ApiStatInfo) {
	// Handle logic for timeout alert
}
