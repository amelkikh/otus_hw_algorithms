package plugins

type Square12 struct {
}

func (s Square12) Draw(x, y, size int) string {
	if x*x+y*y <= (size-4)*(size-4) {
		return "# "
	}
	return ". "

}
