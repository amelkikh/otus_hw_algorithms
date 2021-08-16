package plugins

type Square8 struct {
}

func (s Square8) Draw(x, y, size int) string {
	if x == 0 || y == 0 {
		return "# "
	}
	return ". "

}
