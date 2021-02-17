package adapter

type DBHandler interface {
	Find(dest interface{}, conds ...interface{}) (err error)
	Create(value interface{}) (err error)
}
