package cmd

import (
	"echo.go.dev/pkg/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "The main app command",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Arch: %s\n", runtime.GOARCH)
		fmt.Printf("CPUs: %d\n\n", runtime.NumCPU())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.Path, "config", "c", "", "config file (default is config.toml)")
	rootCmd.AddCommand(cmdServer)
	rootCmd.AddCommand(cmdCreateUser)
	rootCmd.AddCommand(cmdSetPassword)
	rootCmd.AddCommand(cmdMigrate)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
