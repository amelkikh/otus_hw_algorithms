package plugins

type Square3 struct {
}

func (s Square3) Draw(x, y, size int) string {
	if x+y == size-1 {
		return "# "
	}
	return ". "

}
