package scheduler

import (
	"context"
	"i3/pkg/logger"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func New() *Scheduler {
	cronLogger := logger.NewCronLogger()

	var opt []cron.Option

	tz, _ := time.LoadLocation("Asia/Jakarta")

	opt = append(opt,
		cron.WithLogger(cronLogger),
		cron.WithLocation(tz),
		cron.WithChain(cron.Recover(cronLogger)),
	)

	c := cron.New(opt...)

	return &Scheduler{
		cron: c,
	}
}

func (s *Scheduler) ScheduleFunc(spec string, cmd func()) (cron.EntryID, error) {
	return s.cron.AddFunc(spec, cmd)
}

func (s *Scheduler) ScheduleJob(duration time.Duration, job cron.Job) cron.EntryID {
	return s.cron.Schedule(cron.Every(duration), job)
}

func (s *Scheduler) Remove(id cron.EntryID) {
	s.cron.Remove(id)
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() context.Context {
	return s.cron.Stop()
}
