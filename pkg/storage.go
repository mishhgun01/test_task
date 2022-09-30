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
	// ключ в формате год/месяц/денб/час
	key := strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + "/" + strconv.Itoa(time.Now().Hour())
	data, err := s.client.Get(key).Result()
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
	// ключ в формате год/месяц/денб/час
	key := strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + "/" + strconv.Itoa(time.Now().Hour())
	s.client.Set(key, jsonData, s.dur)
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
