package pfcontroller

import (
	"arknet/modules/pf/model"
	"gorm.io/gorm"
)

// PF rule

func CreatePfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
	result := db.Create(pfrb)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPfrulebasicByID(db *gorm.DB, pfrbID int) (*pfmodel.Pfrulebasic, error) {
	var pfrb pfmodel.Pfrulebasic
	result := db.First(&pfrb, pfrbID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pfrb, nil
}

func GetIplanByID(db *gorm.DB, iplanID int) (*pfmodel.InternetPlan, error) {
  var iplan pfmodel.InternetPlan
  result := db.First(&iplan, iplanID)
  if result.Error != nil {
    return nil, result.Error
  }
  return &iplan, nil
}

func UpdatePfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
	result := db.Save(pfrb)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateIplan(db *gorm.DB, iplan *pfmodel.InternetPlan) error {
  result := db.Save(iplan)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func DeletePfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
	result := db.Delete(pfrb)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteInternetPlan(db *gorm.DB, iplan *pfmodel.InternetPlan) error {
  result := db.Delete(iplan)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func GetPfrulebasics(db *gorm.DB) ([]pfmodel.Pfrulebasic, error) {
	var pfrbs []pfmodel.Pfrulebasic
	result := db.Find(&pfrbs)
	if result.Error != nil {
		return nil, result.Error
	}
	return pfrbs, nil
}

//PF Queue

func CreatePfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
	result := db.Create(pfq)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateIplan(db *gorm.DB, iplan *pfmodel.InternetPlan) error {
	result := db.Create(iplan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPfqueueByID(db *gorm.DB, pfqID int) (*pfmodel.Pfqueue, error) {
	var pfq pfmodel.Pfqueue
	result := db.First(&pfq, pfqID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pfq, nil
}

func UpdatePfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
	result := db.Save(pfq)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeletePfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
	result := db.Delete(pfq)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPfqueues(db *gorm.DB) ([]pfmodel.Pfqueue, error) {
	var pfqs []pfmodel.Pfqueue
	result := db.Find(&pfqs)
	if result.Error != nil {
		return nil, result.Error
	}
	return pfqs, nil
}

func GetIplans(db *gorm.DB) ([]pfmodel.InternetPlan, error) {
	var iplans []pfmodel.InternetPlan
	result := db.Find(&iplans)
	if result.Error != nil {
		return nil, result.Error
	}
	return iplans, nil
}
