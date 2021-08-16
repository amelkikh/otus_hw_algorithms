package plugins

type Square4 struct {
}

func (s Square4) Draw(x, y, size int) string {
	if y < size-x+5 {
		return "# "
	}
	return ". "

}
