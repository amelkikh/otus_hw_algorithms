package plugins

type Square6 struct {
}

func (s Square6) Draw(x, y, size int) string {
	if x < 10 || y < 10 {
		return "# "
	}
	return ". "

}
