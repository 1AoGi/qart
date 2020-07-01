package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"qart/controllers/base"
	"qart/models/request"
	"qart/models/response"
	"qart/qrweb/qr"
	"qart/qrweb/utils"
	"time"
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
	fmt.Println("get file", header.Filename, "with size", header.Size)

	buf, err := utils.GetImageThumbnail(f)
	defer f.Close()
	if err != nil {
		log.Println("downsample err ", err)
		c.Fail(nil, 2, err.Error())
		return
	}

	tag := fmt.Sprintf("%x", sha256.Sum256(buf.Bytes()))
	filePath := utils.GetUploadPath(tag + ".png")
	err = utils.Write(filePath, buf.Bytes())
	if err != nil {
		log.Println("write file err ", err)
		c.Fail(nil, 3, err.Error())
		return
	}
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

	t0 := time.Now()
	data, err := qr.Draw(operation)
	t1 := time.Now()
	log.Printf("render in %s\n", t1.Sub(t0).String())
	if err != nil {
		c.Fail(nil, 2, err.Error())
		return
	}
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
