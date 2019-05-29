package cmd

import (
	"fmt"
	"log"

	cups "github.com/jbpratt78/go-cups"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print a file",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: optional use default, name as arg as well
		conn := cups.NewDefaultConnection()

		id, err := conn.Dests[0].PrintFile(filepath, "text/plain")
		if err != nil || id == -1 {
			log.Fatal(err)
		}

		fmt.Printf("JobID: %d\n", id)
	},
}

var filepath string

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.Flags().StringVarP(&filepath, "file", "f", "", "File path to print")
}
