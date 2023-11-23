package pfcontroller

import (
  "github.com/rbaylon/arknet/modules/pf/model"
)

// PF rule

func createPfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
    result := db.Create(pfrb)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func getPfrulebasicByID(db *gorm.DB, pfrbID uint) (*pfmodel.Pfrulebasic, error) {
    var pfrb pfmodel.Pfrulebasic
    result := db.First(&pfrb, pfrbID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &pfrb, nil
}

func updatePfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
    result := db.Save(pfrb)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func deletePfrulebasic(db *gorm.DB, pfrb *pfmodel.Pfrulebasic) error {
    result := db.Delete(pfrb)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

//PF Queue

func createPfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
    result := db.Create(pfq)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func getPfqueueByID(db *gorm.DB, pfqID uint) (*pfmodel.Pfqueue, error) {
    var pfq pfmodel.Pfqueue
    result := db.First(&pfq, pfqID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &pfq, nil
}

func updatePfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
    result := db.Save(pfq)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func deletePfqueue(db *gorm.DB, pfq *pfmodel.Pfqueue) error {
    result := db.Delete(pfq)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

