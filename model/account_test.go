package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAccount(t *testing.T) schema.Account {
	// Create user data
	user := createUser(t)

	data := schema.Account{
		Id:     "",
		UserId: user.Id,
		Name:   util.RandomAccountName(),
		Icon:   util.RandomAccountIcon(),
	}

	result, err := testAccountModel.Add(&data, nil)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Validate Data
	require.NotEmpty(t, result.Id)
	require.Equal(t, data.UserId, result.UserId)
	require.Equal(t, data.Name, result.Name)
	require.Equal(t, data.Icon, result.Icon)

	return data
}

func TestCreateAccount(t *testing.T) {
	createAccount(t)
}

func TestGetAccount(t *testing.T) {
	// Create test data
	dataCreated := createAccount(t)

	// Get data
	dataQried, err := testAccountModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataCreated.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestUpdateAccount(t *testing.T) {
	// Create test data
	dataCreated := createAccount(t)

	// Update data
	dataUpdate := schema.Account{
		Id:   dataCreated.Id,
		Name: util.RandomAccountName(),
	}

	err := testAccountModel.Update(&dataUpdate, nil)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testAccountModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried[0])

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataUpdate.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestDeleteAccount(t *testing.T) {
	// Create test data
	dataCreated := createAccount(t)

	// Delete data
	err := testAccountModel.Delete(dataCreated.Id, nil)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testAccountModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
