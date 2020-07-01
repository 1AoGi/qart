package request

import (
	"github.com/creasty/defaults"
	"math/rand"
	"strconv"
)

type Operation struct {
	Image        string `json:"image" default:"default"`
	Dx           int    `json:"dx" default:"4"`
	Dy           int    `json:"dy" default:"4"`
	Size         int    `json:"size" default:"0"`
	URL          string `json:"url" default:"https://example.com"`
	Version      int    `json:"version" default:"6"` // range in [0,9]
	Mask         int    `json:"mask" default:"2"`
	RandControl  bool   `json:"randcontrol" default:"false"`
	Dither       bool   `json:"dither" default:"false"`
	OnlyDataBits bool   `json:"onlydatabits" default:"false"`
	SaveControl  bool   `json:"savecontrol" default:"false"`
	Seed         string `json:"seed"`
	Scale        int    `json:"scale" default:"4"`
	Rotation     int    `json:"rotate" default:"0"` // range in [0,3]
}

func (op *Operation) SetDefaults() {
	if defaults.CanUpdate(op.Seed) {
		op.Seed = strconv.FormatInt(rand.Int63(), 10)
	}
}

func (op *Operation) GetVersion() int {
	if op.Version < 0 {
		return 0
	}
	if op.Version > 9 {
		return 9
	}
	return op.Version
}

func (op *Operation) GetRotation() int {
	if op.Rotation < 0 {
		return 0
	}
	if op.Rotation > 3 {
		return 3
	}
	return op.Rotation
}

func (op *Operation) GetScale() int {
	if op.Version >= 12 && op.Scale >= 4 {
		return op.Scale / 2
	}
	return op.Scale
}

func (op *Operation) GetSeed() int64 {
	seed, err := strconv.ParseInt(op.Seed, 10, 64)
	if err != nil {
		return int64(rand.Int63())
	}
	return seed
}

func NewOperation() (*Operation, error) {
	operation := &Operation{}
	var err error
	if err = defaults.Set(operation); err != nil {
		return nil, err
	}
	return operation, nil
}
