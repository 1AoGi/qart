package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"qart/controllers/base"
	"qart/controllers/sessionutils"
	"qart/models/qr"
	"qart/models/request"
	"qart/qrweb/utils"
)

type ShareController struct {
	base.QArtController
}

func (c *ShareController) CreateShare() {
	var err error
	share := &request.Share{}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, share); err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
	config := c.GetSession(sessionutils.SessionKey(share.Image, "config"))
	if config == nil {
		c.Fail(nil, 2, "Image not found")
		return
	}
	image := config.(*qr.Image)
	pngData := image.Code.PNG()
	sha := fmt.Sprintf("%x", sha256.Sum256(pngData))
	if err := utils.Write(utils.GetQrsavePath(sha), pngData); err != nil {
		panic(err)
	}
	c.Success(struct {
		Id string `json:"id"`
	}{
		sha,
	}, 0)
}

func (c *ShareController) Get() {
	sha := c.Ctx.Input.Param(":sha")
	data, _, err := utils.Read(utils.GetQrsavePath(sha))
	if err != nil {
		c.Redirect("/image/placeholder/400x400/QR%20Code%20Not%20Found", 302)
	}
	c.Ctx.Output.ContentType(".png")
	err = c.Ctx.Output.Body(data)
}
