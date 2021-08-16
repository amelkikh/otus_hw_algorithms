package plugins

type Square1 struct {
}

func (s *Square1) Draw(x, y, _ int) string {
	if x > y {
		return "# "
	}
	return ". "

}
