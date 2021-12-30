package modules

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"github.com/imroc/req"
	"github.com/liamg/tml"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type TempMail struct {
	Email string
	Token string
}

func authenticate() {

}

func sendAccount(content string) {
	result, err := req.Post("WEBHOOK URL")
	if err != nil {
		tml.Println("<red>[ERRO]</red> Verifique o status do seu webhook")
		return
	}

	if result.Response().StatusCode == 200 {
		tml.Println("<purple>[SEND]</purple> Conta enviada pro seu servidor")
	} else {
		tml.Println("<red>[FAIl]</red> Não foi possível identificar se a mensagem do webhook foi enviada")
	}
}

func GenMail() TempMail {
	url := "https://api.internal.temp-mail.io/api/v3/email/new"

	payload := strings.NewReader(fmt.Sprintf("{ \"domain\": \"kazinsky111.com\", \n \"name\": \"%s\" }", RandStringRunes(32)))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	email := gjson.Get(string(body), "email").String()
	token := gjson.Get(string(body), "token").String()

	return TempMail{Token: token, Email: email}
}

func solveCaptcha() string {
	client := api2captcha.NewClient(os.Getenv("2CAPTCHA"))
	cap := api2captcha.HCaptcha{
		SiteKey: "f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34",
		Url:     "https://mysite.com/captcha.html",
	}
	req := cap.ToRequest()
	code, err := client.Solve(req)
	if err != nil {

	}

	return code
}

/*
embeds example

{
  "embeds": [
    {
      "title": "Meow!",
      "color": 1127128
    },
    {
      "title": "Meow-meow!",
      "color": 14177041
    }
  ]
}
*/
