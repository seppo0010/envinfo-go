package envinfo

type System struct {
	OS     string
	CPU    string
	Memory string
	Shell  string
}

func GetSystem() *System {
	return &System{
		OS: getOS(),
	}
}
