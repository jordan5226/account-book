package acctbook

import (
	"account-book/model"

	"github.com/gin-gonic/gin"
)

type IAcctBook interface {
	SetRout(grp *gin.RouterGroup)
	//
	GetEntries(c *gin.Context)
	CreateEntry(c *gin.Context)
	UpdateEntry(c *gin.Context)
	DeleteEntry(c *gin.Context)
	//
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	//
	GetProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
	//
	GetStores(c *gin.Context)
	CreateStore(c *gin.Context)
	UpdateStore(c *gin.Context)
	DeleteStore(c *gin.Context)
	//
	GetTypes(c *gin.Context)
	CreateType(c *gin.Context)
	UpdateType(c *gin.Context)
	DeleteType(c *gin.Context)
	//
	GetAccounts(c *gin.Context)
	CreateAccount(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type AcctBook struct {
	mdlEntry *model.EntryModel
	mdlUser  *model.UserModel
	mdlPrj   *model.ProjectModel
	mdlStore *model.StoreModel
	mdlType  *model.TypeModel
	mdlAcct  *model.AccountModel
}

func New() IAcctBook {
	ab := new(AcctBook)
	ab.mdlEntry = model.NewEntryModel()
	ab.mdlUser = model.NewUserModel()
	ab.mdlPrj = model.NewProjectModel()
	ab.mdlStore = model.NewStoreModel()
	ab.mdlType = model.NewTypeModel()
	ab.mdlAcct = model.NewAccountModel()
	return ab
}

func (a *AcctBook) SetRout(grp *gin.RouterGroup) {
	grp.GET("/entry/:time/*uid", a.GetEntries)
	grp.POST("/entry", a.CreateEntry)
	grp.PUT("/entry", a.UpdateEntry)
	grp.DELETE("/entry/:uid/*id", a.DeleteEntry)

	grp.GET("/user", a.GetUser)
	grp.POST("/user", a.CreateUser)
	grp.PUT("/user", a.UpdateUser)
	grp.DELETE("/user", a.DeleteUser)

	grp.GET("/prj/:uid", a.GetProjects)
	grp.POST("/prj", a.CreateProject)
	grp.PUT("/prj", a.UpdateProject)
	grp.DELETE("/prj/:id", a.DeleteProject)

	grp.GET("/store/:uid", a.GetStores)
	grp.POST("/store", a.CreateStore)
	grp.PUT("/store", a.UpdateStore)
	grp.DELETE("/store/:id", a.DeleteStore)

	grp.GET("/type", a.GetTypes)
	grp.POST("/type", a.CreateType)
	grp.PUT("/type", a.UpdateType)
	grp.DELETE("/type/:id", a.DeleteType)

	grp.GET("/acct/:uid", a.GetAccounts)
	grp.POST("/acct", a.CreateAccount)
	grp.PUT("/acct", a.UpdateAccount)
	grp.DELETE("/acct/:id", a.DeleteAccount)
}
