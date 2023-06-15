package keeper

import (
	"context"
	"strconv"

	"oracle/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FeedPrice(goCtx context.Context, msg *types.MsgFeedPrice) (*types.MsgFeedPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	usd, err := strconv.ParseUint(msg.Price, 10, 32)
	if err != nil {
		panic("Error: usd value")
	}

	price, found := k.Keeper.GetPrice(ctx, msg.Denom)

	if !found {
		price = types.Price{
			Denom: msg.Denom,
			Usd:   usd,
		}
	} else {
		price.Usd = usd
	}

	k.Keeper.SetPrice(ctx, price)

	return &types.MsgFeedPriceResponse{}, nil
}
