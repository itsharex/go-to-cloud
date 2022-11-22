package stages

type Stage interface {
	Run() error // Agent端调用
}
