package encoders

import (
	"io"
	"github.com/zccaliDev/learnGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeSingleUsers(body io.ReadCloser) (user models.User) {
	var data, _ = ioutil.ReadAll(body);

	if err := json.Unmarshal(data, &user); err != nil {
		log.Println(err)
		return
	}

	return
}

