package controllers

import (
	"crypto/sha256"
	"fmt"
	"log"
	"qart/controllers/base"
	"qart/models/response"
	"qart/qrweb/utils"
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
	fmt.Println(c.GetString("key"))
	c.Success(struct {
		Key string `json:"key"`
	}{"1"}, 0)
}
