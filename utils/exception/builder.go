package exception

// When will generate error wrapper for you to do something
func When(err error) *Wrapper {
	return &Wrapper{
		err: err,
	}
}
