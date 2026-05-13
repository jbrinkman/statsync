package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jbrinkman/statsync/internal/config"
	"github.com/jbrinkman/statsync/internal/model"
	"github.com/jbrinkman/statsync/internal/store"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tracked projects",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.New(config.ValkeyAddr())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer s.Close()

		ctx := context.Background()
		projects, err := s.ListProjects(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		typeFilter, _ := cmd.Flags().GetString("type")

		var filtered []*model.Project
		for _, p := range projects {
			if typeFilter != "" && p.Type != typeFilter {
				continue
			}
			filtered = append(filtered, p)
		}

		if len(filtered) == 0 {
			fmt.Println("No projects found.")
			return
		}

		for _, p := range filtered {
			status := p.ComputedStatus()
			itemCount := len(p.WorkItems)
			fmt.Printf("%-30s %-15s %-15s %d items\n", p.Name, p.Type, status, itemCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().String("type", "", "Filter by project type")
}
