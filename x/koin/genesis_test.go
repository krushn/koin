package koin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "koin/testutil/keeper"
	"koin/testutil/nullify"
	"koin/x/koin"
	"koin/x/koin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KoinKeeper(t)
	koin.InitGenesis(ctx, *k, genesisState)
	got := koin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
