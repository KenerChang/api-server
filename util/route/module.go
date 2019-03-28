package route

type Module struct {
	Name string
	Entrypoints []Entrypoint
	InitFunc func()
}