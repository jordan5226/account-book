package acctbook

import (
	"account-book/lib/pgdb/schema"
	"account-book/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *AcctBook) GetEntries(c *gin.Context) {
	// Get Parameter
	dateStr := c.Param("time")
	uid := c.Param("uid")
	date, err := time.Parse("2006-01-02", dateStr) // According to RFC3339, must key in "2006-01-02" as format.
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Query Data
	entries, err := a.mdlEntry.Get(date, uid, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Response Data
	c.JSON(http.StatusOK, middleware.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Entries []schema.Entry `json:"entries"`
		}{
			Entries: entries,
		},
	})
}

func (a *AcctBook) CreateEntry(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Entry)
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
	_, err = a.mdlEntry.Add(input, nil)
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

func (a *AcctBook) UpdateEntry(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Entry)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Update Data
	err := a.mdlEntry.Update(input, nil)
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

func (a *AcctBook) DeleteEntry(c *gin.Context) {
	// Get Parameter
	uid := c.Param("uid")
	id := c.Param("id")

	// Delete row from DB
	err := a.mdlEntry.Delete(uid, id, nil)
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
