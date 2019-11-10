package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		// Use:   "",
		Short: "openports, get nauti",
		Long:  `openports is a firewall testing tool. It allows you to create a range of listening ports.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// },
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("port", "p", []string{"0"}, "Port, Port range or Port list")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}
