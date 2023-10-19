package secretai_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "secretai/testutil/keeper"
	"secretai/testutil/nullify"
	"secretai/x/secretai"
	"secretai/x/secretai/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SecretaiKeeper(t)
	secretai.InitGenesis(ctx, *k, genesisState)
	got := secretai.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
