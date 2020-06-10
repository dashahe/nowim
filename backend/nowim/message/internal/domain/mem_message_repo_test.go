package domain

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func testMessageRepo(t *testing.T, repo MessageRepo) {

	var senderID, receiverID int64
	senderID, receiverID = 1, 2

	for i := 0; i < 100*1000; i += 1 {
		if i % 2 == 0 {
			senderID, receiverID = 3, 4
		} else {
			senderID, receiverID = 1, 2
		}
		err := repo.SaveMessage(NewMessage(senderID, receiverID, time.Now().UnixNano(), strconv.FormatInt(int64(i), 10)))
		if err != nil {
			t.Errorf("save message err: %s", err)
		}
	}

	messages, _ := repo.QueryMessage(1, 2, 100)
	for i, m := range messages {
		expect := int64(1000*1000 - 1 - 2*i)
		assert.Equal(t, strconv.FormatInt(expect, 10), m.Content)
		if i > 0 {
			assert.True(t, m.MessageID < messages[i-1].MessageID)
		}
	}
}

func TestMemMessageRepo(t *testing.T) {
	testMessageRepo(t, NewMemMessageRepo())
}