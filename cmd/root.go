package cmd

import (
	"github.com/spf13/cobra"
)

var (
	nodeURL string
)

var rootCmd = &cobra.Command{
	Use:   "tia",
	Short: "A CLI explorer for Celestia blockchain",
	Long: `tia is a command-line interface for exploring the Celestia blockchain.
It allows you to retrieve information about blocks, transactions, and more.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&nodeURL, "node", "https://rpc.celestia.pops.one", "Celestia node URL")
}
