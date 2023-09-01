package main

import (
	// "fmt"
	// "log"
	"fmt"
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

		// send mail after update to notify user
		subject := fmt.Sprintf("Cherava: Content update on workflow -  %s", w.Name)
		body := fmt.Sprintf("Your workflow ${w.name} had the following update received\n\t\n\t\t %v \t    ", w.Data)
		err := sendMail(w.Email, subject, body)
		if err != nil {
			fmt.Printf("Error sending mail %s", err)
		}
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
