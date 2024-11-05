package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&ReqCheckersTorram{}, // Add our new message type
	)

	// Register the CheckersTorram service instead of Msg
	msgservice.RegisterMsgServiceDesc(registry, &_CheckersTorram_serviceDesc)
}
