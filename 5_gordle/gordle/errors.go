package gordle

type corpusError string

// the implementation fo the error interface by corpusError
func (e corpusError) Error() string {
	return string(e)
}
