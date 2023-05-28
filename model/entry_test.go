package model

import (
	"account-book/lib/pgdb/schema"
	"account-book/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createEntry(t *testing.T) schema.Entry {
	// Create user data
	user := createUser(t)
	entryType := createType(t)
	acct := createAccount(t)
	prj := createProject(t)
	store := createStore(t)

	data := schema.Entry{
		Id:       "",
		UserId:   user.Id,
		Time:     schema.LocalTime(time.Now()),
		Behavior: int(util.RandomInt(0, 2)),
		Amount:   int(util.RandomMoney()),
		Type:     entryType.Id,
		Account:  acct.Id,
		Project:  prj.Id,
		Store:    store.Id,
		Note:     util.RandomString(8),
	}

	result, err := testEntryModel.Add(&data)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Validate Data
	require.NotEmpty(t, result.Id)
	require.Equal(t, data.UserId, result.UserId)
	require.NotZero(t, result.Time)
	require.Equal(t, data.Behavior, result.Behavior)
	require.Equal(t, data.Amount, result.Amount)
	require.Equal(t, data.Type, result.Type)
	require.Equal(t, data.Account, result.Account)
	require.Equal(t, data.Project, result.Project)
	require.Equal(t, data.Store, result.Store)
	require.Equal(t, data.Note, result.Note)

	return data
}

func TestCreateEntry(t *testing.T) {
	createEntry(t)
}

func TestGetEntry(t *testing.T) {
	// Create test data
	dataCreated := createEntry(t)

	// Get data
	dataQried, err := testEntryModel.Get(time.Time(dataCreated.Time), dataCreated.UserId)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried)

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.WithinDuration(t, time.Time(dataCreated.Time), time.Time(dataQried[0].Time), time.Second)
	require.Equal(t, dataCreated.Behavior, dataQried[0].Behavior)
	require.Equal(t, dataCreated.Amount, dataQried[0].Amount)
	require.Equal(t, dataCreated.Type, dataQried[0].Type)
	require.Equal(t, dataCreated.Account, dataQried[0].Account)
	require.Equal(t, dataCreated.Project, dataQried[0].Project)
	require.Equal(t, dataCreated.Store, dataQried[0].Store)
	require.Equal(t, dataCreated.Note, dataQried[0].Note)
}

func TestUpdateEntry(t *testing.T) {
	// Create test data
	dataCreated := createEntry(t)

	// Update data
	dataUpdate := schema.Entry{
		Id:     dataCreated.Id,
		Amount: int(util.RandomMoney()),
	}

	err := testEntryModel.Update(&dataUpdate)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testEntryModel.GetByID(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)
	require.NotEmpty(t, dataQried[0])

	// Validate Data
	require.Equal(t, dataCreated.Id, dataQried[0].Id)
	require.Equal(t, dataCreated.UserId, dataQried[0].UserId)
	require.WithinDuration(t, time.Time(dataCreated.Time), time.Time(dataQried[0].Time), time.Second)
	require.Equal(t, dataCreated.Behavior, dataQried[0].Behavior)
	require.Equal(t, dataUpdate.Amount, dataQried[0].Amount)
	require.Equal(t, dataCreated.Type, dataQried[0].Type)
	require.Equal(t, dataCreated.Account, dataQried[0].Account)
	require.Equal(t, dataCreated.Project, dataQried[0].Project)
	require.Equal(t, dataCreated.Store, dataQried[0].Store)
	require.Equal(t, dataCreated.Note, dataQried[0].Note)
}

func TestDeleteEntry(t *testing.T) {
	// Create test data
	dataCreated := createEntry(t)

	// Delete data
	err := testEntryModel.Delete(dataCreated.UserId, dataCreated.Id)

	// Validate Command
	require.NoError(t, err)

	// Get Data
	dataQried, err := testEntryModel.GetByID(dataCreated.Id)

	// Validate Command
	require.NoError(t, err)
	require.Empty(t, dataQried)
}
