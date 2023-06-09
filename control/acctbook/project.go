package acctbook

import (
	"account-book/lib/pgdb/schema"
	"account-book/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *AcctBook) GetProjects(c *gin.Context) {
	// Get Parameter
	uid := c.Param("uid")

	// Query Data
	prjs, err := a.mdlPrj.Get(uid, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	if len(prjs) == 0 {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: "Query project data failed!",
		})
		return
	}

	// Response Data
	c.JSON(http.StatusOK, middleware.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Prjs []schema.Project `json:"prjs"`
		}{
			Prjs: prjs,
		},
	})
}

func (a *AcctBook) CreateProject(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Project)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Generate Data
	_id, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, middleware.HttpFailResponse{
			Status:  "fail",
			Message: "Fail to generate uuid!",
		})
		return
	}

	input.Id = _id.String()

	// Write input data to DB
	_, err = a.mdlPrj.Add(input, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Success Response
	c.JSON(http.StatusOK, middleware.HttpFailResponse{
		Status:  "success",
		Message: "success",
	})
}

func (a *AcctBook) UpdateProject(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Project)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Update Data
	err := a.mdlPrj.Update(input, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Success Response
	c.JSON(http.StatusOK, middleware.HttpFailResponse{
		Status:  "success",
		Message: "success",
	})
}

func (a *AcctBook) DeleteProject(c *gin.Context) {
	// Get Parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: "Invalid Data!",
		})
		return
	}

	// Delete row from DB
	err := a.mdlPrj.Delete(id, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Success Response
	c.JSON(http.StatusOK, middleware.HttpFailResponse{
		Status:  "success",
		Message: "success",
	})
}
