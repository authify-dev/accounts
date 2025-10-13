package utils

import (
	"foundation/utils/cerrs"
)

type Result[R any] struct {
	Data R
	Err  cerrs.CustomErrorInterface
}
