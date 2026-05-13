package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jbrinkman/statsync/internal/model"
)

func (s *Store) SaveProject(ctx context.Context, p *model.Project) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("failed to marshal project: %w", err)
	}
	cmd := s.client.B().JsonSet().Key(p.ValkeyKey()).Path("$").Value(string(data)).Build()
	return s.client.Do(ctx, cmd).Error()
}

func (s *Store) GetProject(ctx context.Context, slug string) (*model.Project, error) {
	key := "project:" + slug
	cmd := s.client.B().JsonGet().Key(key).Build()
	result, err := s.client.Do(ctx, cmd).ToString()
	if err != nil {
		return nil, fmt.Errorf("project %q not found", slug)
	}

	// JSON.GET returns an array wrapper at root path
	var projects []model.Project
	if err := json.Unmarshal([]byte(result), &projects); err != nil {
		// Try without array wrapper
		var p model.Project
		if err2 := json.Unmarshal([]byte(result), &p); err2 != nil {
			return nil, fmt.Errorf("failed to unmarshal project: %w", err)
		}
		return &p, nil
	}
	if len(projects) == 0 {
		return nil, fmt.Errorf("project %q not found", slug)
	}
	return &projects[0], nil
}

func (s *Store) DeleteProject(ctx context.Context, slug string) error {
	key := "project:" + slug
	cmd := s.client.B().JsonDel().Key(key).Build()
	n, err := s.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return fmt.Errorf("failed to delete project %q: %w", slug, err)
	}
	if n == 0 {
		return fmt.Errorf("project %q not found", slug)
	}
	return nil
}

func (s *Store) ProjectExists(ctx context.Context, slug string) (bool, error) {
	key := "project:" + slug
	cmd := s.client.B().Exists().Key(key).Build()
	n, err := s.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (s *Store) ListProjects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project
	var cursor uint64

	for {
		cmd := s.client.B().Scan().Cursor(cursor).Match("project:*").Count(100).Build()
		resp := s.client.Do(ctx, cmd)

		entry, err := resp.AsScanEntry()
		if err != nil {
			return nil, fmt.Errorf("failed to scan projects: %w", err)
		}

		for _, key := range entry.Elements {
			getCmd := s.client.B().JsonGet().Key(key).Build()
			result, err := s.client.Do(ctx, getCmd).ToString()
			if err != nil {
				continue
			}
			var items []model.Project
			if err := json.Unmarshal([]byte(result), &items); err != nil {
				var p model.Project
				if err2 := json.Unmarshal([]byte(result), &p); err2 != nil {
					continue
				}
				projects = append(projects, &p)
				continue
			}
			if len(items) > 0 {
				projects = append(projects, &items[0])
			}
		}

		cursor = entry.Cursor
		if cursor == 0 {
			break
		}
	}

	return projects, nil
}
