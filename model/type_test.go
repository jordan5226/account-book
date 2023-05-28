package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createType(t *testing.T) schema.Type {
	data := schema.Type{
		Id:   "",
		Name: util.RandomTypeName(),
		Icon: util.RandomTypeIcon(),
	}

	result, err := testTypeModel.Add(&data)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Validate Data
	require.NotEmpty(t, result.Id)
	require.Equal(t, data.Name, result.Name)
	require.Equal(t, data.Icon, result.Icon)

	return data
}

func TestCreateType(t *testing.T) {
	createType(t)
}

func TestGetType(t *testing.T) {
	// Create test data
	dataCreated := createType(t)

	// Get data
	dataQried, err := testTypeModel.Get()

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	lastIdx := len(dataQried) - 1
	require.Equal(t, dataCreated.Id, dataQried[lastIdx].Id)
	require.Equal(t, dataCreated.Name, dataQried[lastIdx].Name)
	require.Equal(t, dataCreated.Icon, dataQried[lastIdx].Icon)
}

func TestUpdateType(t *testing.T) {
	// Create test data
	dataCreated := createType(t)

	// Update data
	dataUpdate := schema.Type{
		Id:   dataCreated.Id,
		Name: "UpdatedTypeName", //util.RandomTypeName(),
	}

	err := testTypeModel.Update(&dataUpdate)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testTypeModel.GetByID(dataUpdate.Id)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataUpdate.Name, dataQried[0].Name)
	require.Equal(t, dataCreated.Icon, dataQried[0].Icon)
}

func TestDeleteType(t *testing.T) {
	// Create test data
	dataCreated := createType(t)

	// Delete data
	err := testTypeModel.Delete(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testTypeModel.GetByID(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
