// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bondi-push-notification/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.RESTRouter("/v1/getVapidPublicKey", &controllers.VapidController{})
	beego.RESTRouter("/v1/savePushSubscription", &controllers.SubscriptionController{})
	beego.RESTRouter("/v1/sendPushNotification/:messageId", &controllers.SendPushNotification{})
}
