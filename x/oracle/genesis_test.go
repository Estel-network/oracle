package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "oracle/testutil/keeper"
	"oracle/testutil/nullify"
	"oracle/x/oracle"
	"oracle/x/oracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PriceList: []types.Price{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PriceList, got.PriceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
