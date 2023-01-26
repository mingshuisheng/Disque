package errorUtils

import "emperror.dev/errors"

func Combine(errs ...error) error {
	if len(errs) != 0 && errs[0] != nil {
		return errors.Combine(errs...)
	}
	return nil
}
