package web3

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pinardZ/debox/bscserver/web3/erc20"
	"math/big"
)

type Token interface {
	Balance(walletAddressStr string) (*big.Int, error)
}

func NewToken(tokenAddrStr string, backend bind.ContractBackend) (Token, error) {
	tokenAddr := common.HexToAddress(tokenAddrStr)
	erc20Token, err := erc20.NewErc20(tokenAddr, backend)
	return token{erc20Token}, err
}

type token struct {
	*erc20.Erc20
}

func (t token) Balance(walletAddressStr string) (*big.Int, error) {
	return t.BalanceOf(nil, common.HexToAddress(walletAddressStr))
}
