package time

import (
	"context"
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	tm := time.NewTicker(time.Second * 10)
	defer tm.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			t.Log("超时或者被取消")
			goto end
		case now := <-tm.C:
			t.Log(now.Unix())
		}
	}
end:
	t.Log("退出循环")
}
