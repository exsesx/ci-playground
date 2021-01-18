package server

func Init() {
	r := NewRouter()
	err := r.Run(":8000")

	if err != nil {
		panic("The server failed to start up")
	}
}
