package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) schema.User {
	data := schema.User{
		Id:       "",
		Name:     util.RandomName(),
		Uid:      util.RandomUID(),
		Pwd:      util.MakeMD5(util.RandomPwd()),
		Currency: util.RandomCurrency(),
		CreateAt: schema.LocalTime(time.Now()),
	}

	result, err := testUserModel.Add(&data)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Validate Data
	require.NotEmpty(t, result.Id)
	require.Equal(t, data.Name, result.Name)
	require.Equal(t, data.Uid, result.Uid)
	require.Equal(t, data.Pwd, result.Pwd)
	require.Equal(t, data.Currency, result.Currency)
	require.NotZero(t, result.CreateAt)

	return data
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestGetUser(t *testing.T) {
	// Create test data
	dataCreated := createUser(t)

	// Get data
	dataQried, err := testUserModel.Get(dataCreated.Uid, dataCreated.Pwd)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Uid, dataQried[0].Uid)
	require.Equal(t, dataCreated.Pwd, dataQried[0].Pwd)
	require.Equal(t, dataCreated.Currency, dataQried[0].Currency)
	require.WithinDuration(t, time.Time(dataCreated.CreateAt), time.Time(dataQried[0].CreateAt), time.Second)
}

func TestUpdateUser(t *testing.T) {
	// Create test data
	dataCreated := createUser(t)

	// Update data
	dataUpdate := schema.User{
		Id:   dataCreated.Id,
		Name: util.RandomName(),
	}

	err := testUserModel.Update(&dataUpdate)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testUserModel.GetByID(dataUpdate.Id)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried[0])

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataUpdate.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Uid, dataQried[0].Uid)
	require.Equal(t, dataCreated.Pwd, dataQried[0].Pwd)
	require.Equal(t, dataCreated.Currency, dataQried[0].Currency)
	require.WithinDuration(t, time.Time(dataCreated.CreateAt), time.Time(dataQried[0].CreateAt), time.Second)
}

func TestDeleteUser(t *testing.T) {
	// Create test data
	dataCreated := createUser(t)

	// Delete data
	err := testUserModel.Delete(dataCreated.Id, dataCreated.Uid, dataCreated.Pwd)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testUserModel.GetByID(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
