package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/hello/testutil/keeper"
	"github.com/cosmonaut/hello/x/hello/keeper"
	"github.com/cosmonaut/hello/x/hello/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HelloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
