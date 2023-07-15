package common

type AdapterError struct {
	Code int
	Msg  string
}

func (a *AdapterError) Error() string {
	return a.Msg
}
