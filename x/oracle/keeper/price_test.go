package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "oracle/testutil/keeper"
	"oracle/testutil/nullify"
	"oracle/x/oracle/keeper"
	"oracle/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPrice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Price {
	items := make([]types.Price, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetPrice(ctx, items[i])
	}
	return items
}

func TestPriceGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPrice(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrice(ctx,
			item.Denom,
		)
		_, found := keeper.GetPrice(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestPriceGetAll(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPrice(ctx)),
	)
}
