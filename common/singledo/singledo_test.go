package singledo

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	single := NewSingle(time.Millisecond * 30)
	foo := 0
	shardCount := 0
	call := func() (interface{}, error) {
		foo++
		return nil, nil
	}

	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			_, _, shard := single.Do(call)
			if shard {
				shardCount++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, 1, foo)
	assert.Equal(t, 9, shardCount)
}

func TestTimer(t *testing.T) {
	single := NewSingle(time.Millisecond * 30)
	foo := 0
	call := func() (interface{}, error) {
		foo++
		return nil, nil
	}

	single.Do(call)
	time.Sleep(10 * time.Millisecond)
	_, _, shard := single.Do(call)

	assert.Equal(t, 1, foo)
	assert.True(t, shard)
}
