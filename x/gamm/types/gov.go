package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeChangePoolParamIncentives = "ChangePoolParamProposal"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeChangePoolParamIncentives)
	govtypes.RegisterProposalTypeCodec(&ChangePoolParamProposal{}, "osmosis/ChangePoolParamProposal")
}

var _ govtypes.Content = &ChangePoolParamProposal{}

func NewChangePoolParamProposal(title, description string, poolID uint64, exitFee sdk.Dec, swapFee sdk.Dec) govtypes.Content {
	return &ChangePoolParamProposal{
		Title:       title,
		Description: description,
		PoolId:      poolID,
		ExitFee:     exitFee,
		SwapFee:     swapFee,
	}
}

func (p *ChangePoolParamProposal) GetTitle() string { return p.Title }

func (p *ChangePoolParamProposal) GetDescription() string { return p.Description }

func (p *ChangePoolParamProposal) ProposalRoute() string { return RouterKey }

func (p *ChangePoolParamProposal) ProposalType() string {
	return ProposalTypeChangePoolParamIncentives
}

func (p *ChangePoolParamProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *ChangePoolParamProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Update Pool Params Proposal:
  Title:       %s
  Description: %s
  PoolId:		%d
  ExitFee:     %s
  SwapFee:     %s
`, p.Title, p.Description, p.PoolId, p.ExitFee, p.SwapFee))
	return b.String()
}
