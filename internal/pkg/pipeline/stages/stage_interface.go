package stages

type Stage interface {
	Stub() error // Master端使用
	Run() error  // Agent端调用
}
