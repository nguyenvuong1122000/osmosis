package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	"github.com/osmosis-labs/osmosis/x/gamm/client/cli"
)

var ChangePoolParamHandler = govclient.NewProposalHandler(cli.NewChangeFeeProposal, emptyRestHandler)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-ibc-client",
		Handler:  nil,
	}
}
