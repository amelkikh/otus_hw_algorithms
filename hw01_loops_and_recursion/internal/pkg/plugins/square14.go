package plugins

type Square14 struct {
}

func (s Square14) Draw(x, y, size int) string {
	offset := -3
	c := 1
	if (x-size+offset)*(x-size+offset)+(y-size+offset)*(y-size+offset) >= (size-c)*(size-c) {
		return "# "
	}
	return ". "

}
