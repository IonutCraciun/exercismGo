package erratum

// Use func
func Use(o ResourceOpener, input string) (err error) {
	var res Resource

	// open with resourceopener
	// while err is different than nil
	// check if the error is TransientError
	// if it isn't return it
	// if it is, retry to open the resource until is it succesfull
	res, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}
	defer res.Close()

	// defer this function
	// if frob called "panic" that means r different than nil
	// save the type of error that was "panicked" and check it's type
	// for FrobError type, call Defrob and save error
	// for other type of errors, just save it
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	res.Frob(input)
	// return err that is nil when it's reachable
	return err
}
