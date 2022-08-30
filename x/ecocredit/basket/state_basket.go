package basket

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Validate performs basic validation of the Basket state type
func (m *Basket) Validate() error {
	if m.Id == 0 {
		return ecocredit.ErrParseFailure.Wrapf("id cannot be zero")
	}

	if err := ValidateBasketDenom(m.BasketDenom); err != nil {
		return errors.Wrap(err, "basket denom") // returns parse error
	}

	if err := ValidateBasketName(m.Name); err != nil {
		return errors.Wrap(err, "name") // returns parse error
	}

	if err := core.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return errors.Wrap(err, "credit type abbrev") // returns parse error
	}

	if err := m.DateCriteria.Validate(); err != nil {
		return err // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Curator).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("curator: %s", err)
	}

	return nil
}
