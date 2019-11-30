package cmd

import (
	"github.com/jessecorson/openports/pkg/ports"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for open TCP ports",
	Long:  "Scan for open TCP ports. The -p flag can be used as a single port, a range, or a list",
	Run: func(cmd *cobra.Command, args []string) {
		ports.Scan(viper.GetString("target"), viper.GetStringSlice("port"))
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("target", "t", "localhost", "IP Address or HostName to scan")
	viper.BindPFlag("target", rootCmd.PersistentFlags().Lookup("target"))
}
