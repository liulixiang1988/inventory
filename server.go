package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/liulixiang1988/inventory/models"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func main() {
	models.InitDb()
	defer models.X.Close()

	m := martini.Classic()

	m.Use(render.Renderer())
	m.Use(auth.BasicFunc(ValidateUser))

	m.Get("/inventories", AllInventories)
	m.Get("/inventories/:code", GetInventory)
	m.Post("/material_inv", binding.Bind(models.Material_Inventory{}), AddMaterial_Inv)

	//martini整合到现有服务当中
	http.Handle("/", m)
	//m.Run()
	//修改端口
	if err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", m); err != nil {
		log.Fatalln(err)
	}
}

func ValidateUser(username, password string) bool {
	user, err := models.GetUser(username)
	if err != nil {
		log.Println("用户登录出错", err)
		return false
	}
	if user == nil {
		return false
	}
	return auth.SecureCompare(username, user.User_id) && auth.SecureCompare(password, user.Password)
}

func AllInventories(r render.Render) {
	r.JSON(200, models.GetAllInventory())
}

func GetInventory(params martini.Params, r render.Render) {
	code := params["code"]
	inv, err := models.GetInventory(code)
	if err != nil {
		r.JSON(http.StatusInternalServerError, err)
	}
	r.JSON(http.StatusOK, inv)
}

func AddMaterial_Inv(m_inv models.Material_Inventory, r render.Render) {
	affected, err := models.AddMaterial_Inventory(m_inv)
	if err != nil {
		r.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error",
			"message": err})
	}
	msg := fmt.Sprintf("成功插入%d条数据", affected)
	r.JSON(http.StatusOK, map[string]interface{}{"status": "ok",
		"message": msg})
}
