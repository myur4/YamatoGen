package main

import (
	"fmt"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
	"github.com/liamg/tml"
	"github.com/mbndr/figlet4go"
	"log"
	"time"
)

func init() {
	launcher.NewBrowser().MustGet()
}

type Config struct {
	simBrand   string
	fiveSimKey string
	smsRuKey   string
	webhook    string
}

type config struct {
	simBrand   string `env:"PORT"`
	fiveSimKey string `env:"PORT"`
	smsRuKey   string `env:"PORT"`
	webhook    string `env:"PORT" envDefault:"3000"`
	cooldown   time.Duration
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	figlet := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "invita"
	options.FontColor = []figlet4go.Color{figlet4go.ColorCyan}

	renderStr, _ := figlet.RenderOpts("Yamato", options)

	fmt.Println(renderStr)

	tml.Println("<yellow>[AUTENTICANDO]</yellow> Iniciando autenticação com o servidor de login")
	//	tml.Printf("<blue>[QUESTION]</blue> <bold>Informe o número de threads (max. 10):</bold>")
	//devtools.StartBrowser()

	//SMS.CheckBalance()
	//SMS.FiveBuySIM()
	fmt.Scanln()
}
