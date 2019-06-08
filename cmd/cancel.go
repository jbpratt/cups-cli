package cmd

import (
	"log"
	"strconv"

	cups "github.com/jbpratt78/go-cups"
	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a print job by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("not enough args")
		}
		conn := cups.NewDefaultConnection()
		currDest := conn.Dests[0]
		arg, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = currDest.CancelJob(arg)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}
