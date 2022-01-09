package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestChangeFeeProposalMarshalUnmarshal(t *testing.T) {

	tests := []struct {
		proposal *types.ChangePoolParamProposal
	}{
		{ // empty title
			proposal: &types.ChangePoolParamProposal{
				Title:       "proposal",
				Description: "proposal to update fee incentives",
				PoolId:      uint64(0),
				ExitFee:     sdk.Dec{},
				SwapFee:     sdk.Dec{},
			},
		},
		{
			proposal: &types.ChangePoolParamProposal{
				Title:       "",
				Description: "",
				PoolId:      uint64(0),
				ExitFee:     sdk.Dec{},
				SwapFee:     sdk.Dec{},
			},
		},
	}

	for _, test := range tests {
		bz, err := proto.Marshal(test.proposal)
		require.NoError(t, err)
		decoded := types.ChangePoolParamProposal{}
		err = proto.Unmarshal(bz, &decoded)
		require.NoError(t, err)
		require.Equal(t, *test.proposal, decoded)
	}
}
