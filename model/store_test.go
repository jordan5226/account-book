package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createStore(t *testing.T) schema.Store {
	// Create user data
	user := createUser(t)

	data := schema.Store{
		Id:     "",
		UserId: user.Id,
		Name:   util.RandomStoreName(),
		Icon:   util.RandomStoreIcon(),
	}

	result, err := testStoreModel.Add(&data)

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

func TestCreateStore(t *testing.T) {
	createStore(t)
}

func TestGetStore(t *testing.T) {
	// Create test data
	dataCreated := createStore(t)

	// Get data
	dataQried, err := testStoreModel.Get(dataCreated.UserId)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataCreated.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestUpdateStore(t *testing.T) {
	// Create test data
	dataCreated := createStore(t)

	// Update data
	dataUpdate := schema.Store{
		Id:   dataCreated.Id,
		Name: util.RandomStoreName(),
	}

	err := testStoreModel.Update(&dataUpdate)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testStoreModel.Get(dataCreated.UserId)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried[0])

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataUpdate.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestDeleteStore(t *testing.T) {
	// Create test data
	dataCreated := createStore(t)

	// Delete data
	err := testStoreModel.Delete(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testStoreModel.Get(dataCreated.UserId)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
