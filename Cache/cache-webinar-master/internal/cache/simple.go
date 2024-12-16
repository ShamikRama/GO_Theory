package cache

import (
	"context"
	"sync"

	"github.com/glebnaz/cache-webinar/internal/model"
)

type Simple struct {
	mu   sync.Mutex
	data map[string][]model.Post
}

func NewSimpleCache() *Simple {
	return &Simple{
		data: make(map[string][]model.Post),
	}
}

func (s *Simple) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range subs {

	}
}

func (s *Simple) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {

}
