package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "oracle/testutil/keeper"
	"oracle/testutil/nullify"
	"oracle/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPriceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPrice(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPriceRequest
		response *types.QueryGetPriceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPriceRequest{
				Denom: msgs[0].Denom,
			},
			response: &types.QueryGetPriceResponse{Price: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPriceRequest{
				Denom: msgs[1].Denom,
			},
			response: &types.QueryGetPriceResponse{Price: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPriceRequest{
				Denom: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Price(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPriceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPrice(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPriceRequest {
		return &types.QueryAllPriceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PriceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Price), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Price),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PriceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Price), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Price),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PriceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Price),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PriceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
