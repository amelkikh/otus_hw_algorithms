package plugins

type Square25 struct {
}

func (s Square25) Draw(x, y, size int) string {
	if x%6 == 0 || y%6 == 0 {
		return "# "
	}
	return ". "
}
