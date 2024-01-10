package osapi

import (
	"arknet/modules/common/model"
	"arknet/modules/os/controller"
	"arknet/modules/os/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// IP Addresses
func GetIPaddresses(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := oscontroller.GetIPaddresses(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	}
	return gin.HandlerFunc(fn)
}

func CreateIpaddress(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var addr osmodel.Ipaddress
		if err := c.BindJSON(&addr); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := oscontroller.CreateIpaddress(db, &addr)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, addr)
			log.Println(addr.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func GetIpaddressByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetIpaddressByID(db, id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, res)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

func UpdateIpaddressByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var addr osmodel.Ipaddress
		if err := c.BindJSON(&addr); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		addr.ID = uint(id)
		err = oscontroller.UpdateIpaddress(db, &addr)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, addr)
			log.Println(addr.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func DeleteIpaddressByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetIpaddressByID(db, id)
		if err == nil {
			err := oscontroller.DeleteIpaddress(db, res)
			if err == nil {
				c.IndentedJSON(http.StatusOK, res)
				return
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "rule not deleted"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

// Keyvals
func GetIKeyVals(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := oscontroller.GetIKeyVals(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	}
	return gin.HandlerFunc(fn)
}

func CreateKeyval(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var kv osmodel.Keyval
		if err := c.BindJSON(&kv); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := oscontroller.CreateKeyval(db, &kv)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, kv)
			log.Println(kv.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func GetKeyvalByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetKeyvalByID(db, id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, res)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

func UpdateKeyvalByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var kv osmodel.Keyval
		if err := c.BindJSON(&kv); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		kv.ID = uint(id)
		err = oscontroller.UpdateKeyval(db, &kv)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, kv)
			log.Println(kv.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func DeleteKeyvalByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetKeyvalByID(db, id)
		if err == nil {
			err := oscontroller.DeleteKeyval(db, res)
			if err == nil {
				c.IndentedJSON(http.StatusOK, res)
				return
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "rule not deleted"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

// Interfaces
func GetInterfaces(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := oscontroller.GetInterfaces(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	}
	return gin.HandlerFunc(fn)
}

func CreateInterface(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var iface osmodel.Interface
		if err := c.BindJSON(&iface); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := oscontroller.CreateInterface(db, &iface)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, iface)
			log.Println(iface.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func GetInterfaceByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetInterfaceByID(db, id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, res)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

func UpdateInterfaceByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var iface osmodel.Interface
		if err := c.BindJSON(&iface); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		iface.ID = uint(id)
		err = oscontroller.UpdateInterface(db, &iface)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, iface)
			log.Println(iface.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func DeleteInterfaceByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := oscontroller.GetInterfaceByID(db, id)
		if err == nil {
			err := oscontroller.DeleteInterface(db, res)
			if err == nil {
				c.IndentedJSON(http.StatusOK, res)
				return
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "rule not deleted"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

// set routes
func Setroutes(r *gin.Engine, db *gorm.DB) {
  // IP Address
	r.GET("/api/v1/os/ips", common.BasicAuth, GetIPaddresses(db))
	r.GET("/api/v1/os/ip/:id", common.BasicAuth, GetIpaddressByID(db))
	r.PUT("/api/v1/os/ip/:id", common.BasicAuth, UpdateIpaddressByID(db))
	r.DELETE("/api/v1/os/ip/:id", common.BasicAuth, DeleteIpaddressByID(db))
	r.POST("/api/v1/os/ip/new", common.BasicAuth, CreateIpaddress(db))

  // keyval
  r.GET("/api/v1/os/kvs", common.BasicAuth, GetIKeyVals(db))
  r.GET("/api/v1/os/kv/:id", common.BasicAuth, GetKeyvalByID(db))
  r.PUT("/api/v1/os/kv/:id", common.BasicAuth, UpdateKeyvalByID(db))
  r.DELETE("/api/v1/os/kv/:id", common.BasicAuth, DeleteKeyvalByID(db))
  r.POST("/api/v1/os/kv/new", common.BasicAuth, CreateKeyval(db))

  // Interfaces
  r.GET("/api/v1/os/interfaces", common.BasicAuth, GetInterfaces(db))
  r.GET("/api/v1/os/interface/:id", common.BasicAuth, GetInterfaceByID(db))
  r.PUT("/api/v1/os/interface/:id", common.BasicAuth, UpdateInterfaceByID(db))
  r.DELETE("/api/v1/os/interface/:id", common.BasicAuth, DeleteInterfaceByID(db))
  r.POST("/api/v1/os/interface/new", common.BasicAuth, CreateInterface(db))
}
