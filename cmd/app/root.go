package app

import (
	"demo/lib/env"
	"fmt"
	"github.com/json-iterator/go/extra"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo job",
	Long:  `demo job service collection`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	extra.RegisterFuzzyDecoders()

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := config.Load(
		file.NewSource(file.WithPath(filepath.Join("config", env.Mode(), "mysql.yaml"))),
	)
	if err != nil {
		panic(err)
	}
}
