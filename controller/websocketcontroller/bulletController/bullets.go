package bulletController

import (
	"sync"
)

// Bullets 项目中所有的直播
type Bullets struct {
	m    map[int]*Bullet //直播集合
	lock sync.RWMutex
}

func NewBullets() *Bullets {
	return &Bullets{
		m: make(map[int]*Bullet),
	}
}

// GetBullet 获取直播间弹幕结构体
func (b *Bullets) GetBullet(id int) *Bullet {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.m[id]
}

// AddBullet 项目中新增直播间
func (b *Bullets) AddBullet(id int) *Bullet {
	var bullet Bullet
	b.lock.Lock()
	b.m[id] = &bullet
	b.lock.Unlock()
	return &bullet
}
