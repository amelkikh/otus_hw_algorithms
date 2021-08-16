package plugins

type Square21 struct {
}

func (s Square21) Draw(x, y, size int) string {
	if x%(y+1) == 0 {
		return "# "
	}
	return ". "
}
