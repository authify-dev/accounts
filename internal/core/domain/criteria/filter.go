package criteria

// FilterField representa el campo sobre el que se aplica el filtro.
type FilterField string

// FilterValue restringe los tipos permitidos para el valor del filtro.
type FilterValue interface {
	int | string
}

// SQLOperator representa un operador SQL válido.
type Operator string

// Enumeración de operadores SQL válidos.
const (
	OperatorEqual        Operator = "="
	OperatorNotEqual     Operator = "<>"
	OperatorGreaterThan  Operator = ">"
	OperatorGreaterEqual Operator = ">="
	OperatorLessThan     Operator = "<"
	OperatorLessEqual    Operator = "<="
	OperatorLike         Operator = "LIKE"
	OperatorNotLike      Operator = "NOT LIKE"
	OperatorIn           Operator = "IN"
	OperatorNotIn        Operator = "NOT IN"
)

// Filter representa una condición de filtro utilizando genéricos.
type Filter struct {
	Field    FilterField
	Operator Operator
	Value    interface{}
}
