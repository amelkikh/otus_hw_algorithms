package plugins

type Square18 struct {
}

func (s Square18) Draw(x, y, size int) string {
	if (x < 2 || y < 2) && x+y > 0 {
		return "# "
	}
	return ". "
}
