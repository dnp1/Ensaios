package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

func createAdminModule (baseModuleID int64) {
	var err error
	var adminModule = &model.ModuleData{
		Name:"admin",
		ParentModuleID: sql.NullInt64{
			Int64: baseModuleID,
			Valid: true,
		},
	}
	err = db.Module.Create(adminModule)
	throwPanic(err)
}

func createBaseModule () *model.ModuleData {
	var err error
	var baseModule = &model.ModuleData{
		Name:"",
	}
	err = db.Module.Create(baseModule)
	throwPanic(err)

	var loginOp = &model.OperationData{
		Name:"_login",
		ModuleID: baseModule.ID,
	}

	err = db.Operation.Create(loginOp)
	throwPanic(err)

	var baseOp = &model.OperationData{
		Name:"",
		ModuleID: baseModule.ID,
	}

	err = db.Operation.Create(baseOp)
	throwPanic(err)

	return baseModule;
}


func Bootstrap() {
	baseModule := createBaseModule()
	createAdminModule(baseModule.ID)
}