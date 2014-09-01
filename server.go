package main

import (
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"time"
)

func main() {
	initDb()
	defer X.Close()

	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/inventories", AllInventories)
	m.Get("/inventories/:code", Get)

	//martini整合到现有服务当中
	http.Handle("/", m)
	//m.Run()
	//修改端口
	http.ListenAndServe(":8080", m)
}

func AllInventories(r render.Render) {
	r.JSON(200, GetAllInventory())
}

func Inventory(params martini.Params, r render.Render) {

}
