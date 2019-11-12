package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Listen on TCP ports",
	Long:  `Listen on ports specified`,
	Run: func(cmd *cobra.Command, args []string) {
		openports(viper.GetStringSlice("port"))
	},
}
