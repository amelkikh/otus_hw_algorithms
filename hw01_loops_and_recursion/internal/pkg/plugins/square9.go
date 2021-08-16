package plugins

type Square9 struct {
}

func (s Square9) Draw(x, y, size int) string {
	if x-y > 11 || y-x > 11 {
		return "# "
	}
	return ". "

}
