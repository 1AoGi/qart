package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tautcony/qart/controllers/base"
	"github.com/tautcony/qart/controllers/sessionutils"
	"github.com/tautcony/qart/internal/qr"
	"github.com/tautcony/qart/internal/utils"
	"github.com/tautcony/qart/models/request"
	"github.com/tautcony/qart/models/response"
	"image"
	"image/png"
	"log"
)

type UploadController struct {
	base.QArtController
}

type RenderController struct {
	base.QArtController
}

// @Title Upload image
// @Description Upload image for further operation
// @Success 200 {object} models.response.BaseResponse
// @Param   image     formData   string true       "upload file name"
// @router / [post]
func (c *UploadController) Post() {
	f, header, err := c.GetFile("image")
	if err != nil {
		log.Println("get file err ", err)
		c.Fail(nil, 1, err.Error())
		return
	}
	log.Println("get file", header.Filename, "with size", header.Size)

	img, err := utils.GetImageThumbnail(f)
	defer f.Close()
	if err != nil {
		log.Println("down sampling err ", err)
		c.Fail(nil, 2, err.Error())
		return
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
	tag := fmt.Sprintf("%x", sha256.Sum256(buf.Bytes()))
	c.SetSession(sessionutils.SessionKey(tag, "image"), img) // store image data in session

	c.JSON(&response.BaseResponse{
		Data: struct {
			Id string `json:"id"`
		}{
			tag,
		},
		Success: true,
		Code:    0,
		Message: "0",
	})
}

func (c *RenderController) Post() {
	operation, err := request.NewOperation()
	if err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, operation); err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
	sessionKey := sessionutils.SessionKey(operation.Image, "image")
	if operation.Image == "default" && c.GetSession(sessionKey) == nil {
		data, _, _ := utils.Read(utils.GetUploadPath("default.png"))
		defaultImage, err := png.Decode(bytes.NewBuffer(data))
		if err == nil {
			c.SetSession(sessionKey, defaultImage)
		}
	}

	sessionData := c.GetSession(sessionKey)
	if sessionData == nil {
		c.Fail(nil, 2, "image not found, please upload first")
		return
	}
	uploadImage := sessionData.(image.Image)
	img, err := qr.Draw(operation, uploadImage)
	if err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
	var data []byte
	switch {
	case img.SaveControl:
		data = img.Control
	default:
		data = img.Code.PNG()
	}
	c.SetSession(sessionutils.SessionKey(operation.Image, "config"), img)
	if c.GetString("debug") == "1" {
		c.Ctx.Output.ContentType(".png")
		err = c.Ctx.Output.Body(data)
		if err != nil {
			panic(err)
		}
		return
	}

	c.Success(struct {
		Image string `json:"image"`
	}{
		"data:image/png;base64," + base64.StdEncoding.EncodeToString(data),
	}, 0)
}
