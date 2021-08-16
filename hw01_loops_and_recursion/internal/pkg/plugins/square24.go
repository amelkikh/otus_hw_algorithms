package plugins

type Square24 struct {
}

func (s Square24) Draw(x, y, size int) string {
	if x == y || x+y == size-1 {
		return "# "
	}
	return ". "
}
