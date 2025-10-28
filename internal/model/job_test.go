package model

import "testing"

func TestJob(t *testing.T) {
	// Test create job
	job := CreateJob("Job A")
	if job.Status != JobToDo {
		t.Error("job status should be JobToDo")
	}

	// Test set job to doing
	job.SetDoing()
	if job.Status != JobDoing {
		t.Error("job status should be JobDoing")
	}

	// Test set job to done
	job.SetDone()
	if job.Status != JobDone {
		t.Error("job status should be JobDone")
	}
}
