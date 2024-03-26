package testing

import (
	"sync"
	"testing"
	"time"
)

/*
Mistake 87: Not dealing with the time API efficiently

Devemos ter cuidado ao testar código que faz uso da API "time".
*/
type Event struct {
	Timestamp time.Time
	Data      string
}

type now func() time.Time

type Cache struct {
	mu     sync.RWMutex
	events []Event
	now    now
}

func (c *Cache) Add(events []Event) {}
func (c *Cache) TrimOlderThan(since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	t := time.Now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

func NewCache() *Cache {
	return &Cache{
		events: make([]Event, 0),
		now:    time.Now,
	}
}

func TestCache_TrimOlderThan(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.04Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.05Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.06Z")},
	}
	cache := &Cache{now: func() time.Time {
		return parseTime(t, "2020-01-01T12:00:00.06Z")
	}}
	cache.Add(events)
	cache.TrimOlderThan(15 * time.Millisecond)
	// ...
}

func parseTime(t *testing.T, timestamp string) time.Time {
	return time.Time{}
}
