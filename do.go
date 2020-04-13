package catch

// Do calls the given function and recovers an Error panic. If an Error panic is
// recovered, its Cause is returned.
//
// A recover of something that isn't an Error or a *Error will cause a new panic.
func Do(doer func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case *Error:
				err = r.Cause
			case Error:
				err = r.Cause
			default:
				panic(r)
			}
		}
	}()
	doer()
	return
}
