package cmd

import (
	"fmt"
	"log"

	jsonnet "github.com/google/go-jsonnet"
	"github.com/spf13/cobra"

	tankaNative "github.com/grafana/tanka/pkg/jsonnet/native"
)

var options struct {
	extCodeFiles map[string]string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "jrender <file>",
	Short:      "Extended Jsonnet renderer",
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"file"},
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		vm := jsonnet.MakeVM()
		for _, nf := range tankaNative.Funcs() {
			vm.NativeFunction(nf)
		}

		for k, v := range options.extCodeFiles {
			contents, _, err := vm.ImportAST(v, v)
			if err != nil {
				log.Fatal(err)
			}

			vm.ExtNode(k, contents)
		}

		result, err := vm.EvaluateFile(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringToStringVar(&options.extCodeFiles, "ext-code-file", map[string]string{}, "Read the code from the file")
}
