package errors

import "fmt"

type dbError struct {
	error error
}

type CannotInsertRecordError dbError

func NewCannotInsertRecordError(err error) *CannotInsertRecordError {
	return &CannotInsertRecordError{error: fmt.Errorf("cannot insert the record, for more details: %v", err.Error())}
}
func (cnt CannotInsertRecordError) Error() string {
	return cnt.error.Error()
}

type CannotUpdateRecordError dbError

func NewCannotUpdateRecordError(err error) *CannotUpdateRecordError {
	return &CannotUpdateRecordError{error: fmt.Errorf("cannot update the record, for more details: %v", err.Error())}
}
func (cnt CannotUpdateRecordError) Error() string {
	return cnt.error.Error()
}

type CannotDeleteRecordError dbError

func NewCannotDeleteRecordError(err error) *CannotDeleteRecordError {
	return &CannotDeleteRecordError{error: fmt.Errorf("cannot delete the record, for more details: %v", err.Error())}
}
func (cnt CannotDeleteRecordError) Error() string {
	return cnt.error.Error()
}

type NotFoundRecordError dbError

func NewNotFoundRecordError(err error) *NotFoundRecordError {
	return &NotFoundRecordError{error: fmt.Errorf("cannot find the record, for more details: %v", err.Error())}
}
func (cnt NotFoundRecordError) Error() string {
	return cnt.error.Error()
}
