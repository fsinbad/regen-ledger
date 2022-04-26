package server_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	ecocredittypes "github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/module"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	ff, ecocreditSubspace, bankKeeper, accountKeeper, distKeeper := setup(t)
	s := testsuite.NewIntegrationTestSuite(ff, ecocreditSubspace, bankKeeper, accountKeeper, distKeeper)
	suite.Run(t, s)
}

func TestGenesis(t *testing.T) {
	ff, ecocreditSubspace, bankKeeper, _, _ := setup(t)
	s := testsuite.NewGenesisTestSuite(ff, ecocreditSubspace, bankKeeper)
	suite.Run(t, s)
}

func setup(t *testing.T) (*server.FixtureFactory, paramstypes.Subspace, bankkeeper.BaseKeeper, authkeeper.AccountKeeper, *mocks.MockDistributionKeeper) {
	ff := server.NewFixtureFactory(t, 8)
	baseApp := ff.BaseApp()
	cdc := ff.Codec()
	amino := codec.NewLegacyAmino()

	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())
	params.RegisterInterfaces(cdc.InterfaceRegistry())

	authKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	distKey := sdk.NewKVStoreKey(disttypes.StoreKey)
	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	baseApp.MountStore(authKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(bankKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(distKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(paramsKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(tkey, sdk.StoreTypeTransient)

	authSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, authtypes.ModuleName)
	bankSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, banktypes.ModuleName)
	ecocreditSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, ecocredittypes.ModuleName)

	maccPerms := map[string][]string{
		minttypes.ModuleName:       {authtypes.Minter},
		ecocredittypes.ModuleName:  {authtypes.Burner},
		basket.BasketSubModuleName: {authtypes.Burner, authtypes.Minter},
	}

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc, authKey, authSubspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc, bankKey, accountKeeper, bankSubspace, nil,
	)

	ctrl := gomock.NewController(t)
	distKeeper := mocks2.NewMockDistributionKeeper(ctrl)
	ecocreditModule := ecocredit.NewModule(ecocreditSubspace, accountKeeper, bankKeeper, distKeeper)
	ff.SetModules([]module.Module{ecocreditModule})

	return ff, ecocreditSubspace, bankKeeper, accountKeeper, distKeeper
}
