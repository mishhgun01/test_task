package pkg

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

// Кэш вычисляемой статистики
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
func (s *Storage) GetData() (GasStatistics, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := s.client.Get(time.Now().Month().String() + strconv.Itoa(time.Now().Hour())).Result()
	if err != nil {
		return GasStatistics{}, err
	}
	var output GasStatistics
	err = json.Unmarshal([]byte(data), &output)
	if err != nil {
		return GasStatistics{}, err
	}
	return output, nil
}

// загрузка данных в кэш
func (s *Storage) UpdateData(data GasStatistics) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	jsonData, err := json.Marshal(data)
	s.client.Set(time.Now().Month().String()+strconv.Itoa(time.Now().Hour()), jsonData, s.dur)
	return err
}

func (s *Storage) ClearAll() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.FlushAll().Err()
}

func (s *Storage) Close() error {
	return s.client.Close()
}
