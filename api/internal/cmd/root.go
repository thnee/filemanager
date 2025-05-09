package cmd

import (
	"fmt"
	"log"
	"main/internal/g"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "fm",
	Short: "Filemanager",
	Long:  `Web file management application`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file location")
}

func initConfig() {
	// If user specified a config file as CLI param, then use that.
	if configFile != "" {
		viper.SetConfigFile(configFile)

		// Use the default location and name.
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.UnmarshalKey("users", &g.Users)
	viper.UnmarshalKey("file-areas", &g.FileAreas)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
