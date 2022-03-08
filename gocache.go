package GoCache

import "sync"

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

type Getter interface {
	// 回调函数Get，缓存中不存在查询的数据时，调用Get以将源数据加入缓存
	Get(key string) ([]byte, error)
}


type GetterFunc func(key string) ([]byte, error)

// 函数类型实现某一个接口，称为接口型函数。调用时既能传入函数作为参数，也可以传入实现了该接口的结构体作为参数

// 接口型函数只能应用于接口内部只定义了一个方法的情况
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter parameter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group {
		name: name,
		getter: getter,
		mainCache: cache {cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}