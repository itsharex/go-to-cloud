package conf

type KindPair struct {
	Key     Kind   `json:"key"`
	ValueCN string `json:"valueCN"`
	ValueEN string `json:"valueEN"`
}

var Kinds []KindPair

type Kind string

const (
	Root  Kind = "root"
	Dev   Kind = "dev"
	Ops   Kind = "ops"
	Qa    Kind = "qa"
	Guest Kind = "guest"
)

func init() {
	Kinds = []KindPair{
		{Dev, "研发", "Dev"},
		{Ops, "运维", "Ops"},
		{Qa, "质量", "QA"},
		{Guest, "游客", "Guest"},
	}
}
