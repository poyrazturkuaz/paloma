// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/crypto/types"
)

// StakingValidatorI is an autogenerated mock type for the ValidatorI type
type StakingValidatorI struct {
	mock.Mock
}

// ConsPubKey provides a mock function with given fields:
func (_m *StakingValidatorI) ConsPubKey() (types.PubKey, error) {
	ret := _m.Called()

	var r0 types.PubKey
	if rf, ok := ret.Get(0).(func() types.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PubKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBondedTokens provides a mock function with given fields:
func (_m *StakingValidatorI) GetBondedTokens() cosmos_sdktypes.Int {
	ret := _m.Called()

	var r0 cosmos_sdktypes.Int
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Int)
	}

	return r0
}

// GetCommission provides a mock function with given fields:
func (_m *StakingValidatorI) GetCommission() cosmos_sdktypes.Dec {
	ret := _m.Called()

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.Dec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	return r0
}

// GetConsAddr provides a mock function with given fields:
func (_m *StakingValidatorI) GetConsAddr() (cosmos_sdktypes.ConsAddress, error) {
	ret := _m.Called()

	var r0 cosmos_sdktypes.ConsAddress
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.ConsAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.ConsAddress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsensusPower provides a mock function with given fields: _a0
func (_m *StakingValidatorI) GetConsensusPower(_a0 cosmos_sdktypes.Int) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Int) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetDelegatorShares provides a mock function with given fields:
func (_m *StakingValidatorI) GetDelegatorShares() cosmos_sdktypes.Dec {
	ret := _m.Called()

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.Dec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	return r0
}

// GetMinSelfDelegation provides a mock function with given fields:
func (_m *StakingValidatorI) GetMinSelfDelegation() cosmos_sdktypes.Int {
	ret := _m.Called()

	var r0 cosmos_sdktypes.Int
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Int)
	}

	return r0
}

// GetMoniker provides a mock function with given fields:
func (_m *StakingValidatorI) GetMoniker() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOperator provides a mock function with given fields:
func (_m *StakingValidatorI) GetOperator() cosmos_sdktypes.ValAddress {
	ret := _m.Called()

	var r0 cosmos_sdktypes.ValAddress
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.ValAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.ValAddress)
		}
	}

	return r0
}

// GetStatus provides a mock function with given fields:
func (_m *StakingValidatorI) GetStatus() stakingtypes.BondStatus {
	ret := _m.Called()

	var r0 stakingtypes.BondStatus
	if rf, ok := ret.Get(0).(func() stakingtypes.BondStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(stakingtypes.BondStatus)
	}

	return r0
}

// GetTokens provides a mock function with given fields:
func (_m *StakingValidatorI) GetTokens() cosmos_sdktypes.Int {
	ret := _m.Called()

	var r0 cosmos_sdktypes.Int
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Int)
	}

	return r0
}

// IsBonded provides a mock function with given fields:
func (_m *StakingValidatorI) IsBonded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsJailed provides a mock function with given fields:
func (_m *StakingValidatorI) IsJailed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUnbonded provides a mock function with given fields:
func (_m *StakingValidatorI) IsUnbonded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUnbonding provides a mock function with given fields:
func (_m *StakingValidatorI) IsUnbonding() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SharesFromTokens provides a mock function with given fields: amt
func (_m *StakingValidatorI) SharesFromTokens(amt cosmos_sdktypes.Int) (cosmos_sdktypes.Dec, error) {
	ret := _m.Called(amt)

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Int) cosmos_sdktypes.Dec); ok {
		r0 = rf(amt)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Int) error); ok {
		r1 = rf(amt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SharesFromTokensTruncated provides a mock function with given fields: amt
func (_m *StakingValidatorI) SharesFromTokensTruncated(amt cosmos_sdktypes.Int) (cosmos_sdktypes.Dec, error) {
	ret := _m.Called(amt)

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Int) cosmos_sdktypes.Dec); ok {
		r0 = rf(amt)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Int) error); ok {
		r1 = rf(amt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TmConsPublicKey provides a mock function with given fields:
func (_m *StakingValidatorI) TmConsPublicKey() (crypto.PublicKey, error) {
	ret := _m.Called()

	var r0 crypto.PublicKey
	if rf, ok := ret.Get(0).(func() crypto.PublicKey); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(crypto.PublicKey)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokensFromShares provides a mock function with given fields: _a0
func (_m *StakingValidatorI) TokensFromShares(_a0 cosmos_sdktypes.Dec) cosmos_sdktypes.Dec {
	ret := _m.Called(_a0)

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Dec) cosmos_sdktypes.Dec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	return r0
}

// TokensFromSharesRoundUp provides a mock function with given fields: _a0
func (_m *StakingValidatorI) TokensFromSharesRoundUp(_a0 cosmos_sdktypes.Dec) cosmos_sdktypes.Dec {
	ret := _m.Called(_a0)

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Dec) cosmos_sdktypes.Dec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	return r0
}

// TokensFromSharesTruncated provides a mock function with given fields: _a0
func (_m *StakingValidatorI) TokensFromSharesTruncated(_a0 cosmos_sdktypes.Dec) cosmos_sdktypes.Dec {
	ret := _m.Called(_a0)

	var r0 cosmos_sdktypes.Dec
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Dec) cosmos_sdktypes.Dec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Dec)
	}

	return r0
}

type mockConstructorTestingTNewStakingValidatorI interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakingValidatorI creates a new instance of StakingValidatorI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakingValidatorI(t mockConstructorTestingTNewStakingValidatorI) *StakingValidatorI {
	mock := &StakingValidatorI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
