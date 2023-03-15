package conf

type KindPair struct {
	Key     string `json:"key"`
	ValueCN string `json:"valueCN"`
	ValueEN string `json:"valueEN"`
}

var Kinds []KindPair

const (
	KindRoot  = "root"
	KindDev   = "dev"
	KindOps   = "ops"
	KindQa    = "qa"
	KindGuest = "guest"
)

func init() {
	Kinds = []KindPair{
		{KindDev, "研发", "Dev"},
		{KindOps, "运维", "Ops"},
		{KindQa, "质量", "QA"},
		{KindGuest, "游客", "Guest"},
	}
}
