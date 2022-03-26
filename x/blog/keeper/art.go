package keeper

import (
	"encoding/binary"
	"github.com/cosmonaut/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendArt(ctx sdk.Context, art types.Art) uint64 {
	// Get the current number of posts in the store
	count := k.GetArtCount(ctx)
	// Assign an ID to the post based on the number of posts in the store
	art.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ArtKey))
	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, art.Id)
	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&art)
	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)
	// Update the post count
	k.SetArtCount(ctx, count+1)
	return count
}

func (k Keeper) GetArtCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ArtCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.ArtCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetArtCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ArtCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.ArtCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Post-count- to count
	store.Set(byteKey, bz)
}
