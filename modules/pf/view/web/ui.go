package pfui

import (
	"arknet/modules/common/model"
	"arknet/modules/pf/controller"
	"arknet/modules/pf/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
  "strconv"
)

func uiGetPfqueues(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, _ := pfcontroller.GetIplans(db)
		c.HTML(
			http.StatusOK,
			"a_plans.html",
			gin.H{
				"title": "Internet Plans",
				"plans": res,
			},
		)

	}
	return gin.HandlerFunc(fn)
}

func uiCreatePfqueue(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var iplan pfmodel.InternetPlan
		if c.Request.Method == "POST" {
			if err := c.Bind(&iplan); err != nil {
				c.HTML(
					http.StatusOK,
					"a_plan.html",
					gin.H{
						"title": "Internet Plan",
						"flash": err,
					},
				)
				return
			}
      var err error
      id, err := strconv.Atoi(c.Request.PostFormValue("id"))
      if err == nil {
        iplan.ID = uint(id)
        if c.Request.PostFormValue("delete") == "yes" {
          err = pfcontroller.DeleteInternetPlan(db, &iplan)
        } else {
          err = pfcontroller.UpdateIplan(db, &iplan)
        }
      } else {
			  err = pfcontroller.CreateIplan(db, &iplan)
      }
			if err == nil {
				c.Redirect(http.StatusMovedPermanently, "/plans")
				return
			} else {
				c.HTML(
					http.StatusOK,
					"a_plan.html",
					gin.H{
						"title": "Internet Plan",
						"flash": err,
					},
				)
				return
			}
		} else {
      id, err := strconv.Atoi(c.Query("id"))
      if err == nil {
        iplan, err := pfcontroller.GetIplanByID(db, id)
        if err != nil {
          c.HTML(
            http.StatusOK,
            "a_plan.html",
            gin.H{
              "title": "Internet Plan",
              "flash": err,
            },
          )
          return
        } else {
          del := "no"
          if c.Query("delete") == "yes" {
            del = "yes"
          }
          c.HTML(
            http.StatusOK,
            "a_plan.html",
            gin.H{
              "title": "Internet Plan",
              "flash": " ",
              "data": iplan,
              "delete": del,
            },
          )
        }
      } else {
        c.HTML(
          http.StatusOK,
          "a_plan.html",
          gin.H{
            "title": "Internet Plan",
            "flash": " ",
          },
        )
      }
		}
	}

	return gin.HandlerFunc(fn)
}

func Setroutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/plans", common.BasicAuth, uiGetPfqueues(db))
	r.GET("/plan", common.BasicAuth, uiCreatePfqueue(db))
	r.POST("/plan", common.BasicAuth, uiCreatePfqueue(db))
}
