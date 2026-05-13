package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jbrinkman/statsync/internal/config"
	"github.com/jbrinkman/statsync/internal/store"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [slug]",
	Short: "Delete a project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.New(config.ValkeyAddr())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer s.Close()

		ctx := context.Background()
		if err := s.DeleteProject(ctx, args[0]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("Deleted project %q\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
