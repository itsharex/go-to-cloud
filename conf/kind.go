package conf

var Kinds []string

const (
	KindRoot  = "root"
	KindDev   = "dev"
	KindOps   = "ops"
	KindQa    = "qa"
	KindGuest = "guest"
)

func init() {
	Kinds = []string{
		// KindRoot, Only ONE
		KindDev,
		KindOps,
		KindQa,
		KindGuest,
	}
}
