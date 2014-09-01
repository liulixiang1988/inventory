package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"log"
	"time"
)

type Inventory struct {
	Id               int64     `json:"-"`
	Inventory_code   string    `json:"inventory_code"`
	Inventory_name   string    `json:"inventory_name"`
	Remark           string    `json:"remark"`
	Create_time      time.Time `json:"create_time"`
	Create_user_id   string    `json:"create_user_id"`
	Create_user_name string    `json:"create_user_name"`
	Update_time      time.Time `json:"update_time"`
	Update_user_id   string    `json:"update_user_id"`
	Update_user_name string    `json:"update_user_name"`
	Is_available     bool      `json:"is_available"`
}

type Material_Inventory struct {
	Id               int64     `json:"-"`
	Item_code        string    `json:"item_code"`
	Inventory_code   string    `json:"inventory_code"`
	Create_time      time.Time `json:"create_time"`
	Create_user_id   string    `json:"create_user_id"`
	Create_user_name string    `json:"create_user_name"`
	Update_time      time.Time `json:"update_time"`
	Update_user_id   string    `json:"update_user_id"`
	Update_user_name string    `json:"update_user_name"`
	Is_available     bool      `json:"is_available"`
}

var X *xorm.Engine

func InitDb() {
	const conStr = "driver={sql server};server=127.0.0.1;prot=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=Data_Center"
	var err error
	X, err = xorm.NewEngine("odbc", conStr)
	if err != nil {
		log.Println("连接数据库错误：", err)
	}
	err = X.Sync(new(Inventory))
	if err != nil {
		log.Println("同步Inventory表结构时出错", err)
	}
	err = X.Sync(new(Material_Inventory))
	if err != nil {
		log.Println("同步Material_Inventory表结构时出错", err)
	}
}

func GetAllInventory() []*Inventory {
	inventories := make([]*Inventory, 0)
	err := X.Find(&inventories)
	if err != nil {
		return nil
	}
	return inventories
}

func GetInventory(code string) (*Inventory, error) {
	inventory := &Inventory{Inventory_code: code}
	has, err := X.Get(inventory)
	if err != nil {
		return nil, err
	}
	if has == false {
		return nil, nil
	}
	return inventory, nil
}

func AddMaterial_Inventory(meterial_inv Material_Inventory) (int64, error) {
	affected, err := X.Insert(meterial_inv)
	return affected, err
}
