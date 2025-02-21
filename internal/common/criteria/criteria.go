package criteria

type Criteria struct {
	Filters []Filter
}

type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}
