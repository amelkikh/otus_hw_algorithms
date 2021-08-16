package plugins

type Square7 struct {
}

func (s Square7) Draw(x, y, size int) string {
	if x > 15 && y > 15 {
		return "# "
	}
	return ". "

}
