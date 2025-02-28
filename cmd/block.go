package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/21state/tia/pkg/rpc"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block [height]",
	Short: "Get block information",
	Long: `Retrieve information about a specific block by height.
If no height is provided, the latest block will be fetched.

Example:
  tia block 123456
  tia block latest`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := rpc.NewClient(nodeURL)
		ctx := context.Background()

		var height int64 = 0 // Default to latest block
		if len(args) > 0 && args[0] != "latest" {
			var err error
			height, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid block height: %v", err)
			}
		}

		block, err := client.GetBlock(ctx, height)
		if err != nil {
			return fmt.Errorf("failed to get block: %v", err)
		}
		
		fmt.Printf("Block #%d\n", block.Block.Header.Height)
		fmt.Printf("Hash: %s\n", block.BlockID.Hash)
		fmt.Printf("Time: %s\n", block.Block.Header.Time)
		fmt.Printf("Proposer: %s\n", block.Block.Header.ProposerAddress)
		fmt.Printf("Transactions: %d\n", len(block.Block.Data.Txs))
		fmt.Printf("App Hash: %s\n", block.Block.Header.AppHash)
		fmt.Printf("Consensus Hash: %s\n", block.Block.Header.ConsensusHash)
		fmt.Printf("Last Block ID Hash: %s\n", block.Block.Header.LastBlockID.Hash)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
