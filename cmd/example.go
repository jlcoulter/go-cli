package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jlcoulter/go-cli-template/internal/example"
)

var exampleCmd = &cobra.Command{
	Use:   "example [name]",
	Short: "Run the example command",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := "world"
		if len(args) > 0 {
			name = args[0]
		}
		result := example.Greet(name)

		if jsonOut {
			fmt.Printf(`{"greeting":"%s"}`+"\n", result)
		} else {
			fmt.Println(result)
		}
		return nil
	},
}