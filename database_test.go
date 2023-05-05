package main

import (
	"testing"
	"time"
)

func TestCreateNewWorkflow(t *testing.T) {

	if err := Connect(); err != nil {
		t.Error("Error connecting to db")
	}

	// getWorkflowsByUserID("872fd941-af4e-4f8d-8a75-ee25347b1038")

	// getWorkflowByID("fdfa0697-9772-4cbf-8a93-367179bb6025")

	workflow := Workflow{
		ID:          "fdfa0827-9772-4cbf-8a93-367179bb6090",
		User:        "872fd941-af4e-4f8d-8a75-ee25347b1038",
		Data:        "6",
		Selector:    "#repo-stars-counter-star",
		Cron:        "* * * * *",
		LastUpdated: time.Now().UnixMilli(),
		URL:         "https://github.com/Roshan-R/termv-rs",
		Name:        "Termv Stars",
		Email:       "example@example.com",
	}

	err := createNewWorkflow(workflow)

	if err != nil {
		t.Error("Error creating workflow")
	}

}
