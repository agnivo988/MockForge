package scenario

import "github.com/gin-gonic/gin"

func MatchRule(c *gin.Context, rules []Rule) *Rule {
	var selected *Rule

	for _, rule := range rules {
		matched := true
		for _, cond := range rule.Conditions {
			if !matchCondition(c, cond) {
				matched = false
				break
			}
		}

		if matched {
			if selected == nil || rule.Priority > selected.Priority {
				r := rule
				selected = &r
			}
		}
	}
	return selected
}
