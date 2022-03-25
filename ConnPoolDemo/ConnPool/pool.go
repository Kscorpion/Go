package ConnPool

import (
	DB "ConnPoolDemo/DbConn"
	"sync"
)

type Pool struct {
	mu      sync.Mutex
	minConn int             // 最小连接数
	maxConn int             // 最大连接数
	numConn int             // 池已申请的连接数
	conns   chan *DB.DBConn //当前池中空闲连接实例
	close   bool
}

// NewPool 初始化池实例
func NewPool(min, max int) *Pool {
	p := &Pool{
		minConn: min,
		maxConn: max,
		numConn: min,
		conns:   make(chan *DB.DBConn, max),
		close:   false,
	}
	for i := 0; i < min; i++ {
		p.conns <- DB.NewDBConn()
	}
	return p
}

// Get 从池中取出连接
func (p *Pool) Get() *DB.DBConn {
	if p.close {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.numConn >= p.maxConn || len(p.conns) > 0 { // 保证了池申请连接数量不超过最大连接数
		d := <-p.conns // 若池中没有可取的连接，则等待其他请求返回连接至池中再取
		return d
	}
	p.numConn++
	return DB.NewDBConn() //申请新的连接
}

// Put 将连接返回池中
func (p *Pool) Put(d *DB.DBConn) {
	if p.close {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.conns <- d
}

// Close 关闭池
func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for d := range p.conns {
		d.Close()
	}
	p.close = true
}
