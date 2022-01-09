package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

func (k Keeper) HandleChangeParamProposal(ctx sdk.Context, p *types.ChangePoolParamProposal) error {
	return k.EditPoolFeeParams(ctx, p.PoolId, p.ExitFee, p.SwapFee)
}
