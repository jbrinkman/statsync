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

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.New(config.ValkeyAddr())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer s.Close()

		name := args[0]
		slug := model.Slugify(name)
		ctx := context.Background()

		exists, err := s.ProjectExists(ctx, slug)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if exists {
			fmt.Fprintf(os.Stderr, "Error: project %q already exists\n", slug)
			os.Exit(1)
		}

		projType, _ := cmd.Flags().GetString("type")
		framework, _ := cmd.Flags().GetString("framework")
		assignee, _ := cmd.Flags().GetString("assignee")

		p := &model.Project{
			Name:      name,
			Slug:      slug,
			Type:      projType,
			Framework: framework,
			Assignee:  assignee,
			WorkItems: []model.WorkItem{},
			UpdatedAt: time.Now(),
		}

		if err := s.SaveProject(ctx, p); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("Added project %q (%s)\n", name, slug)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("type", "", "Project type (e.g., integration, documentation)")
	addCmd.Flags().String("framework", "", "Target framework name")
	addCmd.Flags().String("assignee", "", "Person assigned to this project")
}
