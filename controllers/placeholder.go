package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
)

type PlaceHolderController struct {
	beego.Controller
}

type PlaceHolder struct {
	Width  int
	Height int
	Random string
}

var SvgTemplate = `<svg width="{{.Width}}" height="{{.Height}}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {{.Width}} {{.Height}}" preserveAspectRatio="none">
	<defs>
		<style type="text/css">#holder_{{.Random}} text { fill:rgba(255,255,255,.75);font-weight:normal;font-family:Helvetica, monospace;font-size:{{fontsize .}}pt } </style>
	</defs>
	<g id="holder_{{.Random}}">
		<rect width="{{.Width}}" height="{{.Height}}" fill="#777"></rect>
		<g>
			<text x="0" y="{{fontsize .}}">{{.Width}}x{{.Height}}</text>
		</g>
	</g>
</svg>`

func GetTemplate() (*template.Template, error) {
	fm := template.FuncMap{"fontsize": func(ph *PlaceHolder) int {
		if ph.Width/ph.Height > 10 {
			return ph.Height / 10
		}
		return ph.Width / 20
	}}
	tpl, err := template.New("svg").Funcs(fm).Parse(SvgTemplate)
	return tpl, err
}

func (p *PlaceHolderController) Get() {
	width := 0
	height := 0
	var err error
	size := p.Ctx.Input.Param(":size")
	if size != "" {
		seps := strings.Split(size, "x")
		if len(seps) == 2 {
			width, err = strconv.Atoi(seps[0])
			if err != nil {
				width = 0
			}
			height, err = strconv.Atoi(seps[1])
			if err != nil {
				height = 0
			}
		}
	}
	if width == 0 && height == 0 {
		width = 100
		height = 100
	}

	placeHolder := &PlaceHolder{
		Width:  width,
		Height: height,
		Random: strconv.Itoa(int(rand.Int31())),
	}
	tpl, err := GetTemplate()
	if err != nil {
		panic(err)
	}
	var svg bytes.Buffer
	err = tpl.Execute(&svg, placeHolder)
	if err != nil {
		panic(err)
	}
	p.Ctx.Output.ContentType(".svg")
	p.Ctx.Output.Body(svg.Bytes())
}
