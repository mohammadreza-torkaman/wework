package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitJob = "submit_job"

var _ sdk.Msg = &MsgSubmitJob{}

func NewMsgSubmitJob(creator string, title string, description string, tags string, postDeadline string, jobDeadline string, maxPrice string, location string, jobtype uint64) *MsgSubmitJob {
	return &MsgSubmitJob{
		Creator:      creator,
		Title:        title,
		Description:  description,
		Tags:         tags,
		PostDeadline: postDeadline,
		JobDeadline:  jobDeadline,
		MaxPrice:     maxPrice,
		Location:     location,
		Jobtype:      jobtype,
	}
}

func (msg *MsgSubmitJob) Route() string {
	return RouterKey
}

func (msg *MsgSubmitJob) Type() string {
	return TypeMsgSubmitJob
}

func (msg *MsgSubmitJob) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitJob) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitJob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
