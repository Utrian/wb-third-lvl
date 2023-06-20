package calendar

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

type Event struct {
	Id          uint64    `json:"id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (e *Event) Value() ([]byte, error) {
	return json.Marshal(e)
}

type Calendar struct {
	events map[uint64][]Event
}

func NewCalendar() *Calendar {
	events := make(map[uint64][]Event)
	return &Calendar{events: events}
}

func (c *Calendar) NewEvent(query *url.Values) (*Event, error) {
	uid, _, date, title, dscr, err := parseQuery(query)
	if err != nil {
		return &Event{}, err
	}

	if !c.isExist(uid) {
		return &Event{}, errors.New("unknown user")
	}

	lastEvn := len(c.events[uid]) - 1
	id := c.events[uid][lastEvn].Id + 1

	evn := Event{Id: id, Date: date, Title: title, Description: dscr}

	c.events[uid] = append(c.events[uid], evn)

	return &evn, nil
}

func (c *Calendar) findEvn(uid, eid uint64) (*Event, error) {
	evns := c.events[uid]

	for _, evn := range evns {
		if evn.Id == eid {
			return &evn, nil
		}
	}

	return nil, errors.New("unknown event")
}

func (c *Calendar) Update(query *url.Values) (*Event, error) {
	uid, evnId, date, title, dscr, err := parseQuery(query)
	if err != nil {
		return &Event{}, err
	}

	eid, err := strconv.ParseUint(evnId, 10, 64)
	if err != nil {
		return &Event{}, err
	}

	if !c.isExist(uid) {
		return &Event{}, errors.New("unknown user")
	}

	evn, err := c.findEvn(uid, eid)
	if err != nil {
		return &Event{}, err
	}

	if date != (time.Time{}) {
		evn.Date = date
	}

	if title != "" {
		evn.Title = title
	}

	if dscr != "" {
		evn.Description = dscr
	}

	return evn, nil
}

func (c *Calendar) Delete(query *url.Values) (*Event, error) {
	uid, evnId, _, _, _, err := parseQuery(query)
	if err != nil {
		return &Event{}, err
	}

	if !c.isExist(uid) {
		return &Event{}, errors.New("unknown user")
	}

	eid, err := strconv.ParseUint(evnId, 10, 64)
	if err != nil {
		return &Event{}, err
	}

	for i, v := range c.events[uid] {
		if v.Id == eid {
			removedEvent := c.events[uid][i]

			lastEID := len(c.events[uid]) - 1

			c.events[uid][i] = c.events[uid][lastEID]
			c.events[uid][lastEID] = Event{}
			c.events[uid] = c.events[uid][:lastEID]

			return &removedEvent, nil
		}
	}

	return &Event{}, errors.New("unknown event")
}

func (c *Calendar) isExist(uid uint64) bool {
	if _, ok := c.events[uid]; !ok {
		return false
	}
	return true
}

func parseQuery(query *url.Values) (uint64, string, time.Time, string, string, error) {
	var (
		uid   uint64
		eid   string
		date  time.Time
		title string
		dscr  string
	)
	uid, err := strconv.ParseUint(query.Get("user_id"), 10, 64)
	if err != nil {
		return uid, eid, date, title, dscr, errors.New("incorrect uid")
	}

	eid = query.Get("event_id")

	dateLayout := "2006-01-02T15:04:05-07:00"
	date, err = time.Parse(dateLayout, query.Get("date"))
	if err != nil {
		return uid, eid, date, title, dscr, errors.New("incorrect date")
	}

	title = query.Get("title")
	dscr = query.Get("description")

	return uid, eid, date, title, dscr, nil
}

func (c *Calendar) DayFilter(query *url.Values) ([]Event, error) {
	uid, _, date, _, _, err := parseQuery(query)
	if err != nil {
		return []Event{}, err
	}

	y := date.Year()
	d := date.YearDay()
	rslt := []Event{}

	evns := c.events[uid]
	for _, v := range evns {
		if v.Date.Year() == y && v.Date.YearDay() == d {
			rslt = append(rslt, v)
		}
	}

	return rslt, nil
}

func (c *Calendar) WeekFilter(query *url.Values) ([]Event, error) {
	uid, _, date, _, _, err := parseQuery(query)
	if err != nil {
		return []Event{}, err
	}

	y := date.Year()
	wd := date.YearDay() - int(date.Weekday())
	rslt := []Event{}

	evns := c.events[uid]
	for _, v := range evns {
		dateWd := date.YearDay() - int(date.Weekday())
		if v.Date.Year() == y && dateWd == wd {
			rslt = append(rslt, v)
		}
	}

	return rslt, nil
}

func (c *Calendar) MonthFilter(query *url.Values) ([]Event, error) {
	uid, _, date, _, _, err := parseQuery(query)
	if err != nil {
		return []Event{}, err
	}

	y := date.Year()
	m := date.Month()
	rslt := []Event{}

	evns := c.events[uid]
	for _, v := range evns {
		if v.Date.Year() == y && v.Date.Month() == m {
			rslt = append(rslt, v)
		}
	}

	return rslt, nil
}
