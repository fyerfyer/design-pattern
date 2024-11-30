package updater

import (
	"time"
)

type Updater interface {
	Update()
}

type Viewer interface {
	OutputInPlainText() string
	Output() map[string]string
}

type ScheduledUpdater struct {
	updater Updater
	period  time.Duration
}

func NewScheduledUpdater(updater Updater, periodInSeconds int) *ScheduledUpdater {
	return &ScheduledUpdater{
		updater: updater,
		period:  time.Duration(periodInSeconds) * time.Second,
	}
}

func (s *ScheduledUpdater) Run() {
	ticker := time.NewTicker(s.period)
	defer ticker.Stop()

	for range ticker.C {
		s.updater.Update()
	}
}
