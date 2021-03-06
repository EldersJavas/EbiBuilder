package app

import "encoding/json"

const (
	Debug   = "Debug"
	Release = "Release"
	All     = "All"
)

func UnmarshalProject(data []byte) (Project, error) {
	var r Project
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Project) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Project struct {
	IsGomod     bool   `json:"IsGomod,omitempty"`
	IsEbitenv1  bool   `json:"IsEbitenv1,omitempty"`
	EbiVersion  string `json:"EbiVersion,omitempty"`
	BuildMode   string `json:"BuildMode,omitempty"`
	OutputName  string `json:"OutputName,omitempty"`
	Config      string `json:"Config,omitempty"`
	Path        string `json:"Path,omitempty"`
	GameName    string `json:"GameName,omitempty"`
	GameVersion string `json:"GameVersion,omitempty"`
	Iconfile    string `json:"Iconfile,omitempty"`
}
