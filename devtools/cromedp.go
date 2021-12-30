package devtools

import (
	"YamatoGen/modules"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/imroc/req"
	"github.com/liamg/tml"
	"os"
	"time"
)

func StartBrowser() {
	launch := launcher.New().
		Set("blink-settings", "imagesEnabled=false").
		Devtools(false).
		Headless(false).
		MustLaunch()

	browser := rod.New().ControlURL(launch).Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	MakeAccount(browser, "bunda123oi1@gmail.com", "mariah lc1", os.Getenv("SENHA"), "03092001")
}

func MakeAccount(chrome *rod.Browser, email, username, password, date string) {
	page := chrome.MustPage("https://discord.com/register/")
	page.MustEvalOnNewDocument("document.body.appendChild(Object.assign(document.createElement('script'), {src: 'https://gitcdn.xyz/repo/berstend/puppeteer-extra/stealth-js/stealth.min.js'}))")

	page.MustWaitRequestIdle("https://discord.com/register")

	// Elementos da página
	emailInput := page.MustElement("input[type=\"email")
	usernameInput := page.MustElement("input[aria-label=\"Username\"")
	passwordInput := page.MustElement("input[name=\"password\"")

	// Date selector
	monthSelector := page.MustElement("input[aria-label=\"Month\"")
	daySelector := page.MustElement("input[aria-label=\"Day\"")
	yearSelector := page.MustElement("input[aria-label=\"Year\"")

	emailInput.MustInput(email)
	usernameInput.MustInput(username)
	passwordInput.MustInput(password)

	// Button
	submitButton := page.MustElement("button[type=\"submit\"")

	monthSelector.Focus()
	page.Keyboard.Press('a')

	daySelector.Focus()
	page.Keyboard.Press('3')

	yearSelector.Focus()
	page.Keyboard.Press('1')
	page.Keyboard.Press('9')
	page.Keyboard.Press('9')
	page.Keyboard.Press('9')

	submitButton.MustClick()

	time.Sleep(time.Second * 5)

	page.Keyboard.Press(input.Tab)
	page.Keyboard.Press(input.Enter)

	page.MustWaitRequestIdle()
	page.MustWaitNavigation()
	page.MustWaitRequestIdle()

	ChangeEmail(page, chrome)
}

func ChangeEmail(page *rod.Page, chrome *rod.Browser) {

	tml.Println(`<blue>[MAIL]</blue> Iniciando confirmação de e-mail)`)
	time.Sleep(time.Second * 3)

	page.MustWaitRequestIdle("discord.com")

	closeModal := page.MustElement("#app-mount > div:nth-child(7) > div.layer-2KE1M9 > div > div > div.content-1KT39n.theme-light > button > div > svg")
	closeModal.MustClick()

	// Buttons
	settingsButton := page.MustElement("#app-mount > div.app-1q1i1E > div > div.layers-3iHuyZ.layers-3q14ss > div > div > div > div > div.sidebar-2K8pFh > section > div.container-3baos1 > div.flex-1xMQg5.flex-1O1GKY.horizontal-1ae9ci.horizontal-2EEEnY.flex-1O1GKY.directionRow-3v3tfG.justifyStart-2NDFzi.alignStretch-DpGPf3.noWrap-3jynv6 > button:nth-child(3) > div > svg")
	settingsButton.MustClick()

	time.Sleep(time.Second * 1)

	editEmailButton := page.MustElement("#app-mount > div.app-1q1i1E > div > div.layers-3iHuyZ.layers-3q14ss > div:nth-child(2) > div > div.contentRegion-3nDuYy > div > div > main > div > div:nth-child(1) > div.children-rWhLdy > div > div.background-1QDuV2 > div > div:nth-child(2) > button")
	editEmailButton.MustClick()

	time.Sleep(time.Second * 1)

	submitButton := page.MustElement("#app-mount > div:nth-child(7) > div.layer-2KE1M9 > div > div > form > div.flex-1xMQg5.flex-1O1GKY.horizontalReverse-2eTKWD.horizontalReverse-3tRjY7.flex-1O1GKY.directionRowReverse-m8IjIq.justifyStart-2NDFzi.alignStretch-DpGPf3.noWrap-3jynv6.footer-2gL1pp > button.button-38aScr.lookFilled-1Gx00P.colorBrand-3pXr91.sizeMedium-1AC_Sl.grow-q77ONN")

	time.Sleep(time.Second * 1)

	// Input
	emailInput := page.MustElement("#app-mount > div:nth-child(7) > div.layer-2KE1M9 > div > div > form > div.content-1LAB8Z.content-2Cyhe6.thin-1ybCId.scrollerBase-289Jih > div:nth-child(1) > div > input")
	passwordInput := page.MustElement("#app-mount > div:nth-child(7) > div.layer-2KE1M9 > div > div > form > div.content-1LAB8Z.content-2Cyhe6.thin-1ybCId.scrollerBase-289Jih > div.password-3REpJO > div > input")

	time.Sleep(time.Second * 1)

	tempMail := modules.GenMail()
	mailEndpoint := fmt.Sprintf("https://api.internal.temp-mail.io/api/v3/email/%s/messages", tempMail.Email)

	time.Sleep(time.Second * 1)

	emailInput.MustInput(tempMail.Email)
	passwordInput.MustInput(os.Getenv("SENHA"))

	submitButton.MustClick()

	tempInbox, _ := req.Get(mailEndpoint)

	for true {
		tempInbox, _ := req.Get(mailEndpoint)
		tml.Println(`<yellow>[MAIL AWAITING]</yellow> Esperando e-mail de confirmação.`)
		if len(tempInbox.String()) > 0 {
			tml.Println(`<magenta>[MAIL RECIEVED]</magenta> E-mail de confirmação recebido, prosseguindo para a confirmação.`)
			break
		}
		time.Sleep(time.Second * 5)
	}

	emailContent := tempInbox.String()
	confirmationURL, _ := modules.Substr(emailContent, "Email: ", "\\n")

	confirmationPage := chrome.MustPage(confirmationURL)
	confirmationPage.MustClose()

	tml.Println("<green>[MAIL SUCCESS]</green> Conta confirmada. Enviando webhook.")
}

func SimActivation() {
	tml.Println("<blue>[SIM]</blue> <bold>Verificação de SMS iniciada</bold>")

}

// BOTAO CONFIG
