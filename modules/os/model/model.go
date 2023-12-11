package osmodel

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Ipaddress struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
	Ip          string `json:"ip"`
	Netmask     string `json:"netmask"`
}

func (ia *Ipaddress) Genline() string {
	if nm := string(ia.Netmask[0]); nm == "/" {
		return fmt.Sprintf("%s%s", ia.Ip, ia.Netmask)
	}
	return fmt.Sprintf("%s %s", ia.Ip, ia.Netmask)
}

type Keyval struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Separator string `json:"separator"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

func (kv *Keyval) Genline() string {
	return fmt.Sprintf("%s %s %s", kv.Name, kv.Separator, kv.Value)
}

type Interface struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Ip      string `json:"ip"`
	Netmask string `json:"netmask"`
	Custom  string `json:"custom"`
}

func (iface *Interface) Genline() string {
	if c := string(iface.Custom[0]); c != "#" {
    if nm := string(iface.Netmask[0]); nm == "/" {
      return fmt.Sprintf("inet %s%s\n%s", iface.Ip, iface.Netmask, iface.Custom)
    }
		return fmt.Sprintf("inet %s %s\n%s", iface.Ip, iface.Netmask, iface.Custom)
	}
  if nm := string(iface.Netmask[0]); nm == "/" {
    return fmt.Sprintf("%s%s", iface.Ip, iface.Netmask)
  }
	return fmt.Sprintf("inet %s %s", iface.Ip, iface.Netmask)
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&Ipaddress{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ipaddress migrated")
	err = db.AutoMigrate(&Keyval{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Keyval migrated")
	err = db.AutoMigrate(&Interface{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Interface migrated")
}
