package resulterrors

type immutableError string

func (e immutableError) Error() string {
	return string(e)
}

const FailedResultFromNilError = "cannot create failed result from nil error"
const UnsafeGetValueError = "unsafe behavior: error was not checked before trying to get the value"
