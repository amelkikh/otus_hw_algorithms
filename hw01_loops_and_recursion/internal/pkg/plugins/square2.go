package plugins

type Square2 struct {
}

func (s Square2) Draw(x, y, _ int) string {
	if x == y {
		return "# "
	}
	return ". "

}
