package lang

type Golang struct {
	TplCommon
}

func (m *Golang) Sdk(env string) string {
	const golang = "golang"
	switch env {
	case Go117:
		return golang + ":1.17.0"
	case Go116:
		return golang + ":1.16.0"
	case Go120:
		return golang + ":1.20.0"
	case Go119:
		return golang + ":1.19.0"
	case Go118:
		return golang + ":1.18.0"
	}

	return golang + ":latest"
}
