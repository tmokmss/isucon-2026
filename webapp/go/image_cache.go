package isureal

import "sync"

// 画像の本体の読み取りキャッシュ。取得してよいかの判定はキャッシュの前にメモリで行う。
// 上限を超えたら古い順に捨てる
const imageCacheLimit = 100 << 20

type cachedImage struct {
	data      []byte
	mediaType string
}

type imageCacheStore struct {
	mu    sync.Mutex
	items map[string]cachedImage
	order []string
	size  int
}

var imageCache = &imageCacheStore{items: map[string]cachedImage{}}

func (c *imageCacheStore) get(id string) ([]byte, string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.items[id]
	return item.data, item.mediaType, ok
}

func (c *imageCacheStore) put(id string, data []byte, mediaType string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.items[id]; ok {
		return
	}
	c.items[id] = cachedImage{data: data, mediaType: mediaType}
	c.order = append(c.order, id)
	c.size += len(data)
	for c.size > imageCacheLimit && len(c.order) > 0 {
		old := c.order[0]
		c.order = c.order[1:]
		c.size -= len(c.items[old].data)
		delete(c.items, old)
	}
}

func (c *imageCacheStore) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]cachedImage{}
	c.order = nil
	c.size = 0
}
