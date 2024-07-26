package client

type Pool struct {
	imps []*Imp
}

func NewPool(confPath string) *Pool {
	config := loadConfig(confPath)
	if config == nil {
		return nil
	}
	imp := newImp(config)
	if imp == nil {
		return nil
	}
	return &Pool{
		imps: []*Imp{imp},
	}
}

func (pool *Pool) ProperClient() *Imp {
	//TODO check status
	if len(pool.imps) == 0 {
		return nil
	}
	return pool.imps[0]
}
