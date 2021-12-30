package SMS

import (
	"fmt"
	"github.com/liamg/tml"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FiveBuySIM() string {
	url := "https://5sim.net/v1/user/buy/activation/any/any/discord"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+string(os.Getenv("FIVESIMKEY")))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	orderId := gjson.Get(string(body), "id").String()
	orderStatus := gjson.Get(string(body), "STATUS")
	phoneNumberRaw := gjson.Get(string(body), "phone")

	fmt.Println(orderStatus, orderId, phoneNumberRaw)

	tml.Println(tml.Sprintf("<blue>[5SIM BOUGHT]</blue> <yellow>[%d]</yellow> Número comprado, aguardando SMS do Discord", orderId))

	phoneNumber := gjson.Get(string(body), "phone")
	phoneCountryCode := gjson.Get(string(body), "country")
	phonePrefix := gjson.Get(string(body), "prefix")

	fmt.Println(phonePrefix, phoneNumber, phoneCountryCode)

	code := ""

	for true {
		url := "https://5sim.net/v1/user/check/" + orderId

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Authorization", "Bearer "+string(os.Getenv("FIVESIMKEY")))

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		codeArr := gjson.Get(string(body), "sms.#.code").Array()
		code = codeArr[0].String()

		fmt.Println(code)

		if len(code) > 0 {
			tml.Println("<green>[SMS RECIEVED]</green> SMS de ativação recebido")
			break
		}

		time.Sleep(time.Second * 5)
	}
	return code
}
