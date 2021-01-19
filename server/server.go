package server

func Init() (err error) {
	r := NewRouter()

	return r.Run(":8000")
}
