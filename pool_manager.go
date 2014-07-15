package db

var IpToPoolMap = make(map[string]interface{})

type createPoolCallBack func(hostNPort []string) interface{}

type PoolManager struct{}

func (m PoolManager) GetConnection(cb createPoolCallBack, hostNPort string) interface{} {
	pool, ok := IpToPoolMap[hostNPort]
	if !ok {
		hostAndPort := []string{hostNPort}
		pool = cb(hostAndPort)
		IpToPoolMap[hostNPort] = pool
	}
	return pool
}
