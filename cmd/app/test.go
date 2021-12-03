package app

import (
	"demo/pkg/service/test"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	name string
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test program",
	Long:  `test program, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please run test command, watch detail with -h args")
	},
}

// queryUserCmd represents the query-user command
var queryUserCmd = &cobra.Command{
	Use:   "query-user",
	Short: "query-user program",
	Long:  `query-user program, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		testSrv := test.NewTestService()
		testSrv.Run(name)
	},
}

func parseFlags(cmd *cobra.Command, flags []string) {
	flagMap := map[string]func(cmd *cobra.Command){
		"name": func(cmd *cobra.Command) {
			cmd.Flags().StringVarP(&name, "name", "n", name, "query user with name")
		},
	}
	for _, v := range flags {
		if f, ok := flagMap[v]; ok {
			f(cmd)
		}
	}
}

func init() {
	rootCmd.AddCommand(testCmd)

	parseFlags(queryUserCmd, []string{"name"})
	testCmd.AddCommand(queryUserCmd)
}
