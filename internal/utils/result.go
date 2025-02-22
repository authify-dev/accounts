package utils

// --------------------------------
// UTILS
// --------------------------------
// Utils
//--------------------------------

type Result[R any] struct {
	Data R
	Err  error
}
