package database

type DatabaseError struct {
	err error
}

func newDatabaseError(err error) error {
	return &DatabaseError{
		err: err,
	}
}

func (e *DatabaseError) Error() string {
	return e.err.Error()
}
