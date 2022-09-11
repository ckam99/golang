package sqlx

type whereQuery struct {
	Column   string
	Value    interface{}
	Operator string
	Type     string
}

func addWhereQuery(whereQueries *[]whereQuery, column, oper string, value interface{}, clause string) {
	*whereQueries = append(*whereQueries, whereQuery{
		Column:   column,
		Value:    value,
		Operator: oper,
		Type:     clause,
	})
}
