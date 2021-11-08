package validation

import (
	"context"
	"skeleton/lib/helper"
)

func (u *driverValidation) Delete(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	err := u.driverRepo.Find(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
