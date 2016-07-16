package encoders

import (
	"io"
	"github.com/zccaliDev/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeComment(body io.ReadCloser) (comment models.Comment) {
	var data, _ = ioutil.ReadAll(body);

	if err := json.Unmarshal(data, &comment); err != nil {
		log.Println(err)
		return
	}

	return
}
