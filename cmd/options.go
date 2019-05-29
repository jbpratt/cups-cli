package cmd

import (
	"fmt"
	"log"

	cups "github.com/jbpratt78/go-cups"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Options from the printer's attribute list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("not enough args")
		}
		conn := cups.NewDefaultConnection()
		currDest := conn.Dests[0]
		arg := args[0]
		if arg == "all" {
			fmt.Println("Current options for ", currDest.Name)
			for k, v := range currDest.GetOptions() {
				fmt.Printf("\t%s = %s\n", k, v)
			}
		} else {
			//fmt.Println("Current option: ", arg, currDest.GetOption(arg))
			option, err := currDest.GetOption(arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Current option for", arg, option)
		}
	},
}

func init() {
	rootCmd.AddCommand(optionsCmd)
}
