package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

// HelloCmd represents the Hello command
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Welcome message",
	Long:  `Respond with a welcome message`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, friend! Updated test!")
	},
}
