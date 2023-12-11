package oscontroller

import (
  "arknet/modules/os/model"
  "gorm.io/gorm"
)

// IP address
func CreateIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Create(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetIpaddressByID(db *gorm.DB, ipaID int) (*osmodel.Ipaddress, error) {
	var ipa osmodel.Ipaddress
	result := db.First(&ipa, ipaID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ipa, nil
}

func UpdateIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Save(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteIpaddress(db *gorm.DB, ipa *osmodel.Ipaddress) error {
	result := db.Delete(ipa)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetIPaddresses(db *gorm.DB) ([]osmodel.Ipaddress, error) {
  var addrs []osmodel.Ipaddress
  result := db.Find(&addrs)
  if result.Error != nil {
    return nil, result.Error
  }
  return addrs, nil
}

// Key value pairs
func CreateKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Create(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetKeyvalByID(db *gorm.DB, kvID int) (*osmodel.Keyval, error) {
	var kv osmodel.Keyval
	result := db.First(&kv, kvID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &kv, nil
}

func UpdateKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Save(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteKeyval(db *gorm.DB, kv *osmodel.Keyval) error {
	result := db.Delete(kv)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetIKeyVals(db *gorm.DB) ([]osmodel.Keyval, error) {
  var kvs []osmodel.Keyval
  result := db.Find(&kvs)
  if result.Error != nil {
    return nil, result.Error
  }
  return kvs, nil
}

// Interface
func CreateInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Create(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetInterfaceByID(db *gorm.DB, ifaceID int) (*osmodel.Interface, error) {
	var iface osmodel.Interface
	result := db.First(&iface, ifaceID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &iface, nil
}

func UpdateInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Save(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteInterface(db *gorm.DB, iface *osmodel.Interface) error {
	result := db.Delete(iface)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetInterfaces(db *gorm.DB) ([]osmodel.Interface, error) {
  var ifaces []osmodel.Interface
  result := db.Find(&ifaces)
  if result.Error != nil {
    return nil, result.Error
  }
  return ifaces, nil
}
