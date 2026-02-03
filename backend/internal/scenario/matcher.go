package scenario

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func matchCondition(c *gin.Context, cond Condition) bool {
	var actual string

	switch {
	case cond.Field[:6] == "query.":
		actual = c.Query(cond.Field[6:])
	case cond.Field[:7] == "header.":
		actual = c.GetHeader(cond.Field[7:])
	default:
		return false
	}

	switch cond.Op {
	case "eq":
		return actual == cond.Value
	case "neq":
		return actual != cond.Value
	case "gt":
		a, _ := strconv.Atoi(actual)
		b, _ := strconv.Atoi(cond.Value)
		return a > b
	case "lt":
		a, _ := strconv.Atoi(actual)
		b, _ := strconv.Atoi(cond.Value)
		return a < b
	default:
		return false
	}
}
