package model

import (
	"account-book/lib/pgdb"
	"account-book/util"
	"log"
	"os"
	"testing"
)

var testAccountModel *AccountModel
var testUserModel *UserModel
var testProjectModel *ProjectModel
var testStoreModel *StoreModel
var testTypeModel *TypeModel
var testEntryModel *EntryModel

func TestMain(m *testing.M) {
	util.LoadEnv()
	Init(pgdb.GetConnect())

	if GetDB() == nil {
		log.Fatal("TestMain - Cannot connect to DB!")
	}

	//
	testUserModel = NewUserModel()
	testAccountModel = NewAccountModel()
	testProjectModel = NewProjectModel()
	testStoreModel = NewStoreModel()
	testTypeModel = NewTypeModel()
	testEntryModel = NewEntryModel()

	os.Exit(m.Run())
}
