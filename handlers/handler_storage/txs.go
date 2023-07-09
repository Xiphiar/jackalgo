package handler_storage

import (
	"fmt"

	"github.com/JackalLabs/jackalgo/utils"
	"github.com/cosmos/cosmos-sdk/types"
	storagetypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (s *StorageHandler) BuyStorage(forAddress string, duration int64, bytes int64) (*types.TxResponse, error) {
	if duration <= 0 {
		return nil, fmt.Errorf("cannot use less than 0 months of duration")
	}
	monthsAsHours := duration * 720

	_, err := types.AccAddressFromBech32(forAddress)
	if err != nil {
		return nil, err
	}

	buyMsg := storagetypes.MsgBuyStorage{
		Creator:      s.walletHandler.GetAddress(),
		PaymentDenom: "ujkl",
		ForAddress:   forAddress,
		Duration:     fmt.Sprintf("%dh", monthsAsHours),
		Bytes:        fmt.Sprintf("%d", utils.NumTo3xTB(bytes)),
	}
	res, err := s.walletHandler.SendTx(&buyMsg)

	return res, err
}

func (s *StorageHandler) UpgradeStorage(forAddress string, duration int64, bytes int64) (*types.TxResponse, error) {
	if duration <= 0 {
		return nil, fmt.Errorf("cannot use less than 0 months of duration")
	}
	monthsAsHours := duration * 720

	_, err := types.AccAddressFromBech32(forAddress)
	if err != nil {
		return nil, err
	}

	buyMsg := storagetypes.MsgUpgradeStorage{
		Creator:      s.walletHandler.GetAddress(),
		PaymentDenom: "ujkl",
		ForAddress:   forAddress,
		Duration:     fmt.Sprintf("%dh", monthsAsHours),
		Bytes:        fmt.Sprintf("%d", utils.NumTo3xTB(bytes)),
	}
	res, err := s.walletHandler.SendTx(&buyMsg)

	return res, err
}