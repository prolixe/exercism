package erratum

const testVersion = 2

func Use(f ResourceOpener, input string) (err error) {

	var r Resource
	defer func() {
		if pErr := recover(); pErr != nil {
			if frobErr, ok := pErr.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
				r.Close()
				err = frobErr.inner
			} else {
				r.Close()
				err = pErr.(error)
			}

		}
	}()

	for {
		r, err = f()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			}
			return err
		}
		break
	}

	r.Frob(input)
	err = r.Close()

	return err
}
