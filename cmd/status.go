package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/21state/tia/pkg/rpc"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get node status information",
	Long: `Retrieve the current status of the Celestia node.
This includes information about the node, sync status, and validator details.

Example:
  tia status`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := rpc.NewClient(nodeURL)
		ctx := context.Background()

		status, err := client.GetStatus(ctx)
		if err != nil {
			return fmt.Errorf("failed to get status: %v", err)
		}

		// Node Info
		fmt.Println("Node Information:")
		fmt.Printf("  Moniker:      %s\n", status.NodeInfo.Moniker)
		fmt.Printf("  Network:      %s\n", status.NodeInfo.Network)
		fmt.Printf("  Version:      %s\n", status.NodeInfo.Version)
		fmt.Printf("  Node ID:      %s\n", status.NodeInfo.ID)
		fmt.Printf("  Listen Addr:  %s\n", status.NodeInfo.ListenAddr)
		fmt.Printf("  RPC Address:  %s\n", status.NodeInfo.Other.RPCAddress)
		fmt.Printf("  Protocol:     P2P=%s, Block=%s, App=%s\n", 
			status.NodeInfo.ProtocolVersion.P2P,
			status.NodeInfo.ProtocolVersion.Block,
			status.NodeInfo.ProtocolVersion.App)

		// Sync Info
		fmt.Println("\nSync Status:")
		fmt.Printf("  Latest Block Height: %d\n", status.SyncInfo.LatestBlockHeight)
		fmt.Printf("  Latest Block Time:   %s\n", formatTime(status.SyncInfo.LatestBlockTime))
		fmt.Printf("  Catching Up:         %v\n", status.SyncInfo.CatchingUp)
		fmt.Printf("  Latest Block Hash:   %s\n", status.SyncInfo.LatestBlockHash)
		fmt.Printf("  Latest App Hash:     %s\n", status.SyncInfo.LatestAppHash)

		// Validator Info
		fmt.Println("\nValidator Information:")
		fmt.Printf("  Address:      %s\n", status.ValidatorInfo.Address)
		fmt.Printf("  Voting Power: %s\n", status.ValidatorInfo.VotingPower)
		fmt.Printf("  Pub Key:      %s (%s)\n", status.ValidatorInfo.PubKey.Value, status.ValidatorInfo.PubKey.Type)

		return nil
	},
}

func formatTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05 UTC")
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
