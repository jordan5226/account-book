package acctbook

import (
	"account-book/lib/pgdb/schema"
	"account-book/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *AcctBook) GetStores(c *gin.Context) {
	// Get Parameter
	uid := c.Param("uid")

	// Query Data
	stores, err := a.mdlStore.Get(uid)
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
			Stores []schema.Store `json:"stores"`
		}{
			Stores: stores,
		},
	})
}

func (a *AcctBook) CreateStore(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Store)
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
	_, err = a.mdlStore.Add(input)
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

func (a *AcctBook) UpdateStore(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.Store)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Update Data
	err := a.mdlStore.Update(input)
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

func (a *AcctBook) DeleteStore(c *gin.Context) {
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
	err := a.mdlStore.Delete(id)
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
