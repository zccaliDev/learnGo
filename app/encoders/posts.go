package encoders

import (
	"io"
	"github.com/zccaliDev/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodePost(body io.ReadCloser) (post models.Post) {
	var data, _ = ioutil.ReadAll(body);

	if err := json.Unmarshal(data, &post); err != nil {
		log.Println(err)
		return
	}

	return
}
