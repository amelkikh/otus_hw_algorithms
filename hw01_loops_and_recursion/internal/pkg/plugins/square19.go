package plugins

type Square19 struct {
}

func (s Square19) Draw(x, y, size int) string {
	if x == 0 || y == 0 || x == size-1 || y == size-1 {
		return "# "
	}
	return ". "
}
