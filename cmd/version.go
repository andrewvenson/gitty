package cmd

import (
	"fmt"
	//"os"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version of gitty",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("gowhere v1.1")
	},
}
