package router

func hash(verb string, path string) string {
	return verb + " " + path
}
