package erratum

type Resource interface {
	Close() error
	Frob(string)
	Defrob(string)
}

type ResourceOpener func() (Resource, error)

func Use(opener ResourceOpener, input string) (returnError error) {
	r, err := opener()
	if err != nil {
		_, ok := err.(TransientError)
		if ok {
			return Use(opener, input)
		}

		return err
	}

	defer func() error {
		if rec := recover(); rec != nil {
			frobErr, ok := rec.(FrobError)
			if ok {
				r.Defrob(frobErr.defrobTag)
				returnError = frobErr.err
			} else {
				returnError = rec.(error)
			}

			r.Close()
		}
		return nil
	}()

	r.Frob(input)
	r.Close()

	return nil
}
