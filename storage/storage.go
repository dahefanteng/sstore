package badger

type Engine interface {
	Set(key,value string) error
	Get(key string) (string,error)
}
