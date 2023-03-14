package middlewares

type policy struct {
	Kind     string // 角色
	Resource string
	Act      string
}

var rules []policy

func init() {
	rules = make([]policy, 0, 100)
	rules = append(rules, policy{"1", "2", "3"})
}
