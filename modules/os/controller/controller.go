package oscontroller

import (
	"arknet/modules/os/model"
)

// IP address
func createIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Create(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getIpaddressByID(db *gorm.DB, ipaID uint) (*osmodel.Ipaddress, error) {
	var ipa osmodel.Ipaddress
	result := db.First(&ipa, ipaID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ipa, nil
}

func updateIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Save(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Delete(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Key value pairs
func createKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Create(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getKeyvalByID(db *gorm.DB, kvID uint) (*osmodel.Keyval, error) {
	var kv osmodel.Keyval
	result := db.First(&kv, kvID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &kv, nil
}

func updateKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Save(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Delete(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Interface
func createInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Create(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getInterfaceByID(db *gorm.DB, ifaceID uint) (*osmodel.Interface, error) {
	var iface osmodel.Interface
	result := db.First(&iface, ifaceID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &iface, nil
}

func updateInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Save(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Delete(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
