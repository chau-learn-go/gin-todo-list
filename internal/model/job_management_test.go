package model

import "testing"

func TestJobManagement(t *testing.T) {
	// Test init job management
	jobManagement := newJobManagement()

	// Test add new job
	_ = jobManagement.Add(&Job{
		Name: "job1",
	})

	// Test add exist Job
	err := jobManagement.Add(&Job{
		Name: "job1",
	})
	if err == nil {
		t.Error("job should not be added to job list")
	}

	// Add 1 new job
	_ = jobManagement.Add(&Job{
		Name: "job2",
	})

	// Test do job2
	_ = jobManagement.Doing("job2")

	job2Status, err := jobManagement.Get("job2")
	if err != nil {
		t.Error(err)
	}
	if job2Status.Status != JobDoing {
		t.Error("job2 status should be JobDone")
	}

	// Test done job2
	_ = jobManagement.Done("job2")
	job2Status, err = jobManagement.Get("job2")
	if err != nil {
		t.Error(err)
	}
	if job2Status.Status != JobDone {
		t.Error("job2 status should be JobDone")
	}
}
