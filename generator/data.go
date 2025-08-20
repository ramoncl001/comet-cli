package generator

type ProjectData struct {
	FileName    string
	ProjectName string
	Module      string
}

type ControllerData struct {
	FileName string
	Package  string
	Name     string
}

type ServiceData struct {
	FileName  string
	Package   string
	Name      string
	LowerName string
}

type MiddlewareData struct {
	FileName string
	Package  string
	Name     string
}
