package errs

type Handler func(err error) Error

func Handle(err error, handlers ...Handler) error {
	switch err.(type) {
	case nil, Error:
		return err

	default:
		for _, handle := range handlers {
			if handle != nil {
				if customErr := handle(err); customErr != nil {
					return handle(err).CausedBy(err)
				}
			}
		}
		return err
	}
}
