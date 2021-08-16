package plugins

type Square10 struct {
}

func (s Square10) Draw(x, y, size int) string {
	if x > y && y >= x/2 {
		return "# "
	}
	return ". "

}
