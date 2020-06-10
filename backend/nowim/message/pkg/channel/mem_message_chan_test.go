package channel

import (
	"github.com/stretchr/testify/assert"
	"nowim.message/internal/domain"
	"testing"
	"time"
)

func TestMemMessageChannel(t *testing.T) {
	mmchan := NewMemMessageChannel()

	go func() {
		for i := 0; i < 10000; i += 1 {
			err := mmchan.PushMessage(0, domain.NewMessage(int64(i), int64(i), time.Now().UnixNano(), ""))
			if err != nil {
				t.Errorf("push message err: %+v", err)
			}
		}
		t.Log("producer exit")
	}()

	stop := false
	cnt := 0

	go func() {
		for {
			m, err := mmchan.Consume(0)
			if err != nil {
				t.Errorf("consume err: %+v", err)
			}
			t.Logf("receive message %d", m.SenderID)
			cnt += 1

			if stop {
				return
			}
		}
	}()
	time.Sleep(3*time.Second)
	stop = true
	assert.Equal(t, 10000, cnt)
}
