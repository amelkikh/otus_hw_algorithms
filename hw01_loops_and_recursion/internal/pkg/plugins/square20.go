package plugins

type Square20 struct {
}

func (s Square20) Draw(x, y, size int) string {
	if (x+y)%2 == 0 {
		return "# "
	}
	return ". "
}
