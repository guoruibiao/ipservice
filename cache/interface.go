package cache

import "github.com/guoruibiao/ipservice/constants"

type CacheInterface interface {
	Cache(key string, e constants.Entry) (bool, error)
	Get(key string)(constants.Entry, error)
}
