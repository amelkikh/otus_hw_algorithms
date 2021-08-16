package plugins

type Square11 struct {
}

func (s Square11) Draw(x, y, size int) string {
	if x == size-2 || y == size-2 || x == 1 || y == 1 {
		return "# "
	}
	return ". "

}
