package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/helpers"
	"github.com/PutraFajarF/backend-ats-app-cap/service"
	"github.com/gin-gonic/gin"
)

type JobsAppController interface {
	GetAllJobsApplicant(ctx *gin.Context)
	ApplicantGetJobsByID(ctx *gin.Context)
}

type jobsAppController struct {
	jobsService service.JobsService
}

func NewJobsAppController(jobServ service.JobsService) JobsAppController {
	return &jobsAppController{
		jobsService: jobServ,
	}
}

func (c *jobsAppController) GetAllJobsApplicant(ctx *gin.Context) {
	jobs, err := c.jobsService.AllJobs()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	response := helpers.BuildResponse(true, "Success to get all jobs data", jobs)
	ctx.JSON(http.StatusOK, response)
}

func (c *jobsAppController) ApplicantGetJobsByID(ctx *gin.Context) {
	var jobs entity.Jobs

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed to get id", "No param id were found", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	jobs.ID = id

	job, err := c.jobsService.GetJobByID(int(jobs.ID))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		messError := fmt.Sprintf("failed to get jobs by id")
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to get jobs data by id", job)
	ctx.JSON(http.StatusOK, response)
}
