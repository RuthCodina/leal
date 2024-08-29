package api

func Start(port string) {
	h := Handlers{}
	r := routes(&h)
	server := newServer(port, r)
	server.Start()
}
