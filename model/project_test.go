package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createProject(t *testing.T) schema.Project {
	// Create user data
	user := createUser(t)

	data := schema.Project{
		Id:     "",
		UserId: user.Id,
		Name:   util.RandomProjectName(),
		Icon:   util.RandomProjectIcon(),
	}

	result, err := testProjectModel.Add(&data, nil)

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

func TestCreateProject(t *testing.T) {
	createProject(t)
}

func TestGetProject(t *testing.T) {
	// Create test data
	dataCreated := createProject(t)

	// Get data
	dataQried, err := testProjectModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataCreated.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestUpdateProject(t *testing.T) {
	// Create test data
	dataCreated := createProject(t)

	// Update data
	dataUpdate := schema.Project{
		Id:   dataCreated.Id,
		Name: util.RandomProjectName(),
	}

	err := testProjectModel.Update(&dataUpdate, nil)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testProjectModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried[0])

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.Equal(t, dataUpdate.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestDeleteProject(t *testing.T) {
	// Create test data
	dataCreated := createProject(t)

	// Delete data
	err := testProjectModel.Delete(dataCreated.Id, nil)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testProjectModel.Get(dataCreated.UserId, nil)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
