package plugins

type Square5 struct {
}

func (s Square5) Draw(x, y, size int) string {
	if y == x/2 {
		return "# "
	}
	return ". "

}
