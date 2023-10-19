package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "secretai/testutil/keeper"
	"secretai/x/secretai/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SecretaiKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
