package api

import (
	"gin-todo-list/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jobManagement = model.GetJobManagementInstance()

func InitJob() {
	_ = jobManagement.Add(&model.Job{Name: "job1", Status: model.JobToDo})
	_ = jobManagement.Add(&model.Job{Name: "job2", Status: model.JobDoing})
	_ = jobManagement.Add(&model.Job{Name: "job3", Status: model.JobToDo})
	_ = jobManagement.Add(&model.Job{Name: "job4", Status: model.JobDoing})
	_ = jobManagement.Add(&model.Job{Name: "job5", Status: model.JobToDo})
}

func GetJobList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"items": jobManagement.GetAll()})
}

func GetJob(c *gin.Context) {
	job, err := jobManagement.Get(c.Param("name"))
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"item": job})
}

func PostJob(c *gin.Context) {
	var job model.Job
	if err := c.ShouldBind(&job); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	err := jobManagement.Add(&job)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "job added"})
}

func DoingJob(c *gin.Context) {
	jobName := c.Param("name")
	if err := jobManagement.Doing(jobName); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "job doing"})
}

func DoneJob(c *gin.Context) {
	jobName := c.Param("name")
	if err := jobManagement.Done(jobName); err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"message": "job done"})
}

func DeleteJob(c *gin.Context) {
	jobName := c.Param("name")
	if err := jobManagement.Delete(jobName); err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"message": "job deleted"})
}
