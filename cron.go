package main

import (
	// "fmt"
	// "log"
	"time"

	"github.com/go-co-op/gocron"
)

var s *gocron.Scheduler

func setupCron() error {
	s = gocron.NewScheduler(time.UTC)
	setCronForAllWorkflows()
	s.StartAsync()
	return nil
}

func updateIfChanged(w Workflow) {
	// log.Println(fmt.Sprintf("checking for update with workflow id: %s", w.ID))
	d, _ := scrape(w.URL, w.Selector)
	if d != w.Data {
		updateWorkflow(w, d)
	}
}

func setCronForAllWorkflows() {
	ws, _ := getAllWorkflows()
	for _, w := range ws.Workflows {
		setCronForWorkflow(w)
	}
}

func setCronForWorkflow(w Workflow) {
	fn1 := func() { updateIfChanged(w) }
	s.Cron(w.Cron).Do(fn1)
}
