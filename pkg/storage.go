package pkg

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

// Кэш приходящей статистики
type Storage struct {
	mu     sync.Mutex
	client *redis.Client
	dur    time.Duration
}

func NewStorage(conn string, dur time.Duration) *Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "",
		DB:       0,
	})
	return &Storage{client: client, dur: dur}
}

// получение данных из кэша
func (s *Storage) GetData() ([]IncomingData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	data := s.client.MGet().String()
	var output []IncomingData
	err := json.Unmarshal([]byte(data), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// загрузка данных в кэш
func (s *Storage) UpdateData(data Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, transaction := range data.Data {
		for _, items := range transaction {
			for _, item := range items {
				err := s.client.Set(item.Time, item, s.dur).Err()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
