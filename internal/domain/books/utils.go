package books

import (
	"fmt"
	"strings"
)

func getFilterQuery(baseQuery string, param *QueryFilterDTO) (q string, args []interface{}) {
	q = baseQuery
	if param.Title != "" {
		q += fmt.Sprintf(" where title ilike $%d", len(args)+1)
		args = append(args, param.Title)
	}

	if param.AuthorID != nil {
		if !strings.Contains(q, "where") {
			q += " where"
		}
		q += fmt.Sprintf(" author_id=$%d", len(args)+1)
		args = append(args, param.AuthorID)
	}

	if param.OrderBy != "" {
		q += " order by " + param.OrderBy
		if param.Sort != "" {
			q += param.Sort
		}
	}

	if param.Limit != nil {
		q += fmt.Sprintf(" limit $%d", len(args)+1)
		args = append(args, param.Limit)
	}

	if param.Offset != nil {
		q += fmt.Sprintf(" offset $%d", len(args)+1)
		args = append(args, param.Offset)
	}
	return q, args
}
