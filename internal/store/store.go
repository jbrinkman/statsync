package store

import (
	"context"
	"fmt"

	"github.com/valkey-io/valkey-go"
)

type Store struct {
	client valkey.Client
}

func New(addr string) (*Store, error) {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to valkey at %s: %w", addr, err)
	}

	// Verify connection
	if err := client.Do(context.Background(), client.B().Ping().Build()).Error(); err != nil {
		client.Close()
		return nil, fmt.Errorf("valkey at %s is unreachable: %w", addr, err)
	}

	return &Store{client: client}, nil
}

func (s *Store) Close() {
	s.client.Close()
}

func (s *Store) Client() valkey.Client {
	return s.client
}
