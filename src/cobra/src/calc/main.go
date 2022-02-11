package calc

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var Cmd = &cobra.Command{
	Use:   "calc",
	Short: "calc cmd",
	Long:  `calc cmd for cobra test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := initService(cmd)
		if err != nil {
			fmt.Printf("Failed to init err = %s", err)
		}
		c, err := readConfigFile(configFile)
		if err != nil {
			fmt.Printf("Failed to read config file , configFile = %s , err = %s\n", configFile, err)
		}
		fmt.Println(c)
		fmt.Print(c.Num1 + c.Num2)
		return nil
	},
}

func init() {
	Cmd.Flags().String("config", "", "transfer config file path")
}

func initService(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	configFile = viper.GetString("config")
	return nil
}
