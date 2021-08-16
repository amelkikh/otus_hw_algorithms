package plugins

type Square15 struct {
}

func (s Square15) Draw(x, y, size int) string {
	if (x-y > 11 && x-y < size-4) || (y-x > 11 && y-x < size-4) {
		return "# "
	}
	return ". "

}
