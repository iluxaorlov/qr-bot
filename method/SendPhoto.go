package method

import (
	"bytes"
	"github.com/iluxaorlov/qrbot/entity"
	"github.com/skip2/go-qrcode"
	"image"
	png2 "image/png"
	"mime/multipart"
	"net/http"
	"strconv"
)

type photo struct {
	ChatId int `json:"chat_id"`
	Photo string `json:"photo"`
}

func SendPhoto(api string, update entity.Update) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	err := writer.WriteField("chat_id", strconv.Itoa(update.Message.Chat.Id))
	if err != nil {
		return err
	}

	photo, _ := writer.CreateFormFile("photo", "qr.png")

	png, err := qrcode.Encode(update.Message.Text, qrcode.Highest, 256)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewBuffer(png))
	if err != nil {
		return err
	}

	err = png2.Encode(photo, img)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", api + "/sendPhoto", buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := new(http.Client)
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}