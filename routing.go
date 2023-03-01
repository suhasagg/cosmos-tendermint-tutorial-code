package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/mymodule"
)

func NewApp() *types.App {
	// Define the bank module
	bankModule := bank.NewModule()

	// Define the auth module
	authModule := auth.NewModule()

	// Define the mymodule module
	myModule := mymodule.NewModule()

	// Define the module manager
	moduleManager := module.NewManager(
		bankModule,
		authModule,
		myModule,
	)

	// Define the application
	app := types.NewApp()

	// Register the module routes with the application router
	moduleManager.RegisterRoutes(app.Router())

	return app
}
