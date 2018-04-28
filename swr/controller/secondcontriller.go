package controller
import (
	"github.com/astaxie/beego"
)
type SecondController struct {
	beego.Controller
}
func (this *SecondController) Get() {
	this.Ctx.WriteString("SecondController **************")
}