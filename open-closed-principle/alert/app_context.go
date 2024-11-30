package alert

type ApplicationContext struct {
	alertRule    *AlertRule
	notification *Notification
	alert        *Alert
}

func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{}
}

func (ctx *ApplicationContext) InitializeBeans() {
	ctx.alertRule = &AlertRule{}
	ctx.notification = &Notification{}
	ctx.alert = &Alert{}

	ctx.alert.AddAlertHandler(NewTpsAlertHandler(ctx.alertRule, ctx.notification))
	ctx.alert.AddAlertHandler(NewErrorAlertHandler(ctx.alertRule, ctx.notification))
	// Register the new handler
	ctx.alert.AddAlertHandler(NewTimeoutAlertHandler(ctx.alertRule, ctx.notification))
}

func (ctx *ApplicationContext) GetAlert() *Alert {
	return ctx.alert
}
