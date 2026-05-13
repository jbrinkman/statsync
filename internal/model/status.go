package model

import (
	"fmt"
	"strings"
)

type Status string

const (
	StatusNotStarted   Status = "Not Started"
	StatusInProgress   Status = "In Progress"
	StatusInReview     Status = "In Review"
	StatusSubmitted    Status = "Submitted"
	StatusAwaitingMerge Status = "Awaiting Merge"
	StatusMerged       Status = "Merged"
	StatusDone         Status = "Done"
	StatusBlocked      Status = "Blocked"
	StatusPaused       Status = "Paused"
	StatusClosed       Status = "Closed"
	StatusNotABug      Status = "Not a Bug"
	StatusDropped      Status = "Dropped"
)

var AllStatuses = []Status{
	StatusNotStarted, StatusInProgress, StatusInReview,
	StatusSubmitted, StatusAwaitingMerge, StatusMerged,
	StatusDone, StatusBlocked, StatusPaused, StatusClosed, StatusNotABug,
	StatusDropped,
}

// statusLevel maps statuses to their progression level.
// Synonyms share the same level.
var statusLevel = map[Status]int{
	StatusNotStarted:    1,
	StatusInProgress:    2,
	StatusInReview:      3,
	StatusSubmitted:     4,
	StatusAwaitingMerge: 4,
	StatusMerged:        5,
	StatusDone:          6,
	StatusClosed:        6,
	StatusNotABug:       6,
	StatusDropped:       6,
}

// NormalizeStatus accepts a case-insensitive status string and returns the canonical form.
func NormalizeStatus(s string) (Status, error) {
	lower := strings.ToLower(strings.TrimSpace(s))
	for _, st := range AllStatuses {
		if strings.ToLower(string(st)) == lower {
			return st, nil
		}
	}
	return "", fmt.Errorf("invalid status %q, valid statuses: %s", s, statusList())
}

func statusList() string {
	names := make([]string, len(AllStatuses))
	for i, s := range AllStatuses {
		names[i] = string(s)
	}
	return strings.Join(names, ", ")
}

// ComputeProjectStatus derives a project's status from its work items.
func ComputeProjectStatus(items []WorkItem) Status {
	if len(items) == 0 {
		return StatusNotStarted
	}

	for _, item := range items {
		if item.Status == StatusBlocked {
			return StatusBlocked
		}
	}
	for _, item := range items {
		if item.Status == StatusPaused {
			return StatusPaused
		}
	}

	allNotStarted := true
	lowestLevel := 7 // higher than any real level

	for _, item := range items {
		if item.Status != StatusNotStarted {
			allNotStarted = false
			level := statusLevel[item.Status]
			if level < lowestLevel {
				lowestLevel = level
			}
		} else {
			// Treat Not Started as In Progress for lowest calculation
			if statusLevel[StatusInProgress] < lowestLevel {
				lowestLevel = statusLevel[StatusInProgress]
			}
		}
	}

	if allNotStarted {
		return StatusNotStarted
	}

	// Map level back to a representative status
	for _, st := range []Status{StatusInProgress, StatusInReview, StatusSubmitted, StatusMerged, StatusDone} {
		if statusLevel[st] == lowestLevel {
			return st
		}
	}

	return StatusInProgress
}
