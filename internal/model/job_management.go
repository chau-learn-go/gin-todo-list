package model

import "errors"

type JobManagement struct {
	jobs map[string]*Job
}

func NewJobManagement() *JobManagement {
	return &JobManagement{
		jobs: make(map[string]*Job),
	}
}

func (j *JobManagement) Add(job *Job) error {
	if _, found := j.jobs[job.Name]; found {
		return errors.New("job already exists")
	}
	j.jobs[job.Name] = job
	return nil
}

func (j *JobManagement) Get(name string) (*Job, error) {
	if job, found := j.jobs[name]; found {
		return job, nil
	}
	return nil, errors.New("job not found")
}

func (j *JobManagement) GetAll() []*Job {
	var jobs []*Job
	for _, job := range j.jobs {
		jobs = append(jobs, job)
	}
	return jobs
}

func (j *JobManagement) Delete(name string) error {
	if _, found := j.jobs[name]; !found {
		return errors.New("job not found")
	}
	delete(j.jobs, name)
	return nil
}

func (j *JobManagement) Doing(jobName string) error {
	job, found := j.jobs[jobName]
	if !found {
		return errors.New("job not found")
	}
	job.Status = JobDoing
	return nil
}

func (j *JobManagement) Done(jobName string) error {
	job, found := j.jobs[jobName]
	if !found {
		return errors.New("job not found")
	}
	job.Status = JobDone
	return nil
}
