package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Event - структура, описывающая событие
type Event struct {
	EventID     int    `json:"event_id,omitempty"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        Date   `json:"date"`
}

// Date - отдельная структура для переопределения метода UnmarshalJSON
// из пакета json, чтобы правильно декодиравать даты, подаваемой клиентом в запросах
type Date struct {
	time.Time
}

// UnmarshalJSON - переопределённый метода из пакета json для парсинга переданной клиентом даты
func (t *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == `""` {
		*t = Date{time.Now()}
		return nil
	}

	timeStr := strings.ReplaceAll(string(data), `"`, "")
	parsedTime, err := time.Parse("2006-01-02T15:04", timeStr)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02T15:04:00Z", timeStr)
		if err != nil {
			parsedTime, err = time.Parse("2006-01-02", timeStr)
			if err != nil {
				return errors.New("wrong date format, example 2006-01-02T15:04")
			}
		}
	}
	*t = Date{parsedTime}
	return nil
}

// ErrorEvent - структура, с помощью которой выводим оишбку на страницу
type ErrorEvent struct {
	Error string `json:"error"`
}

// Result - структура, с помощью которой выводим успешный результат запроса на страницу
type Result struct {
	Message string  `json:"message"`
	Events  []Event `json:"result"`
}

// Store - хранилище событий
type Store struct {
	mu            *sync.RWMutex
	storageEvents map[int]*Event
	eventNumber   int
}

// NewStore - Фунция-конструктор для store
func NewStore() Store {
	return Store{
		storageEvents: make(map[int]*Event),
		eventNumber:   1,
		mu:            new(sync.RWMutex),
	}
}

// CreateEvent - Метод для создания события, ему присваивается свой уникальный ID
// и далее оно кладётся в хранилище
func (s *Store) CreateEvent(e *Event) {
	s.mu.Lock()
	e.EventID = s.eventNumber
	s.storageEvents[e.EventID] = e
	s.eventNumber++
	s.mu.Unlock()
}

// UpdateEvent - Метод для обновления события
func (s *Store) UpdateEvent(e *Event) error {
	s.mu.Lock()
	if _, ok := s.storageEvents[e.EventID]; !ok {
		return fmt.Errorf("еvent with event id %d was not found", e.EventID)
	}
	s.storageEvents[e.EventID] = e
	s.mu.Unlock()
	return nil
}

// DeleteEvent - Метод для удаления события
func (s *Store) DeleteEvent(id int) (event *Event, err error) {
	s.mu.Lock()
	if _, ok := s.storageEvents[id]; !ok {
		return nil, fmt.Errorf("еvent with event id %d was not found", id)
	}
	event = s.storageEvents[id]
	delete(s.storageEvents, id)
	s.mu.Unlock()
	return event, nil
}

//EventsForDay - Метод для получения событий конкретного пользователя в указанный день
func (s *Store) EventsForDay(userID int, date time.Time) []Event {
	events := make([]Event, 0)
	s.mu.RLock()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		if e.Date.Day() == date.Day() && e.Date.Month() == date.Month() && e.Date.Year() == date.Year() && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}

//EventsForWeek - Метод для получения событий конкретного пользователя в указанную неделю аналогично EventsForDay
func (s *Store) EventsForWeek(userID int, date time.Time) []Event {
	events := make([]Event, 0)
	s.mu.RLock()
	yearOne, weekOne := date.ISOWeek()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		// Получаем номер года и недели события, затем сравниваем их с теми, которые в запросе
		yearTwo, weekTwo := e.Date.ISOWeek()
		if yearOne == yearTwo && weekOne == weekTwo && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}

//EventsForMonth - Метод для получения событий конкретного пользователя в указанный месяц аналогично EventsForDay
func (s *Store) EventsForMonth(userID int, date time.Time) []Event {
	events := make([]Event, 0)
	s.mu.RLock()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		if date.Year() == e.Date.Year() && date.Month() == e.Date.Month() && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}
