package author

type FilterParamsDTO struct {
	Limit     uint
	Offset    uint
	OrderBy   []string
	GroupBy   []string
	Ascending string
}
