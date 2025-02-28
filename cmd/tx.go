package cmd

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/21state/tia/pkg/rpc"
	"github.com/spf13/cobra"
)

// txCmd represents the tx command
var txCmd = &cobra.Command{
	Use:   "tx [hash]",
	Short: "Get transaction information",
	Long: `Retrieve information about a specific transaction by hash.

Example:
  tia tx 0x1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF
  tia tx 1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := rpc.NewClient(nodeURL)
		ctx := context.Background()

		// Clean the hash (remove 0x prefix if present)
		hash := strings.TrimPrefix(args[0], "0x")
		
		// Convert hex to bytes
		hashBytes, err := hex.DecodeString(hash)
		if err != nil {
			return fmt.Errorf("invalid transaction hash: %v", err)
		}

		tx, err := client.GetTx(ctx, hashBytes)
		if err != nil {
			return fmt.Errorf("failed to get transaction: %v", err)
		}

		// Display transaction information
		fmt.Printf("Transaction Hash: %s\n", hash)
		fmt.Printf("Height: %d\n", tx.Height)
		fmt.Printf("Index: %d\n", tx.Index)
		fmt.Printf("Result Code: %d\n", tx.TxResult.Code)
		
		if tx.TxResult.Code != 0 {
			fmt.Printf("Error: %s\n", tx.TxResult.Log)
		} else {
			fmt.Printf("Success: %s\n", tx.TxResult.Log)
		}
		
		// Display gas information
		fmt.Printf("Gas Wanted: %d\n", tx.TxResult.GasWanted)
		fmt.Printf("Gas Used: %d\n", tx.TxResult.GasUsed)
		
		// Display the raw transaction data
		fmt.Printf("Raw Transaction: %s\n", base64.StdEncoding.EncodeToString(tx.Tx))
		
		// Display events
		if len(tx.TxResult.Events) > 0 {
			fmt.Println("\nEvents:")
			for _, event := range tx.TxResult.Events {
				fmt.Printf("  Type: %s\n", event.Type)
				if len(event.Attributes) > 0 {
					fmt.Println("  Attributes:")
					for _, attr := range event.Attributes {
						key, err := base64.StdEncoding.DecodeString(string(attr.Key))
						if err != nil {
							key = attr.Key
						}
						
						value, err := base64.StdEncoding.DecodeString(string(attr.Value))
						if err != nil {
							value = attr.Value
						}
						
						fmt.Printf("    %s: %s\n", string(key), string(value))
					}
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
