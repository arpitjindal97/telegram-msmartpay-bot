package main

import (
	"github.com/tebeka/selenium"
	"time"
	//"os"
	"fmt"
)
var msmartpay_email = "mukeshkjindal72@gmailcom"//os.Getenv("msmartpay_email")
var msmartpay_password = "Mukesh@1234"//os.Getenv("msmartpay_password")
func main1(vals [3]string) string {

	caps := selenium.Capabilities{"browserName": "firefox",
	"webdriver.gecko.driver":"/home/arpit/Desktop/geckodriver"}
	wd, err := selenium.NewRemote(caps, "")
	wd.SetAsyncScriptTimeout(5000)

	defer wd.Quit()

	if err := wd.Get("http://super.msmartpay.in/superadmin/login.jsp"); err != nil {
		panic(err)
	}

	elem,err := (wd.FindElement(selenium.ByName,"userName"))
	if err !=nil {
		panic(err)
	}

	elem.SendKeys(msmartpay_email)
	fmt.Println("Email: "+msmartpay_email+"\nPassword: "+msmartpay_password)
	elem,err = (wd.FindElement(selenium.ByName,"password"))
	elem.SendKeys(msmartpay_password)

	elem,err = wd.FindElement(selenium.ByName,"Submit")
	elem.Click()

	time.Sleep(2*time.Second)

	wd.Get("http://super.msmartpay.in/superadmin/TBTransfer.action")

	elem,err = wd.FindElement(selenium.ByID,"reqId")
	if err !=nil {
		panic(err)
	}
	elem.SendKeys(vals[0])

	elem,err = wd.FindElement(selenium.ByID,"requser")
	elem.SendKeys(vals[2])

	elem,err = wd.FindElement(selenium.ByClassName,"cls_btn")
	elem.Click()

	time.Sleep(1*time.Second)

	elem,err = wd.FindElement(selenium.ByName,"company")

	return_str,err := elem.GetAttribute("value")

	var elements, _ = wd.FindElements(selenium.ByXPATH,"//td[@valign='middle']")

	return_str = "Amount before: "+extractAmount(elements)+"\nCompany: "+return_str

	elem,err = wd.FindElement(selenium.ByName,"amount")
	elem.SendKeys(vals[1])

	elem,err = wd.FindElement(selenium.ByName,"intremark")
	elem.SendKeys("sent via bot")

	elem,err = wd.FindElement(selenium.ByName,"extremark")
	elem.SendKeys("made by arpit")

	elements,_ = wd.FindElements(selenium.ByClassName,"cls_btn")
	elem = elements[1]
	elem.Click()

	time.Sleep(1*time.Second)

	elements, _ = wd.FindElements(selenium.ByXPATH,"//td[@valign='middle']")

	return_str  = return_str + ("\nAmount after: "+extractAmount(elements))

	return return_str
}

func extractAmount(elements []selenium.WebElement) string {
	elem := elements[8]
	amount,_ := elem.GetAttribute("innerHTML")

	var new_str string

	for _, r := range amount {
		if string(r) == " " {
			continue
		}
		new_str = new_str + string(r)
	}
	return new_str
}