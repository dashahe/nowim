package domain

import (
	"testing"
	"time"
)

func TestMongoMessageRepo(t *testing.T) {
	testMessageRepo(t, NewMongoMessageRepo())
}

func TestMongoCollection(t *testing.T) {
	repo := NewMongoMessageRepo()
	err := repo.SaveMessage(NewMessage(1, 2, time.Now().UnixNano(), "content"))
	if err != nil {
		t.Errorf("err: %+v", err)
	}
}