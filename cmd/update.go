package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jbrinkman/statsync/internal/config"
	"github.com/jbrinkman/statsync/internal/model"
	"github.com/jbrinkman/statsync/internal/store"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [slug]",
	Short: "Update a project or its work items",
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

		// Update project metadata
		if v, _ := cmd.Flags().GetString("assignee"); v != "" {
			p.Assignee = v
		}
		if v, _ := cmd.Flags().GetString("type"); v != "" {
			p.Type = v
		}
		if v, _ := cmd.Flags().GetString("framework"); v != "" {
			p.Framework = v
		}
		if v, _ := cmd.Flags().GetString("notes"); v != "" {
			p.Notes = v
		}

		// Add a new work item
		if addItem, _ := cmd.Flags().GetString("add-item"); addItem != "" {
			p.WorkItems = append(p.WorkItems, model.WorkItem{
				Label:  addItem,
				Status: model.StatusNotStarted,
				PRs:    []string{},
			})
		}

		// Update work item status or add PR
		if item, _ := cmd.Flags().GetString("item"); item != "" {
			found := false
			for i := range p.WorkItems {
				if p.WorkItems[i].Label == item {
					found = true
					if status, _ := cmd.Flags().GetString("status"); status != "" {
						normalized, err := model.NormalizeStatus(status)
						if err != nil {
							fmt.Fprintln(os.Stderr, err)
							os.Exit(1)
						}
						p.WorkItems[i].Status = normalized
					}
					if pr, _ := cmd.Flags().GetString("add-pr"); pr != "" {
						p.WorkItems[i].PRs = append(p.WorkItems[i].PRs, pr)
					}
					if ecd, _ := cmd.Flags().GetString("ecd"); ecd != "" {
						p.WorkItems[i].ECD = ecd
					}
					if jira, _ := cmd.Flags().GetString("jira-issue"); jira != "" {
						p.WorkItems[i].JiraIssue = jira
					}
					if gh, _ := cmd.Flags().GetString("github-issue"); gh != "" {
						p.WorkItems[i].GithubIssue = gh
					}
					if notes, _ := cmd.Flags().GetString("item-notes"); notes != "" {
						p.WorkItems[i].Notes = notes
					}
					break
				}
			}
			if !found {
				fmt.Fprintf(os.Stderr, "Error: work item %q not found in project %q\n", item, p.Slug)
				os.Exit(1)
			}
		}

		p.UpdatedAt = time.Now()

		if err := s.SaveProject(ctx, p); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("Updated project %q\n", p.Name)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("assignee", "", "Update assignee")
	updateCmd.Flags().String("type", "", "Update project type")
	updateCmd.Flags().String("framework", "", "Update framework")
	updateCmd.Flags().String("notes", "", "Update notes")
	updateCmd.Flags().String("add-item", "", "Add a new work item")
	updateCmd.Flags().String("item", "", "Work item to update")
	updateCmd.Flags().String("status", "", "New status for work item")
	updateCmd.Flags().String("add-pr", "", "Add PR link to work item")
	updateCmd.Flags().String("ecd", "", "Estimated completion date (YYYY-MM-DD)")
	updateCmd.Flags().String("jira-issue", "", "Jira issue key (e.g., PROJ-123)")
	updateCmd.Flags().String("github-issue", "", "GitHub issue URL")
	updateCmd.Flags().String("item-notes", "", "Notes for work item")
}
