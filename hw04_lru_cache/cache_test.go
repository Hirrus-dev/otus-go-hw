package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok, _ := c.Get("aaa")
		require.False(t, ok)

		_, ok, _ = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache, _ := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache, _ = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok, _ := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok, _ = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache, _ = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok, _ = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok, _ = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)
		c.Set("ddd", 400)
		_, ok, _ := c.Get("aaa")
		require.False(t, ok)

		_, ok, _ = c.Get("bbb")
		require.True(t, ok)

		c.Set("eee", 500)
		c.Set("fff", 600)
		val, ok, _ := c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		c.Clear()
		_, ok, _ = c.Get("bbb")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
