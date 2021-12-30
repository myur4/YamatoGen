package SMS

import (
	"fmt"
	"github.com/liamg/tml"
	"net/http"
	"os"
	"strings"
)

var smsNumber string
var smsID string
var phoneCode string

func sendActivationCode(apiSelection, apiKey string) {
	switch apiSelection {
	case "1":
	}
}

func requestSmsCode(phoneNumber, discordToken, solve2Captcha string) {
	url := "https://discord.com/api/v9/users/@me/phone"

	payload := strings.NewReader(fmt.Sprintf(`{ "change_phone_reason": "%s", "captcha_key": "%s","phone": "%s"}`, "user_action_required", code, phoneNumber))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:96.0) Gecko/20100101 Firefox/96.0")
	req.Header.Add("Accept", "/")
	req.Header.Add("Accept-Language", "pt-BR")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", discordToken)
	req.Header.Add("X-Super-Properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRmlyZWZveCIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJwdC1CUiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQ7IHJ2Ojk2LjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvOTYuMCIsImJyb3dzZXJfdmVyc2lvbiI6Ijk2LjAiLCJvc192ZXJzaW9uIjoiMTAiLCJyZWZlcnJlciI6IiIsInJlZmVycmluZ19kb21haW4iOiIiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTA4NDcxLCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==")
	req.Header.Add("X-Discord-Locale", "pt-BR")
	req.Header.Add("X-Debug-Options", "bugReporterEnabled")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cords")
	req.Header.Add("Sec-Fetch-Site", "same-origin")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 204 {
		tml.Println(`<red>[SMS ERROR 1]</red> O discord não aceitou a verificação de número`)
		submitSms("", os.Getenv("SENHA"))
	} else {
		tml.Println(`<green>[SMS AWAIT]</green> <bold>SMS encaminhado, esperando resposta do SIM</bold>`)
	}
}

func submitSms(phoneCode, accountPassword string) {

	tml.Println(`<yellow>[SMS AWAIT]</yellow> <bold>Enviando código de SMS pro Discord</bold>`)

	url := "https://discord.com/api/v9/users/@me/phone"

	payload := strings.NewReader(fmt.Sprintf(`"change_phone_reason": "user_action_required", "password": "%s", "phone_token": "%s"`, accountPassword, phoneCode))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("referer", "https://discord.com/channels/@me,")
	req.Header.Add("Authorization", "token")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 204 {
		tml.Println(`<red>[SMS ERROR 2]</red> O discord não aceitou a verificação de número`)
	} else {
		tml.Println(`<green>[SMS SUCCESS]</green> <bold>Telefone adicionado com sucesso</bold>`)
	}
}
