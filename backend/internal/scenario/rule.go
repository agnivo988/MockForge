package scenario

type Condition struct {
	Field string `json:"field"` // query.userId / header.X-FAIL / body.age
	Op    string `json:"op"`    // eq, neq, gt, lt
	Value string `json:"value"`
}

type Rule struct {
	Name       string      `json:"name"`
	Conditions []Condition `json:"conditions"`
	StatusCode int         `json:"status"`
	Response   any         `json:"response"`
	Priority   int         `json:"priority"`
}
