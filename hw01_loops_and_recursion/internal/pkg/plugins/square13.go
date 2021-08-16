package plugins

type Square13 struct {
}

func (s Square13) Draw(x, y, size int) string {
	if x+y >= size-5 && x+y <= size+3 {
		return "# "
	}
	return ". "

}
