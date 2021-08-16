package plugins

type Square23 struct {
}

func (s Square23) Draw(x, y, size int) string {
	if x%2 == 0 && y%3 == 0 {
		return "# "
	}
	return ". "
}
