package plugins

type Square16 struct {
}

func (s Square16) Draw(x, y, size int) string {
	// top left -> top right -> bottom right -> bottom left
	if x+y > size/2+2 && y-x < size/2-2 && x+y < size+size/2-3 && y-x > -size/2+2 {
		return "# "
	}
	return ". "

}
