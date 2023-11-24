package pfview

import (
	"arknet/modules/common/model"
	"arknet/modules/pf/controller"
	"arknet/modules/pf/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// PF rule
func GetPfrulebasics(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := pfcontroller.GetPfrulebasics(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	}
	return gin.HandlerFunc(fn)
}

func CreatePfrulebasic(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var pfrb pfmodel.Pfrulebasic
		if err := c.BindJSON(&pfrb); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := pfcontroller.CreatePfrulebasic(db, &pfrb)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, pfrb)
			log.Println(pfrb.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func GetRuleByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := pfcontroller.GetPfrulebasicByID(db, id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, res)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

func UpdateRuleByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var pfrb pfmodel.Pfrulebasic
		if err := c.BindJSON(&pfrb); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		pfrb.ID = uint(id)
		err = pfcontroller.UpdatePfrulebasic(db, &pfrb)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, pfrb)
			log.Println(pfrb.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func DeleteRuleByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := pfcontroller.GetPfrulebasicByID(db, id)
		if err == nil {
			err := pfcontroller.DeletePfrulebasic(db, res)
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

// PF Queue
func GetPfqueues(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := pfcontroller.GetPfqueues(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	}
	return gin.HandlerFunc(fn)
}

func CreatePfqueue(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var pfq pfmodel.Pfqueue
		if err := c.BindJSON(&pfq); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := pfcontroller.CreatePfqueue(db, &pfq)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, pfq)
			log.Println(pfq.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func GetQueueByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := pfcontroller.GetPfqueueByID(db, id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, res)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
	}
	return gin.HandlerFunc(fn)
}

func UpdateQueueByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		var pfq pfmodel.Pfqueue
		if err := c.BindJSON(&pfq); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		pfq.ID = uint(id)
		err = pfcontroller.UpdatePfqueue(db, &pfq)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, pfq)
			log.Println(pfq.Genline())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"serv_error": err})
	}
	return gin.HandlerFunc(fn)
}

func DeleteQueueByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		res, err := pfcontroller.GetPfqueueByID(db, id)
		if err == nil {
			err := pfcontroller.DeletePfqueue(db, res)
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
	r.GET("/api/v1/pf/rules", common.BasicAuth, GetPfrulebasics(db))
	r.GET("/api/v1/pf/rules/:id", common.BasicAuth, GetRuleByID(db))
	r.PUT("/api/v1/pf/rules/:id", common.BasicAuth, UpdateRuleByID(db))
	r.DELETE("/api/v1/pf/rules/:id", common.BasicAuth, DeleteRuleByID(db))
	r.POST("/api/v1/pf/rule/new", common.BasicAuth, CreatePfrulebasic(db))

	r.GET("/api/v1/pf/queues", common.BasicAuth, GetPfqueues(db))
	r.GET("/api/v1/pf/queues/:id", common.BasicAuth, GetQueueByID(db))
	r.PUT("/api/v1/pf/queues/:id", common.BasicAuth, UpdateQueueByID(db))
	r.DELETE("/api/v1/pf/queues/:id", common.BasicAuth, DeleteQueueByID(db))
	r.POST("/api/v1/pf/queue/new", common.BasicAuth, CreatePfqueue(db))
}
