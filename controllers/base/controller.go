package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/tautcony/qart/models/response"
	"log"
	"strings"
	"time"
)

var (
	AppVer string
)

var langTypes []*langType // Languages are supported.

// langType represents a language type.
type langType struct {
	Lang, Name string
}

type QArtController struct {
	beego.Controller
	i18n.Locale
}

func (c *QArtController) JSON(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *QArtController) Success(data interface{}, code int) {
	r := &response.BaseResponse{
		Success: true,
		Code:    code,
		Data:    data,
	}
	c.JSON(r)
}

func (c *QArtController) Fail(data interface{}, code int, message string) {
	r := &response.BaseResponse{
		Code:    code,
		Data:    data,
		Message: message,
	}
	c.JSON(r)
}

// Prepare implemented Prepare method for baseRouter.
func (c *QArtController) Prepare() {
	// Setting properties.
	c.Data["AppVer"] = AppVer

	c.Data["PageStartTime"] = time.Now()

	// Redirect to make URL clean.
	if c.setLangVer() {
		i := strings.Index(c.Ctx.Request.RequestURI, "?")
		c.Redirect(c.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

// setLangVer sets site language version.
func (c *QArtController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := c.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		c.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	c.Lang = lang
	c.Data["Lang"] = curLang.Lang
	c.Data["CurLang"] = curLang.Name
	c.Data["RestLangs"] = restLangs

	return isNeedRedir
}

func initLocales() {
	// Initialized language type list.
	var availableLangs map[string]string
	langConfig := beego.AppConfig.String("lang::available_lang")
	err := json.Unmarshal([]byte(langConfig), &availableLangs)
	if err != nil {
		log.Fatalln("Language config invalid", langConfig)
		return
	}

	langTypes = make([]*langType, 0, len(availableLangs))
	for lang, name := range availableLangs {
		langTypes = append(langTypes, &langType{
			Lang: lang,
			Name: name,
		})
	}

	for _, langType := range langTypes {
		log.Println("Loading language: " + langType.Lang)
		if err := i18n.SetMessage(langType.Lang, "conf/"+"locale_"+langType.Lang+".ini"); err != nil {
			log.Println("Fail to set message file: " + err.Error())
			return
		}
	}
}

func init() {
	initLocales()
}
