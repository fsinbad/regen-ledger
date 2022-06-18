package core

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgMintBatchCredits{}

// Route implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgMintBatchCredits) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgMintBatchCredits) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		return sdkerrors.Wrap(err, "malformed issuer address")
	}
	if err := ValidateBatchDenom(m.BatchDenom); err != nil {
		return err
	}
	if len(m.Note) > MaxNoteLength {
		return errBadReq.Wrapf("note must not be longer than %d characters", MaxNoteLength)
	}
	if err = validateBatchIssuances(m.Issuance); err != nil {
		return err
	}
	return validateOriginTx(m.OriginTx, true)
}

// GetSigners returns the expected signers for MsgMintBatchCredits.
func (m *MsgMintBatchCredits) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

func validateBatchIssuances(iss []*BatchIssuance) error {
	if len(iss) == 0 {
		errBadReq.Wrap("issuance list must not be empty")
	}
	for idx, i := range iss {
		if i == nil {
			return errBadReq.Wrapf("issuance[%d] must be defined", idx)
		}
		_, err := sdk.AccAddressFromBech32(i.Recipient)
		if err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("issuance[%d].recipient", idx)
		}
		if i.TradableAmount != "" {
			if _, err := math.NewNonNegativeDecFromString(i.TradableAmount); err != nil {
				return errBadReq.Wrapf("issuance[%d].tradable_amount; %v", idx, err)
			}
		}

		if i.RetiredAmount != "" {
			retiredAmount, err := math.NewNonNegativeDecFromString(i.RetiredAmount)
			if err != nil {
				return errBadReq.Wrapf("issuance[%d].retired_amount; %v", idx, err)
			}

			if !retiredAmount.IsZero() {
				if err = ValidateJurisdiction(i.RetirementJurisdiction); err != nil {
					return errBadReq.Wrapf("issuance[%d].retirement_jurisdiction; %v", idx, err)
				}
			}
		}
	}
	return nil
}

var reOriginTxId = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 _\-]{0,127}$`)
var reOriginTxSource = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 _\-]{0,31}$`)

func validateOriginTx(o *OriginTx, required bool) error {
	if o == nil {
		if required {
			return errBadReq.Wrap("origin_tx is required")
		}
		return nil
	}
	if !reOriginTxId.MatchString(o.Id) {
		return errBadReq.Wrap("origin_tx.id must be at most 128 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}
	if !reOriginTxSource.MatchString(o.Source) {
		return errBadReq.Wrap("origin_tx.source must be at most 32 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}
	return nil
}
