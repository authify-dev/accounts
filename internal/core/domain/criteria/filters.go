package criteria

type Filters struct {
	filters []Filter
}

func NewFilters(filters []Filter) *Filters {
	return &Filters{
		filters: filters,
	}
}

func (f Filters) Get() []Filter {
	return f.filters
}
