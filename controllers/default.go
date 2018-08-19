package controllers

// Operations about   版本信息
type DefaultController struct {
	BaseController
}
// @Title version
// @Description get version info
// @Success 200 {object} controllers.Response
// @router / [get]
func (this *DefaultController) GetAll() {
	this.Data["json"] = Response{0, "success.", "API 1.0"}
	this.ServeJSON()
}