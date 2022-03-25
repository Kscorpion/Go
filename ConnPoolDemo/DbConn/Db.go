package DbConn

type DBConn struct {
	idleTime int // 标记该数据库连接空闲时间
}

// NewDBConn 新建数据库连接
func NewDBConn() *DBConn {
	return &DBConn{idleTime: 0}
}

// Close 关闭数据库连接
func (d *DBConn) Close() {}
