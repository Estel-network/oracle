package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"oracle/x/oracle/types"
)

func (k Keeper) PriceAll(goCtx context.Context, req *types.QueryAllPriceRequest) (*types.QueryAllPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var prices []types.Price
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.KeyPrefix(types.PriceKeyPrefix))

	pageRes, err := query.Paginate(priceStore, req.Pagination, func(key []byte, value []byte) error {
		var price types.Price
		if err := k.cdc.Unmarshal(value, &price); err != nil {
			return err
		}

		prices = append(prices, price)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPriceResponse{Price: prices, Pagination: pageRes}, nil
}

func (k Keeper) Price(goCtx context.Context, req *types.QueryGetPriceRequest) (*types.QueryGetPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPrice(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPriceResponse{Price: val}, nil
}
