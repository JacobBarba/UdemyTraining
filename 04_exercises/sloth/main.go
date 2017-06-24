package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
	//"gopkg.in/headzoo/surf.v1"
	//"github.com/headzoo/surf"
	//"net/http"
)

var executor = "http://localhost:4444/wd/hub"

func click() {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, executor)

	if err != nil {
		panic(err)
	}

	wd.Get("https://u5378697.ct.sendgrid.net/wf/click?upn=avbkyTCefrQKo2L1u1sfmxkNfjyLNUO6PxYgQgpw226udSqC4-2FNl-2FJqjszNrubKYGQYTjmx3rzQi-2BPUw3mLRjVg-2FfHrGpJvIic8FHPmzyZ9vv4ZrAQFdTdbCLZgWCMJJ-2FP28FTTbbswxC7j6wc8heAW5BD2T6AZfWTCBnxUr6ZwUdv6fw5AWFberOmPDtastEifXsfWbUbdvJmx6XZ0xzWZ9zO6TKCnvkUzNT7EFr2Uwb27LH7ZFGRjU-2Fo-2B08Ah8ca41VO3CBo5X6141O87eLhiQU7psnI6rRhIdhAIH8-2B9p5TOLBCt1k-2B5X4JS8nuGs89-2Bu1msprf5IFBtJ3tZaKTWl9aVT9GtQoRoHrPYmhFfOB9bF-2FU-2Fb9bxSEpPZB-2FngFhryYJ2c1JZN0QZoYCoJTNJ-2BKiwLsS4JC4TukPUzU9bQOmRBDtFUt8ahM3x0lIVCRxae5PDGuXACS3QUHugfVA-3D-3D_SrLupllVV2a7IZXnUPzJ4PoX2AtCxvxlbg91u8CK0U22EQcNoud66koIlOC2Tbt03E-2Fz9yc3R-2BBlnN6iPnymfR6WtU001XfZojsjlt4UM6RGntX7iXPRFUTR8qOZaBa-2FaiP6CbsAK3NA2Q2eoveLkdz7XeqwOWetC6SEoQVXiKN1tzBvDU-2FLIgToZxdc0OiAaN8ycofDxkWzPOFjFA1UEb2osCx-2BdpBykYSuj2i3o4E-3D")
	btn, _ := wd.FindElement(selenium.ByClassName, "style__button___3HJEA")

	for true {
		btn.Click()
	}

	time.Sleep(1)

	wd.Close()
	fmt.Println("Done")

}

func main() {
	for i := 0; i < 15; i++{
		go click()
	}

	for true{
		time.Sleep(1)
	}
}
