package calendar

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNoEvent = errors.New("no event")
	ErrNoUser  = errors.New("no user")
)

type Event struct {
	UserID int64
	Id     int64
	Name   string
	Desc   string
	Date   time.Time
}

type Calendar struct {
	list map[int64][]Event
	sync.RWMutex
}

func New() *Calendar {
	return &Calendar{
		list: make(map[int64][]Event),
	}
}

func (c *Calendar) CreateEvent(e *Event) {
	c.Lock()
	defer c.Unlock()

	c.list[e.UserID] = append(c.list[e.UserID], *e)

	(*e).Id = int64(len(c.list[e.UserID]) - 1)
}

func (c *Calendar) UpdateEvent(e *Event) error {
	c.Lock()
	defer c.Unlock()

	usrEvents, ok := c.list[e.UserID]
	if !ok {
		return ErrNoUser
	}

	if e.Id < 0 || e.Id >= int64(len(usrEvents)) {
		return ErrNoEvent
	}

	usrEvents[e.Id] = *e

	return nil
}

func (c *Calendar) RemoveEvent(userid int64, id int64) error {
	c.Lock()
	defer c.Unlock()

	usrEvents, ok := c.list[userid]
	if !ok {
		return ErrNoUser
	}

	if id < 0 || id >= int64(len(usrEvents)) {
		return ErrNoEvent
	}

	c.list[userid] = append(usrEvents[:id], usrEvents[id+1:]...)

	return nil
}

func (c *Calendar) GetEvents(userid int64, from time.Time, to time.Time) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	usrEvents, ok := c.list[userid]
	if !ok {
		return nil, ErrNoUser
	}

	var ret []Event

	for _, v := range usrEvents {
		if v.Date.After(from) && v.Date.Before(to) {
			ret = append(ret, v)
		}
	}

	return ret, nil
}
