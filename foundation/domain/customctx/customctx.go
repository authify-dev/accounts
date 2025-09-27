package customctx

import (
	"context"
	"fmt"
	"foundation/utils/cerrs"
	"runtime"
	"sync"
	"time"
)

type WrapError struct {
	Error  cerrs.CustomErrorInterface `json:"error"`
	CallIn string                     `json:"call_in"`
}

type CustomContext struct {
	parent context.Context
	mu     sync.Mutex
	errors []WrapError
}

func NewCustomContext(ctx context.Context) *CustomContext {
	return &CustomContext{parent: ctx}
}

// Métodos de context.Context (igual que antes)
func (c *CustomContext) Deadline() (time.Time, bool)       { return c.parent.Deadline() }
func (c *CustomContext) Done() <-chan struct{}             { return c.parent.Done() }
func (c *CustomContext) Err() error                        { return c.parent.Err() }
func (c *CustomContext) Value(key interface{}) interface{} { return c.parent.Value(key) }
func (c *CustomContext) Errors() []WrapError               { return c.errors }
func (c *CustomContext) FirstError() WrapError             { return c.errors[0] }
func (c *CustomContext) LastError() WrapError              { return c.errors[len(c.errors)-1] }
func (c *CustomContext) Context() context.Context          { return c.parent }

// AddError ahora acepta wrapping de errores
func (c *CustomContext) addError(err cerrs.CustomErrorInterface, call_in string) {
	if err == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.errors = append(c.errors, WrapError{Error: err, CallIn: call_in})
}

// WrapError añade un nuevo error haciendo wrapping del anterior
func (c *CustomContext) NewError(err cerrs.CustomErrorInterface) cerrs.CustomErrorInterface {

	if err == nil {
		return nil
	}

	pc, _, line, ok := runtime.Caller(1) // 1 => calling frame
	caller := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		caller = fmt.Sprintf("%s:%d", fn.Name(), line)
	}

	c.addError(err, caller)
	return err
}
