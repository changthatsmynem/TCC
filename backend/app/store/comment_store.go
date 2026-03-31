package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"tcc/backend/app/model"
	"time"
)

const dateLayout = "2 January 2006 15:04"

type CommentStore struct {
	mu       sync.Mutex
	path     string
	comments []model.Comment
	nextID   int64
}

func NewCommentStore(path string) (*CommentStore, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	store := &CommentStore{
		path:     path,
		comments: []model.Comment{},
		nextID:   1,
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := store.persist(); err != nil {
				return nil, err
			}
			return store, nil
		}

		return nil, err
	}

	if len(data) == 0 {
		return store, nil
	}

	if err := json.Unmarshal(data, &store.comments); err != nil {
		return nil, err
	}

	if store.comments == nil {
		store.comments = []model.Comment{}
	}

	for _, item := range store.comments {
		if item.ID >= store.nextID {
			store.nextID = item.ID + 1
		}
	}

	return store, nil
}

func (s *CommentStore) List() []model.Comment {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]model.Comment{}, s.comments...)
}

func (s *CommentStore) Add(message string) (model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry := model.Comment{
		ID:        s.nextID,
		Author:    "Blend 285",
		Avatar:    "B",
		Message:   strings.TrimSpace(message),
		CreatedAt: time.Now().Format(dateLayout),
	}

	s.comments = append(s.comments, entry)
	s.nextID++

	if err := s.persist(); err != nil {
		return model.Comment{}, err
	}

	return entry, nil
}

func (s *CommentStore) persist() error {
	data, err := json.MarshalIndent(s.comments, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0o644)
}
