package controller
import (
	"github.com/astaxie/beego"
)
type FistController struct {
	beego.Controller
}
func (this *FistController) Get() {
	this.Ctx.WriteString("FistController **************")
}