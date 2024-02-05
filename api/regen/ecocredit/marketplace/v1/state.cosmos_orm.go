// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package marketplacev1

import (
	context "context"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type SellOrderTable interface {
	Insert(ctx context.Context, sellOrder *SellOrder) error
	InsertReturningID(ctx context.Context, sellOrder *SellOrder) (uint64, error)
	Update(ctx context.Context, sellOrder *SellOrder) error
	Save(ctx context.Context, sellOrder *SellOrder) error
	Delete(ctx context.Context, sellOrder *SellOrder) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*SellOrder, error)
	List(ctx context.Context, prefixKey SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error)
	ListRange(ctx context.Context, from, to SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error)
	DeleteBy(ctx context.Context, prefixKey SellOrderIndexKey) error
	DeleteRange(ctx context.Context, from, to SellOrderIndexKey) error

	doNotImplement()
}

type SellOrderIterator struct {
	ormtable.Iterator
}

func (i SellOrderIterator) Value() (*SellOrder, error) {
	var sellOrder SellOrder
	err := i.UnmarshalMessage(&sellOrder)
	return &sellOrder, err
}

type SellOrderIndexKey interface {
	id() uint32
	values() []interface{}
	sellOrderIndexKey()
}

// primary key starting index..
type SellOrderPrimaryKey = SellOrderIdIndexKey

type SellOrderIdIndexKey struct {
	vs []interface{}
}

func (x SellOrderIdIndexKey) id() uint32            { return 0 }
func (x SellOrderIdIndexKey) values() []interface{} { return x.vs }
func (x SellOrderIdIndexKey) sellOrderIndexKey()    {}

func (this SellOrderIdIndexKey) WithId(id uint64) SellOrderIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type SellOrderBatchKeyIndexKey struct {
	vs []interface{}
}

func (x SellOrderBatchKeyIndexKey) id() uint32            { return 1 }
func (x SellOrderBatchKeyIndexKey) values() []interface{} { return x.vs }
func (x SellOrderBatchKeyIndexKey) sellOrderIndexKey()    {}

func (this SellOrderBatchKeyIndexKey) WithBatchKey(batch_key uint64) SellOrderBatchKeyIndexKey {
	this.vs = []interface{}{batch_key}
	return this
}

type SellOrderSellerIndexKey struct {
	vs []interface{}
}

func (x SellOrderSellerIndexKey) id() uint32            { return 2 }
func (x SellOrderSellerIndexKey) values() []interface{} { return x.vs }
func (x SellOrderSellerIndexKey) sellOrderIndexKey()    {}

func (this SellOrderSellerIndexKey) WithSeller(seller []byte) SellOrderSellerIndexKey {
	this.vs = []interface{}{seller}
	return this
}

type SellOrderExpirationIndexKey struct {
	vs []interface{}
}

func (x SellOrderExpirationIndexKey) id() uint32            { return 3 }
func (x SellOrderExpirationIndexKey) values() []interface{} { return x.vs }
func (x SellOrderExpirationIndexKey) sellOrderIndexKey()    {}

func (this SellOrderExpirationIndexKey) WithExpiration(expiration *timestamppb.Timestamp) SellOrderExpirationIndexKey {
	this.vs = []interface{}{expiration}
	return this
}

type sellOrderTable struct {
	table ormtable.AutoIncrementTable
}

func (this sellOrderTable) Insert(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Insert(ctx, sellOrder)
}

func (this sellOrderTable) Update(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Update(ctx, sellOrder)
}

func (this sellOrderTable) Save(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Save(ctx, sellOrder)
}

func (this sellOrderTable) Delete(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Delete(ctx, sellOrder)
}

func (this sellOrderTable) InsertReturningID(ctx context.Context, sellOrder *SellOrder) (uint64, error) {
	return this.table.InsertReturningID(ctx, sellOrder)
}

func (this sellOrderTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this sellOrderTable) Get(ctx context.Context, id uint64) (*SellOrder, error) {
	var sellOrder SellOrder
	found, err := this.table.PrimaryKey().Get(ctx, &sellOrder, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &sellOrder, nil
}

func (this sellOrderTable) List(ctx context.Context, prefixKey SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return SellOrderIterator{it}, err
}

func (this sellOrderTable) ListRange(ctx context.Context, from, to SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return SellOrderIterator{it}, err
}

func (this sellOrderTable) DeleteBy(ctx context.Context, prefixKey SellOrderIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this sellOrderTable) DeleteRange(ctx context.Context, from, to SellOrderIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this sellOrderTable) doNotImplement() {}

var _ SellOrderTable = sellOrderTable{}

func NewSellOrderTable(db ormtable.Schema) (SellOrderTable, error) {
	table := db.GetTable(&SellOrder{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&SellOrder{}).ProtoReflect().Descriptor().FullName()))
	}
	return sellOrderTable{table.(ormtable.AutoIncrementTable)}, nil
}

type AllowedDenomTable interface {
	Insert(ctx context.Context, allowedDenom *AllowedDenom) error
	Update(ctx context.Context, allowedDenom *AllowedDenom) error
	Save(ctx context.Context, allowedDenom *AllowedDenom) error
	Delete(ctx context.Context, allowedDenom *AllowedDenom) error
	Has(ctx context.Context, bank_denom string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, bank_denom string) (*AllowedDenom, error)
	HasByDisplayDenom(ctx context.Context, display_denom string) (found bool, err error)
	// GetByDisplayDenom returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByDisplayDenom(ctx context.Context, display_denom string) (*AllowedDenom, error)
	List(ctx context.Context, prefixKey AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error)
	ListRange(ctx context.Context, from, to AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error)
	DeleteBy(ctx context.Context, prefixKey AllowedDenomIndexKey) error
	DeleteRange(ctx context.Context, from, to AllowedDenomIndexKey) error

	doNotImplement()
}

type AllowedDenomIterator struct {
	ormtable.Iterator
}

func (i AllowedDenomIterator) Value() (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	err := i.UnmarshalMessage(&allowedDenom)
	return &allowedDenom, err
}

type AllowedDenomIndexKey interface {
	id() uint32
	values() []interface{}
	allowedDenomIndexKey()
}

// primary key starting index..
type AllowedDenomPrimaryKey = AllowedDenomBankDenomIndexKey

type AllowedDenomBankDenomIndexKey struct {
	vs []interface{}
}

func (x AllowedDenomBankDenomIndexKey) id() uint32            { return 0 }
func (x AllowedDenomBankDenomIndexKey) values() []interface{} { return x.vs }
func (x AllowedDenomBankDenomIndexKey) allowedDenomIndexKey() {}

func (this AllowedDenomBankDenomIndexKey) WithBankDenom(bank_denom string) AllowedDenomBankDenomIndexKey {
	this.vs = []interface{}{bank_denom}
	return this
}

type AllowedDenomDisplayDenomIndexKey struct {
	vs []interface{}
}

func (x AllowedDenomDisplayDenomIndexKey) id() uint32            { return 1 }
func (x AllowedDenomDisplayDenomIndexKey) values() []interface{} { return x.vs }
func (x AllowedDenomDisplayDenomIndexKey) allowedDenomIndexKey() {}

func (this AllowedDenomDisplayDenomIndexKey) WithDisplayDenom(display_denom string) AllowedDenomDisplayDenomIndexKey {
	this.vs = []interface{}{display_denom}
	return this
}

type allowedDenomTable struct {
	table ormtable.Table
}

func (this allowedDenomTable) Insert(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Insert(ctx, allowedDenom)
}

func (this allowedDenomTable) Update(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Update(ctx, allowedDenom)
}

func (this allowedDenomTable) Save(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Save(ctx, allowedDenom)
}

func (this allowedDenomTable) Delete(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Delete(ctx, allowedDenom)
}

func (this allowedDenomTable) Has(ctx context.Context, bank_denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, bank_denom)
}

func (this allowedDenomTable) Get(ctx context.Context, bank_denom string) (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	found, err := this.table.PrimaryKey().Get(ctx, &allowedDenom, bank_denom)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &allowedDenom, nil
}

func (this allowedDenomTable) HasByDisplayDenom(ctx context.Context, display_denom string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		display_denom,
	)
}

func (this allowedDenomTable) GetByDisplayDenom(ctx context.Context, display_denom string) (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &allowedDenom,
		display_denom,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &allowedDenom, nil
}

func (this allowedDenomTable) List(ctx context.Context, prefixKey AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AllowedDenomIterator{it}, err
}

func (this allowedDenomTable) ListRange(ctx context.Context, from, to AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AllowedDenomIterator{it}, err
}

func (this allowedDenomTable) DeleteBy(ctx context.Context, prefixKey AllowedDenomIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this allowedDenomTable) DeleteRange(ctx context.Context, from, to AllowedDenomIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this allowedDenomTable) doNotImplement() {}

var _ AllowedDenomTable = allowedDenomTable{}

func NewAllowedDenomTable(db ormtable.Schema) (AllowedDenomTable, error) {
	table := db.GetTable(&AllowedDenom{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&AllowedDenom{}).ProtoReflect().Descriptor().FullName()))
	}
	return allowedDenomTable{table}, nil
}

type MarketTable interface {
	Insert(ctx context.Context, market *Market) error
	InsertReturningID(ctx context.Context, market *Market) (uint64, error)
	Update(ctx context.Context, market *Market) error
	Save(ctx context.Context, market *Market) error
	Delete(ctx context.Context, market *Market) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Market, error)
	HasByCreditTypeAbbrevBankDenom(ctx context.Context, credit_type_abbrev string, bank_denom string) (found bool, err error)
	// GetByCreditTypeAbbrevBankDenom returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByCreditTypeAbbrevBankDenom(ctx context.Context, credit_type_abbrev string, bank_denom string) (*Market, error)
	List(ctx context.Context, prefixKey MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error)
	ListRange(ctx context.Context, from, to MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error)
	DeleteBy(ctx context.Context, prefixKey MarketIndexKey) error
	DeleteRange(ctx context.Context, from, to MarketIndexKey) error

	doNotImplement()
}

type MarketIterator struct {
	ormtable.Iterator
}

func (i MarketIterator) Value() (*Market, error) {
	var market Market
	err := i.UnmarshalMessage(&market)
	return &market, err
}

type MarketIndexKey interface {
	id() uint32
	values() []interface{}
	marketIndexKey()
}

// primary key starting index..
type MarketPrimaryKey = MarketIdIndexKey

type MarketIdIndexKey struct {
	vs []interface{}
}

func (x MarketIdIndexKey) id() uint32            { return 0 }
func (x MarketIdIndexKey) values() []interface{} { return x.vs }
func (x MarketIdIndexKey) marketIndexKey()       {}

func (this MarketIdIndexKey) WithId(id uint64) MarketIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type MarketCreditTypeAbbrevBankDenomIndexKey struct {
	vs []interface{}
}

func (x MarketCreditTypeAbbrevBankDenomIndexKey) id() uint32            { return 1 }
func (x MarketCreditTypeAbbrevBankDenomIndexKey) values() []interface{} { return x.vs }
func (x MarketCreditTypeAbbrevBankDenomIndexKey) marketIndexKey()       {}

func (this MarketCreditTypeAbbrevBankDenomIndexKey) WithCreditTypeAbbrev(credit_type_abbrev string) MarketCreditTypeAbbrevBankDenomIndexKey {
	this.vs = []interface{}{credit_type_abbrev}
	return this
}

func (this MarketCreditTypeAbbrevBankDenomIndexKey) WithCreditTypeAbbrevBankDenom(credit_type_abbrev string, bank_denom string) MarketCreditTypeAbbrevBankDenomIndexKey {
	this.vs = []interface{}{credit_type_abbrev, bank_denom}
	return this
}

type marketTable struct {
	table ormtable.AutoIncrementTable
}

func (this marketTable) Insert(ctx context.Context, market *Market) error {
	return this.table.Insert(ctx, market)
}

func (this marketTable) Update(ctx context.Context, market *Market) error {
	return this.table.Update(ctx, market)
}

func (this marketTable) Save(ctx context.Context, market *Market) error {
	return this.table.Save(ctx, market)
}

func (this marketTable) Delete(ctx context.Context, market *Market) error {
	return this.table.Delete(ctx, market)
}

func (this marketTable) InsertReturningID(ctx context.Context, market *Market) (uint64, error) {
	return this.table.InsertReturningID(ctx, market)
}

func (this marketTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this marketTable) Get(ctx context.Context, id uint64) (*Market, error) {
	var market Market
	found, err := this.table.PrimaryKey().Get(ctx, &market, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &market, nil
}

func (this marketTable) HasByCreditTypeAbbrevBankDenom(ctx context.Context, credit_type_abbrev string, bank_denom string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		credit_type_abbrev,
		bank_denom,
	)
}

func (this marketTable) GetByCreditTypeAbbrevBankDenom(ctx context.Context, credit_type_abbrev string, bank_denom string) (*Market, error) {
	var market Market
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &market,
		credit_type_abbrev,
		bank_denom,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &market, nil
}

func (this marketTable) List(ctx context.Context, prefixKey MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return MarketIterator{it}, err
}

func (this marketTable) ListRange(ctx context.Context, from, to MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return MarketIterator{it}, err
}

func (this marketTable) DeleteBy(ctx context.Context, prefixKey MarketIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this marketTable) DeleteRange(ctx context.Context, from, to MarketIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this marketTable) doNotImplement() {}

var _ MarketTable = marketTable{}

func NewMarketTable(db ormtable.Schema) (MarketTable, error) {
	table := db.GetTable(&Market{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Market{}).ProtoReflect().Descriptor().FullName()))
	}
	return marketTable{table.(ormtable.AutoIncrementTable)}, nil
}

// singleton store
type FeeParamsTable interface {
	Get(ctx context.Context) (*FeeParams, error)
	Save(ctx context.Context, feeParams *FeeParams) error
}

type feeParamsTable struct {
	table ormtable.Table
}

var _ FeeParamsTable = feeParamsTable{}

func (x feeParamsTable) Get(ctx context.Context) (*FeeParams, error) {
	feeParams := &FeeParams{}
	_, err := x.table.Get(ctx, feeParams)
	return feeParams, err
}

func (x feeParamsTable) Save(ctx context.Context, feeParams *FeeParams) error {
	return x.table.Save(ctx, feeParams)
}

func NewFeeParamsTable(db ormtable.Schema) (FeeParamsTable, error) {
	table := db.GetTable(&FeeParams{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&FeeParams{}).ProtoReflect().Descriptor().FullName()))
	}
	return &feeParamsTable{table}, nil
}

type StateStore interface {
	SellOrderTable() SellOrderTable
	AllowedDenomTable() AllowedDenomTable
	MarketTable() MarketTable
	FeeParamsTable() FeeParamsTable

	doNotImplement()
}

type stateStore struct {
	sellOrder    SellOrderTable
	allowedDenom AllowedDenomTable
	market       MarketTable
	feeParams    FeeParamsTable
}

func (x stateStore) SellOrderTable() SellOrderTable {
	return x.sellOrder
}

func (x stateStore) AllowedDenomTable() AllowedDenomTable {
	return x.allowedDenom
}

func (x stateStore) MarketTable() MarketTable {
	return x.market
}

func (x stateStore) FeeParamsTable() FeeParamsTable {
	return x.feeParams
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	sellOrderTable, err := NewSellOrderTable(db)
	if err != nil {
		return nil, err
	}

	allowedDenomTable, err := NewAllowedDenomTable(db)
	if err != nil {
		return nil, err
	}

	marketTable, err := NewMarketTable(db)
	if err != nil {
		return nil, err
	}

	feeParamsTable, err := NewFeeParamsTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		sellOrderTable,
		allowedDenomTable,
		marketTable,
		feeParamsTable,
	}, nil
}
