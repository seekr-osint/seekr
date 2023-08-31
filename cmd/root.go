/*
Copyright © 2023 seekr-osint

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"embed"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/config"
	"github.com/seekr-osint/seekr/api/database"
	"github.com/seekr-osint/seekr/api/seekrauth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	ip      net.IP
	port    uint16
	webfs   embed.FS
	dbpath  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seekr",
	Short: "seekr osint tool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config{
			Ip:           ip,
			Port:         port,
			DataBasePath: dbpath,
		}
		database, err := database.Connect(config.DataBasePath)
		if err != nil {
			log.Panic("error initializing data base: ", err)
		}
		users := seekrauth.Users{
			{
				Username: "hacker",
				Password: "",
			},
			{
				Username: "seekr",
				Password: "",
			},
		}
		log.Panic(api.Serve(config, webfs, database, users))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(fs embed.FS) {
	webfs = fs
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/seekr/cfg.toml)")
	rootCmd.PersistentFlags().IPVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "ip adress to use")
	rootCmd.PersistentFlags().StringVar(&dbpath, "db", "./data.db", "Path to database")
	rootCmd.PersistentFlags().Uint16VarP(&port, "port", "p", 3000, "port")
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// err := viper.SafeWriteConfig()
	// // err := viper.WriteConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME/.config/seekr")
		viper.SetConfigType("toml")
		viper.SetConfigName("cfg")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
