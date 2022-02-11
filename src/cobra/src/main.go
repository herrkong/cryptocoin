package cobra

import (
	"cryptocoin/src/cobra/src/calc"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "cobra cmd",
	Short: "add cobra cmd",
	Long:  `show how to use cobra`,
}

func init() {
	cmd.AddCommand(calc.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
