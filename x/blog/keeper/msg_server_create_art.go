package keeper

import (
	"context"

	"github.com/cosmonaut/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateArt(goCtx context.Context, msg *types.MsgCreateArt) (*types.MsgCreateArtResponse, error) {
	// Get the context
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Create variable of type Post
	var art = types.Art{
		Creator:      msg.Creator,
		Title:        msg.Title,
		Body:         msg.Body,
		Price:        msg.Price,
		OldArt:       msg.OldArt,
		AuthorArt:    msg.AuthorArt,
		AuthorSender: msg.Creator,
		Published:    msg.Published,
		Image:        msg.Image,
	}
	// Add a post to the store and get back the ID
	id := k.AppendArt(ctx, art)
	// Return the ID of the post
	return &types.MsgCreateArtResponse{Id: id}, nil
}
