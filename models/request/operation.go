package request

type Operation struct {
	Image        string `json:"image"`
	Dx           string `json:"dx"`
	Dy           string `json:"dy"`
	Size         string `json:"size"`
	URL          string `json:"url"`
	Version      string `json:"version"`
	Mask         string `json:"mask"`
	RandControl  string `json:"randcontrol"`
	Dither       bool   `json:"dither"`
	OnlyDataBits bool   `json:"onlydatabits"`
	SaveControl  bool   `json:"savecontrol"`
	Seed         string `json:"seed"`
}
