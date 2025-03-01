package controller

func GetTestHTTPServer(u UseCaseInterface) *Server {
	l := LoggerTest{}
	s := New(u, &l)
	return s
}
