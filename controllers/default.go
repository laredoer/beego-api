package controllers

// Operations about default
type DefaultController struct {
	BaseController
}
// @Title welcome
// @Description get welcome info
// @Success 200 {object} Response
// @router / [get]
func (this *DefaultController) GetAll() {
	this.Data["json"] = Response{0, "success.", "API 1.0"}
	this.ServeJSON()
}