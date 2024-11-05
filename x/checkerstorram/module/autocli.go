package checkerstorram

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "checkers-torram/api/checkerstorram/checkerstorram"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "GetStoredGame",
					Use:       "get-game [index]",
					Short:     "Query a game by index",
					Long:      "Get the stored game information for a specific game index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{
							ProtoField: "index",
						},
					},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.CheckersTorram_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod: "CheckersCreateGm",
					Use:       "create-game [black] [red]",
					Short:     "Create a new game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "black"},
						{ProtoField: "red"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
