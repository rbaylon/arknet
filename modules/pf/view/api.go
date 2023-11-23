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

func Setroutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/v1/pf/rules", common.BasicAuth, GetPfrulebasics(db))
	r.GET("/api/v1/pf/rules/:id", common.BasicAuth, GetRuleByID(db))
	r.POST("/api/v1/pf/rule/new", common.BasicAuth, CreatePfrulebasic(db))
}
