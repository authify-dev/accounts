package saga

import (
	"accounts/internal/utils"
	"context"
	"reflect"
)

type SAGA_Step[P any] interface {
	Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[P]
	Rollback(ctx context.Context) error
	Produce() string
}

type SAGA_Controller struct {
	Steps    []SAGA_Step[any]
	Payloads map[string]utils.Result[any]
	PrevSaga *SAGA_Controller
}

func (c *SAGA_Controller) Executed(ctx context.Context) map[string]utils.Result[any] {
	allPayloads := make(map[string]utils.Result[any])
	var lastPayload utils.Result[any]
	for _, step := range c.Steps {

		result := step.Call(ctx, lastPayload, allPayloads)

		// Almacenar el resultado en allPayloads
		lastPayload = result

		name_step := step.Produce()
		if name_step == "" {
			t := reflect.TypeOf(step)
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			name_step = t.Name()
		}

		allPayloads[name_step] = result

		if result.Err != nil {
			c.Rollback(ctx)
			if c.PrevSaga != nil {
				c.PrevSaga.Rollback(ctx)
			}
			break
		}
	}
	c.Payloads = allPayloads
	return c.Payloads
}

func (c SAGA_Controller) Rollback(ctx context.Context) error {
	for i := len(c.Steps) - 1; i >= 0; i-- {
		err := c.Steps[i].Rollback(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c SAGA_Controller) Ok() bool {
	for _, p := range c.Payloads {
		if p.Err != nil {
			return false
		}
	}
	return true
}

func (c SAGA_Controller) Errors() []string {
	var errors []string
	for _, p := range c.Payloads {
		if p.Err != nil {
			errors = append(errors, p.Err.Error())
		}
	}
	return errors
}
