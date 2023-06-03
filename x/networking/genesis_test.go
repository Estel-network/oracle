package networking_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "oracle/testutil/keeper"
	"oracle/testutil/nullify"
	"oracle/x/networking"
	"oracle/x/networking/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NetworkingKeeper(t)
	networking.InitGenesis(ctx, *k, genesisState)
	got := networking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
