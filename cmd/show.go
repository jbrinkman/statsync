package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jbrinkman/statsync/internal/config"
	"github.com/jbrinkman/statsync/internal/store"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [slug]",
	Short: "Show project details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.New(config.ValkeyAddr())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer s.Close()

		ctx := context.Background()
		p, err := s.GetProject(ctx, args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: project %q not found\n", args[0])
			os.Exit(1)
		}

		fmt.Printf("Name:      %s\n", p.Name)
		fmt.Printf("Slug:      %s\n", p.Slug)
		fmt.Printf("Category:  %s\n", p.Category)
		fmt.Printf("Framework: %s\n", p.Framework)
		fmt.Printf("Assignee:  %s\n", p.Assignee)
		fmt.Printf("Status:    %s\n", p.ComputedStatus())
		fmt.Printf("Created:   %s\n", p.CreatedAt.Format("2006-01-02"))
		fmt.Printf("Updated:   %s\n", p.UpdatedAt.Format("2006-01-02"))

		if len(p.WorkItems) > 0 {
			fmt.Println("\nWork Items:")
			for _, item := range p.WorkItems {
				assignee := item.Assignee
				if assignee == "" {
					assignee = p.Assignee
				}
				fmt.Printf("  - %-25s %-15s (%s)\n", item.Label, item.Status, assignee)
				if item.ECD != "" {
					fmt.Printf("    ECD: %s\n", item.ECD)
				}
				if item.Release != "" {
					fmt.Printf("    Release: %s\n", item.Release)
				}
				if item.JiraIssue != "" {
					fmt.Printf("    Jira: %s\n", item.JiraIssue)
				}
				if item.GithubIssue != "" {
					fmt.Printf("    GitHub: %s\n", item.GithubIssue)
				}
				if len(item.PRs) > 0 {
					fmt.Printf("    PRs: %s\n", strings.Join(item.PRs, ", "))
				}
				if len(item.Resources) > 0 {
					for _, r := range item.Resources {
						fmt.Printf("    Resource: %s - %s\n", r.Label, r.URL)
					}
				}
				if item.Notes != "" {
					fmt.Printf("    Notes: %s\n", item.Notes)
				}
			}
		}

		if p.Notes != "" {
			fmt.Printf("\nNotes: %s\n", p.Notes)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
