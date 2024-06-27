package domain

type RedisRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, value interface{}) error
	Del(key string) error
}
