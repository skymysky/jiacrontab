package handle

import (
	"jiacrontab/libs/mailer"
	"jiacrontab/server/conf"
	"net/http"
	"strings"

	"github.com/kataras/iris"
)

func ViewConfig(ctx iris.Context) {

	c := conf.Category()
	r := ctx.Request()

	if r.Method == http.MethodPost {
		mailTo := strings.TrimSpace(r.FormValue("mailTo"))
		mailer.SendMail([]string{mailTo}, "测试邮件", "测试邮件请勿回复！")
	}

	ctx.ViewData("configs", c)
	ctx.View("viewConfig.html")
}

func ReloadConfig(ctx iris.Context) {
	conf.Reload()
	ctx.Redirect("/", http.StatusFound)

}
