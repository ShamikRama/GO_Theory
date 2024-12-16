package cache

import (
	"context"
	"errors"
	"sync"

	"github.com/glebnaz/cache-webinar/internal/model"
	lru "github.com/hashicorp/golang-lru"
)

type LRU struct {
	m    sync.RWMutex
	data *lru.Cache
}

func NewLRU() (*LRU, error) {
	data, err := lru.New(100)
	if err != nil {
		return nil, err
	}
	return &LRU{data: data}, nil
}

func (l *LRU) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	l.m.Lock()
	defer l.m.Unlock()

	for i := range subs {
		toWrite := []model.Post{}
		val, ok := l.data.Get(subs[i])
		if !ok {
			toWrite = []model.Post{post}
		} else {
			toWrite, ok := val.([]model.Post)
			if !ok {
				return errors.New("not a type")
			}
			toWrite = append(toWrite, post)
		}
		l.data.Add(subs[i], toWrite)
	}

	return nil
}

func (l *LRU) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {
	l.m.Lock()
	defer l.m.Unlock()

	val, ok := l.data.Get(id)
	if !ok {
		return []model.Post{}, errors.New("nothing")
	}

	feed, ok := val.([]model.Post)
	if !ok {
		return feed, errors.New("not a type")
	}

	return feed, nil

}
