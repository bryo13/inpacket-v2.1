/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"inpacketv2.1/pkg/ips"
	"inpacketv2.1/pkg/nic"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "inpacketv2.1",
	Short: "Listen to interfaces in your computer",
	Long: `We listen to any or all interfaces in your computer
	and checks if the packet passing throuth the NIC is 
	from a public or private network.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// chooseInterface chooses which interface one wants to listen
var allInterface = &cobra.Command{
	Use:     "allInterfaces",
	Aliases: []string{"all"},
	Short:   "Shows all interface present in a system",
	Run: func(cmd *cobra.Command, args []string) {
		interfaces := nic.AllInterfaces()
		fmt.Println(interfaces)
	},
}

// selectOne watches the choosen interface
var selectOne = &cobra.Command{
	Use:     "select",
	Aliases: []string{"s"},
	Short:   "Watches the choosen interface",
	Run: func(cmd *cobra.Command, args []string) {
		allInterfaces := nic.AllInterfaces()
		chosen := nic.ChooseInterface(allInterfaces)
		fmt.Println(chosen, " has been chosen")
		ips.ReadPacket(chosen)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(allInterface)
	rootCmd.AddCommand(selectOne)
}
