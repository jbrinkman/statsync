package cmd

import (
	"fmt"
	"os"

	"github.com/jbrinkman/statsync/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "statsync",
	Short: "Track integration project status across multiple systems",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.Init(cfgFile)
	})
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.statsync.yaml)")
	rootCmd.PersistentFlags().String("valkey-addr", "", "Valkey server address")
}
