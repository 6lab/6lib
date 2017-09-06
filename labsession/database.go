package labsession

type DATABASE struct {
	// List of Methods to be implemented in the db struct (couchbase, ...)
	accesser interface {
		readAll(*SESSIONMANAGER) error
		readOne(idUser string) (*SESSION, error)
		add(*SESSION) error
		deleteAll() error
		deleteOne(idUser) error
	}
}
