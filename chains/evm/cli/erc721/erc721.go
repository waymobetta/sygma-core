package erc721

import (
	"fmt"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/erc721"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/flags"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmclient"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmgaspricer"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmtransaction"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var ERC721Cmd = &cobra.Command{
	Use:   "erc721",
	Short: "ERC721-related instructions",
	Long:  "ERC721-related instructions",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// fetch global flag values
		url, gasLimit, gasPrice, senderKeyPair, err = flags.GlobalFlagValues(cmd)
		if err != nil {
			return fmt.Errorf("could not get global flags: %v", err)
		}
		return nil
	},
}

func init() {
	ERC721Cmd.AddCommand(mintCmd)
	ERC721Cmd.AddCommand(approveCmd)
	ERC721Cmd.AddCommand(ownerCmd)
	ERC721Cmd.AddCommand(depositCmd)
	ERC721Cmd.AddCommand(addMinterCmd)
}

func initializeErc721Contract() (*erc721.ERC721Contract, error) {
	txFabric := evmtransaction.NewTransaction

	ethClient, err := evmclient.NewEVMClientFromParams(
		url, senderKeyPair.PrivateKey())
	if err != nil {
		log.Error().Err(fmt.Errorf("eth client intialization error: %v", err))
		return nil, err
	}

	gasPricer := evmgaspricer.NewLondonGasPriceClient(
		ethClient,
		&evmgaspricer.GasPricerOpts{UpperLimitFeePerGas: gasPrice},
	)

	transactor := transactor.NewSignAndSendTransactor(txFabric, gasPricer, ethClient)
	erc721Contract := erc721.NewErc721Contract(ethClient, erc721Addr, transactor)

	return erc721Contract, nil
}
