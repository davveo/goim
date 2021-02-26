package cmap

import "sync"

const (
	// 默认32分片
	ShardCount = 32
)

type ConcurrentMap []*concurrentMapShared

type concurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

func NewConcurrentMap() ConcurrentMap {
	m := make(ConcurrentMap, ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = &concurrentMapShared{
			items: make(map[string]interface{}),
		}
	}
	return m
}

func (m ConcurrentMap) getShard(key string) *concurrentMapShared {
	return m[uint(fnv32(key))%uint(ShardCount)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = value
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.items[key]
	return val, ok
}

func (m ConcurrentMap) Remove(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.items, key)
}

func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < ShardCount; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

func (m ConcurrentMap) Keys() []string {
	count := m.Count()
	ch := make(chan string, count)

	// 每一个分片启动一个协程 遍历key
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(ShardCount)
		for _, shard := range m {
			go func(shard *concurrentMapShared) {
				defer wg.Done()
				shard.RLock()
				// 每个分片中的key遍历后都写入统计用的channel
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	keys := make([]string, count)
	// 统计各个协程并发读取Map分片的key
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
