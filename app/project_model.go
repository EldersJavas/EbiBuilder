package app



const (
	Debug = iota
	Release
)

type Project struct {
	EbiVersion string
	BuildMode  uint
	OutPut     string
	Config string
	Path string
}
