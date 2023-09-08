package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "koin/testutil/keeper"
	"koin/x/koin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KoinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
