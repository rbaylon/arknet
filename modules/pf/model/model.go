package pfmodel

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Pfrulebasic struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Action      string `json:"action"`
	Direction   string `json:"direction"`
	Interface   string `json:"interface"`
	Addrfamily  string `json:"addrfamily"`
	Protocol    string `json:"protocol"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Dstport     string `json:"dstport"`
	Custom      string `json:"custom"`
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
	Name          string `json:"name"`
	Onorparent    string `json:"onorparent"`
	Parentoriface string `json:"parentoriface"`
	Bandwidth     string `json:"bandwidth"`
	Custom        string `json:"custom"`
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
}
