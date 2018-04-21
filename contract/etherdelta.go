// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EtherDeltaABI is the input ABI used to generate the binding from.
const EtherDeltaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderFills\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"amountFilled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeMake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeMake\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeMake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"name\":\"changeFeeRebate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"testTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeAccount_\",\"type\":\"address\"}],\"name\":\"changeFeeAccount\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeRebate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeTake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeTake\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"}],\"name\":\"changeAccountLevelsAddr\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountLevelsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"availableVolume\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"},{\"name\":\"feeAccount_\",\"type\":\"address\"},{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"},{\"name\":\"feeMake_\",\"type\":\"uint256\"},{\"name\":\"feeTake_\",\"type\":\"uint256\"},{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Order\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"v\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"r\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"give\",\"type\":\"address\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// EtherDelta is an auto generated Go binding around an Ethereum contract.
type EtherDelta struct {
	EtherDeltaCaller     // Read-only binding to the contract
	EtherDeltaTransactor // Write-only binding to the contract
	EtherDeltaFilterer   // Log filterer for contract events
}

// EtherDeltaCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherDeltaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherDeltaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherDeltaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherDeltaSession struct {
	Contract     *EtherDelta       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherDeltaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherDeltaCallerSession struct {
	Contract *EtherDeltaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherDeltaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherDeltaTransactorSession struct {
	Contract     *EtherDeltaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherDeltaRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherDeltaRaw struct {
	Contract *EtherDelta // Generic contract binding to access the raw methods on
}

// EtherDeltaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherDeltaCallerRaw struct {
	Contract *EtherDeltaCaller // Generic read-only contract binding to access the raw methods on
}

// EtherDeltaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherDeltaTransactorRaw struct {
	Contract *EtherDeltaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherDelta creates a new instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDelta(address common.Address, backend bind.ContractBackend) (*EtherDelta, error) {
	contract, err := bindEtherDelta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherDelta{EtherDeltaCaller: EtherDeltaCaller{contract: contract}, EtherDeltaTransactor: EtherDeltaTransactor{contract: contract}, EtherDeltaFilterer: EtherDeltaFilterer{contract: contract}}, nil
}

// NewEtherDeltaCaller creates a new read-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaCaller(address common.Address, caller bind.ContractCaller) (*EtherDeltaCaller, error) {
	contract, err := bindEtherDelta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaCaller{contract: contract}, nil
}

// NewEtherDeltaTransactor creates a new write-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherDeltaTransactor, error) {
	contract, err := bindEtherDelta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaTransactor{contract: contract}, nil
}

// NewEtherDeltaFilterer creates a new log filterer instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherDeltaFilterer, error) {
	contract, err := bindEtherDelta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaFilterer{contract: contract}, nil
}

// bindEtherDelta binds a generic wrapper to an already deployed contract.
func bindEtherDelta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherDeltaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.EtherDeltaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transact(opts, method, params...)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) AccountLevelsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "accountLevelsAddr")
	return *ret0, err
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AmountFilled(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "amountFilled", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AvailableVolume(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "availableVolume", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeAccount")
	return *ret0, err
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeMake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeMake")
	return *ret0, err
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeRebate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeRebate")
	return *ret0, err
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeTake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeTake")
	return *ret0, err
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) OrderFills(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orderFills", arg0, arg1)
	return *ret0, err
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) Orders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orders", arg0, arg1)
	return *ret0, err
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) TestTrade(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "testTrade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
	return *ret0, err
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "tokens", arg0, arg1)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactor) CancelOrder(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "cancelOrder", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactorSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAccountLevelsAddr(opts *bind.TransactOpts, accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAccountLevelsAddr", accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAdmin", admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeAccount(opts *bind.TransactOpts, feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeAccount", feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeMake(opts *bind.TransactOpts, feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeMake", feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeRebate(opts *bind.TransactOpts, feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeRebate", feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeTake(opts *bind.TransactOpts, feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeTake", feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactorSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) DepositToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "depositToken", token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Order(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "order", tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Trade(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "trade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// EtherDeltaCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the EtherDelta contract.
type EtherDeltaCancelIterator struct {
	Event *EtherDeltaCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherDeltaCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherDeltaCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherDeltaCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaCancel represents a Cancel event raised by the EtherDelta contract.
type EtherDeltaCancel struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Expires    *big.Int
	Nonce      *big.Int
	User       common.Address
	V          uint8
	R          [32]byte
	S          [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0.
//
// Solidity: event Cancel(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32)
func (_EtherDelta *EtherDeltaFilterer) FilterCancel(opts *bind.FilterOpts) (*EtherDeltaCancelIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaCancelIterator{contract: _EtherDelta.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0.
//
// Solidity: event Cancel(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32)
func (_EtherDelta *EtherDeltaFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *EtherDeltaCancel) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaCancel)
				if err := _EtherDelta.contract.UnpackLog(event, "Cancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EtherDeltaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EtherDelta contract.
type EtherDeltaDepositIterator struct {
	Event *EtherDeltaDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherDeltaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherDeltaDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherDeltaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaDeposit represents a Deposit event raised by the EtherDelta contract.
type EtherDeltaDeposit struct {
	Token   common.Address
	User    common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) FilterDeposit(opts *bind.FilterOpts) (*EtherDeltaDepositIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaDepositIterator{contract: _EtherDelta.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EtherDeltaDeposit) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaDeposit)
				if err := _EtherDelta.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EtherDeltaOrderIterator is returned from FilterOrder and is used to iterate over the raw logs and unpacked data for Order events raised by the EtherDelta contract.
type EtherDeltaOrderIterator struct {
	Event *EtherDeltaOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherDeltaOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherDeltaOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherDeltaOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaOrder represents a Order event raised by the EtherDelta contract.
type EtherDeltaOrder struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Expires    *big.Int
	Nonce      *big.Int
	User       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrder is a free log retrieval operation binding the contract event 0x3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e85.
//
// Solidity: event Order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address)
func (_EtherDelta *EtherDeltaFilterer) FilterOrder(opts *bind.FilterOpts) (*EtherDeltaOrderIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Order")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaOrderIterator{contract: _EtherDelta.contract, event: "Order", logs: logs, sub: sub}, nil
}

// WatchOrder is a free log subscription operation binding the contract event 0x3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e85.
//
// Solidity: event Order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address)
func (_EtherDelta *EtherDeltaFilterer) WatchOrder(opts *bind.WatchOpts, sink chan<- *EtherDeltaOrder) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Order")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaOrder)
				if err := _EtherDelta.contract.UnpackLog(event, "Order", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EtherDeltaTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the EtherDelta contract.
type EtherDeltaTradeIterator struct {
	Event *EtherDeltaTrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherDeltaTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaTrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherDeltaTrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherDeltaTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaTrade represents a Trade event raised by the EtherDelta contract.
type EtherDeltaTrade struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Get        common.Address
	Give       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d.
//
// Solidity: event Trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, get address, give address)
func (_EtherDelta *EtherDeltaFilterer) FilterTrade(opts *bind.FilterOpts) (*EtherDeltaTradeIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaTradeIterator{contract: _EtherDelta.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d.
//
// Solidity: event Trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, get address, give address)
func (_EtherDelta *EtherDeltaFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *EtherDeltaTrade) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaTrade)
				if err := _EtherDelta.contract.UnpackLog(event, "Trade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EtherDeltaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EtherDelta contract.
type EtherDeltaWithdrawIterator struct {
	Event *EtherDeltaWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherDeltaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherDeltaWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherDeltaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaWithdraw represents a Withdraw event raised by the EtherDelta contract.
type EtherDeltaWithdraw struct {
	Token   common.Address
	User    common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) FilterWithdraw(opts *bind.FilterOpts) (*EtherDeltaWithdrawIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaWithdrawIterator{contract: _EtherDelta.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EtherDeltaWithdraw) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaWithdraw)
				if err := _EtherDelta.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
