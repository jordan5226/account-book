package acctbook

import (
	"account-book/lib/pgdb/schema"
	"account-book/middleware"
	"account-book/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *AcctBook) GetUser(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.User)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	//
	input.Pwd = util.MakeMD5(input.Pwd)

	// Query Data
	user, err := a.mdlUser.Get(input.Uid, input.Pwd, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: "Query user data failed!",
		})
		return
	}

	user[0].Pwd = ""

	// Response Data
	c.JSON(http.StatusOK, middleware.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			User []schema.User `json:"user"`
		}{
			User: user,
		},
	})
}

func (a *AcctBook) CreateUser(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.User)
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
	input.Pwd = util.MakeMD5(input.Pwd)
	input.CreateAt = schema.LocalTime(time.Now())

	// Write input data to DB
	_, err = a.mdlUser.Add(input, nil)
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

func (a *AcctBook) UpdateUser(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.User)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	//
	input.Pwd = util.MakeMD5(input.Pwd)

	// Update Data
	err := a.mdlUser.Update(input, nil)
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

func (a *AcctBook) DeleteUser(c *gin.Context) {
	// Bind input data; Map the JSON data to struct
	input := new(schema.User)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.HttpFailResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	//
	input.Pwd = util.MakeMD5(input.Pwd)

	// Delete row from DB
	err := a.mdlUser.Delete(input.Id, input.Uid, input.Pwd, nil)
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
