package model

import (
	"regexp"
	"strings"
	"time"
)

type Resource struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type WorkItem struct {
	Label       string     `json:"label"`
	Status      Status     `json:"status"`
	PRs         []string   `json:"prs"`
	ECD         string     `json:"ecd,omitempty"`
	JiraIssue   string     `json:"jiraIssue,omitempty"`
	GithubIssue string     `json:"githubIssue,omitempty"`
	Notes       string     `json:"notes,omitempty"`
	Assignee    string     `json:"assignee,omitempty"`
	Release     string     `json:"release,omitempty"`
	Resources   []Resource `json:"resources,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type Project struct {
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Category  string     `json:"category"`
	Framework string     `json:"framework"`
	Assignee  string     `json:"assignee"`
	WorkItems []WorkItem `json:"workItems"`
	Notes     string     `json:"notes"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts a project name to a kebab-case slug.
func Slugify(name string) string {
	lower := strings.ToLower(strings.TrimSpace(name))
	slug := nonAlphanumeric.ReplaceAllString(lower, "-")
	return strings.Trim(slug, "-")
}

// Status returns the computed project status based on work items.
func (p *Project) ComputedStatus() Status {
	return ComputeProjectStatus(p.WorkItems)
}

// ValkeyKey returns the Valkey key for this project.
func (p *Project) ValkeyKey() string {
	return "project:" + p.Slug
}
