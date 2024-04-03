package pfmodel

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Pfrulebasic struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Action      string `json:"action" form:"action"`
	Direction   string `json:"direction"  form:"direction"`
	Interface   string `json:"interface" form:"interface"`
	Addrfamily  string `json:"addrfamily" form:"addrfamily"`
	Protocol    string `json:"protocol" form:"protocol"`
	Source      string `json:"source" form:"source"`
	Destination string `json:"destination" form:"destination"`
	Dstport     string `json:"dstport" form:"dstport"`
	Custom      string `json:"custom" form:"custom"`
}

func (pfrb *Pfrulebasic) Genline() string {
	if c := string(pfrb.Custom[0]); c != "#" {
		return fmt.Sprintf("%s %s on %s %s proto %s from %s to %s port %s %s",
			pfrb.Action, pfrb.Direction, pfrb.Interface,
			pfrb.Addrfamily, pfrb.Protocol, pfrb.Source, pfrb.Destination, pfrb.Dstport, pfrb.Custom)
	}
	return fmt.Sprintf("%s %s on %s %s proto %s from %s to %s port %s",
		pfrb.Action, pfrb.Direction, pfrb.Interface, pfrb.Addrfamily, pfrb.Protocol, pfrb.Source, pfrb.Destination, pfrb.Dstport)
}

type Pfqueue struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name" form:"name"`
	Onorparent    string `json:"onorparent" form:"onorparent"`
	Parentoriface string `json:"parentoriface" form:"parentorinterface"`
	Bandwidth     string `json:"bandwidth" form:"bandwidth"`
	Custom        string `json:"custom" form:"custom"`
}

type InternetPlan struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" form:"name"`
	Download int    `json:"download" form:"download"`
	Upload   int    `json:"upload" form:"upload"`
	Burst    int    `json:"burst" form:"burst"`
	Duration int    `json:"duration" form:"duration"`
}

func (pfq *Pfqueue) Genline() string {
	if c := string(pfq.Custom[0]); c != "#" {
		return fmt.Sprintf("queue %s %s %s bandwidth  %s %s", pfq.Name, pfq.Onorparent, pfq.Parentoriface, pfq.Bandwidth, pfq.Custom)
	}
	return fmt.Sprintf("queue %s %s %s bandwidth  %s", pfq.Name, pfq.Onorparent, pfq.Parentoriface, pfq.Bandwidth)
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&Pfrulebasic{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pfrulebasic migrated")
	err = db.AutoMigrate(&Pfqueue{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pfqueue migrated")
	err = db.AutoMigrate(&InternetPlan{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Internet plan migrated")
}
