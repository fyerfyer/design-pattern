package alert

type ApiStatInfo struct {
	API               string
	RequestCount      int64
	ErrorCount        int64
	DurationOfSeconds int64
	TimeoutCount      int64
}

func (a *ApiStatInfo) SetTimeoutCount(count int64) {
	a.TimeoutCount = count
}
