package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateArt = "create_art"

var _ sdk.Msg = &MsgCreateArt{}

func NewMsgCreateArt(creator string, title string, body string, price string, oldArt string, authorArt string, authorSender string, published string, image string) *MsgCreateArt {
	return &MsgCreateArt{
		Creator:      creator,
		Title:        title,
		Body:         body,
		Price:        price,
		OldArt:       oldArt,
		AuthorArt:    authorArt,
		AuthorSender: authorSender,
		Published:    published,
		Image:        image,
	}
}

func (msg *MsgCreateArt) Route() string {
	return RouterKey
}

func (msg *MsgCreateArt) Type() string {
	return TypeMsgCreateArt
}

func (msg *MsgCreateArt) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateArt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateArt) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
