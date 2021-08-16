package plugins

type Square22 struct {
}

func (s Square22) Draw(x, y, size int) string {
	if (x+y)%3 == 0 {
		return "# "
	}
	return ". "
}
