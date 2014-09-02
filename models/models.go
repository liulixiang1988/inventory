package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"log"
	"time"
)

type User struct {
	Id           int64     `json:"-" `
	User_id      string    `json:"user_id" xorm:"varchar(50)"`
	User_name    string    `json:"user_name" xorm:"nvarchar(255)"`
	Password     string    `json:"-"  xorm:"nvarchar(255)"`
	Is_active    bool      `json:"is_active"`
	Is_superuser bool      `json:"-"`
	Is_stuff     bool      `json:"is_stuff"`
	Email        string    `json:"email" xorm:"nvarchar(255)"`
	Last_login   time.Time `json:"-"`
}

type Inventory struct {
	Id               int64     `json:"-"`
	Inventory_code   string    `json:"inventory_code" xorm:"nvarchar(255)"`
	Inventory_name   string    `json:"inventory_name" xorm:"nvarchar(255)"`
	Remark           string    `json:"remark"`
	Create_time      time.Time `json:"create_time"`
	Create_user_id   string    `json:"create_user_id"  xorm:"nvarchar(255)"`
	Create_user_name string    `json:"create_user_name"  xorm:"nvarchar(255)"`
	Update_time      time.Time `json:"update_time"`
	Update_user_id   string    `json:"update_user_id"  xorm:"nvarchar(255)"`
	Update_user_name string    `json:"update_user_name"  xorm:"nvarchar(255)"`
	Is_available     bool      `json:"is_available"`
}

type Material_Inventory struct {
	Id               int64     `json:"-"`
	Item_code        string    `json:"item_code"  xorm:"nvarchar(255)"`
	Inventory_code   string    `json:"inventory_code"  xorm:"varchar(128)"`
	Create_time      time.Time `json:"create_time"`
	Create_user_id   string    `json:"create_user_id"  xorm:"nvarchar(255)"`
	Create_user_name string    `json:"create_user_name"  xorm:"nvarchar(255)"`
	Update_time      time.Time `json:"update_time"`
	Update_user_id   string    `json:"update_user_id"  xorm:"nvarchar(255)"`
	Update_user_name string    `json:"update_user_name"  xorm:"nvarchar(255)"`
	Is_available     bool      `json:"is_available"`
}

var X *xorm.Engine

func InitDb() {
	const conStr = "driver={sql server};server=127.0.0.1;prot=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=Data_Center"
	var err error
	X, err = xorm.NewEngine("odbc", conStr)
	if err != nil {
		log.Fatalln("连接数据库错误：", err)
	}
	err = X.Sync(new(User))
	if err != nil {
		log.Fatalln("同步用户表结构时出错", err)
	}
	err = X.Sync(new(Inventory))
	if err != nil {
		log.Fatalln("同步Inventory表结构时出错", err)
	}
	err = X.Sync(new(Material_Inventory))
	if err != nil {
		log.Fatalln("同步Material_Inventory表结构时出错", err)
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

func AddMaterial_Inventory(material_inv Material_Inventory) (int64, error) {
	m_inv := &Material_Inventory{}

	has, err := X.Where("Item_code=?", material_inv.Item_code).Get(m_inv)
	if err != nil {
		log.Println("插入Material_Inventory时出错", err)
	}
	if has == false {
		log.Println("物料：", material_inv.Item_code, material_inv.Inventory_code)
		affected, err := X.Insert(material_inv)
		return affected, err
	} else {
		material_inv.Id = m_inv.Id
		material_inv.Update_time = time.Now()
		log.Println("更新物料：", material_inv.Item_code, material_inv.Inventory_code)
		affected, err := X.Id(m_inv.Id).Update(&material_inv)
		return affected, err
	}
}

func AddUser(user User) (int64, error) {
	affected, err := X.Insert(user)
	return affected, err
}

func GetUser(user_id string) (*User, error) {
	user := &User{User_id: user_id}
	has, err := X.Get(user)
	if err != nil {
		return nil, err
	}
	if has == false {
		return nil, nil
	}
	return user, nil
}
