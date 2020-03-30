// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mn_log/controllers"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func setLoginUser() {
	var setUserId = func(ctx *context.Context) {
		inner_uid, _ := strconv.ParseInt(ctx.Input.Header("uid"), 10, 64)
		inner_muid, _ := strconv.ParseInt(ctx.Input.Header("muid"), 10, 64)
		//inner_resource_id, _ := strconv.ParseInt(ctx.Input.Header("resourceId"), 10, 64)
		//|| inner_resource_id <= 0
		if inner_uid <= 0 || inner_muid <= 0 {
			resp := &controllers.Response{
				Code: 10005,
				Data: map[string]interface{}{},
				Msg:  beego.AppConfig.String("code::10005"),
			}

			hasIndent := beego.BConfig.RunMode != beego.PROD
			ctx.Output.JSON(resp, hasIndent, false)

			return

		}

		ctx.Input.SetData("inner_uid", inner_uid)
		ctx.Input.SetData("inner_muid", inner_muid)
		//ctx.Input.SetData("inner_resource_id", inner_resource_id)

	}
	beego.InsertFilter("/log/*", beego.BeforeStatic, setUserId)
}

func init() {
	beego.Router("/log/list", &controllers.LogController{}, "get:List")
	beego.Router("/order/list", &controllers.OrderController{}, "get:List")
	setLoginUser()
}
