package method

import (
	"bytes"
	"encoding/json"
	"github.com/iluxaorlov/qrbot/entity"
	"io/ioutil"
	"net/http"
)

type request struct {
	Offset int `json:"offset"`
}

func GetUpdates(api string, offset int) ([]entity.Update, error) {
	request := request{
		Offset: offset,
	}

	js, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(api + "/getUpdates", "application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response entity.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}