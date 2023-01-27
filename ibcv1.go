package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	"github.com/cosmos/cosmos-sdk/x/ibc/05-port"
)

func main() {
	// Create a new Cosmos SDK application
	config := sdk.NewConfig()
	config.SetBech32PrefixForAccount("cosmos", "cosmospub")
	config.SetBech32PrefixForValidator("cosmosvaloper")
	config.SetBech32PrefixForConsensusNode("cosmosvalcons")
	config.Seal()
	
	// Add the IBC module to the application
	app := server.NewServer(config)
	app.Initialize()
	app.IBC = ibc.NewIBCModule(app)
	app.IBC.RegisterInterfaces(app.Interfaces)
	
	// Register the IBC connection, channel, and port submodules
	connection.RegisterInterfaces(app.Interfaces)
	channel.RegisterInterfaces(app.Interfaces)
	port.RegisterInterfaces(app.Interfaces)
	
	// Start the application
	app.Start()
}
