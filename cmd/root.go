package cmd

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "open",
		Short: "openports, get nauti",
		Long:  `openports is a firewall testing tool. It allows you to create a range of listening ports.`,
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("range", "r", "int-int", "Port Range")
	rootCmd.PersistentFlags().IntP("first", "f", 80, "First Port")
	rootCmd.PersistentFlags().IntP("last", "l", 0, "Last Port")
	viper.BindPFlag("range", rootCmd.PersistentFlags().Lookup("range"))
	viper.BindPFlag("first", rootCmd.PersistentFlags().Lookup("first"))
	viper.BindPFlag("last", rootCmd.PersistentFlags().Lookup("last"))
}

func getRange(r string) []int {
	var pRange []int
	if r != "int-int" {
		rSplit := strings.Split(r, "-")
		if len(rSplit) > 2 {
			return []int{80}
		}
		for _, a := range rSplit {
			if i, err := strconv.Atoi(a); err == nil {
				pRange = append(pRange, i)
			}
		}
	}
	return pRange
}

func start() {
	portRange := getRange(viper.GetString("range"))
	if len(portRange) < 1 {
		portRange = []int{viper.GetInt("first"), viper.GetInt("last")}
	}
	openports(portRange)
}
