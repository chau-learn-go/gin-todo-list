package model

// Job Status
const (
	JobToDo  = "todo"
	JobDoing = "doing"
	JobDone  = "done"
)

type Job struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// CreateJob Create a job with name
func CreateJob(name string) *Job {
	return &Job{
		Name:   name,
		Status: JobToDo,
	}
}

// SetDoing Set job status to "doing"
func (j *Job) SetDoing() *Job {
	j.Status = JobDoing
	return j
}

// SetDone Set job status to "done"
func (j *Job) SetDone() *Job {
	j.Status = JobDone
	return j
}
