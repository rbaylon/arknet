package pfui

import (
	"arknet/modules/common/model"
	"arknet/modules/pf/controller"
  "arknet/modules/pf/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func uiGetPfqueues(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, _ := pfcontroller.GetPfqueues(db)
    c.HTML(
      http.StatusOK,
      "a_plans.html",
      gin.H{
        "title": "Subscription Plans",
        "plans": res,
      },
    )

	}
	return gin.HandlerFunc(fn)
}

func uiCreatePfqueue(db *gorm.DB) gin.HandlerFunc {
  fn := func(c *gin.Context) {
    var pfq pfmodel.Pfqueue
    if c.Request.Method == "POST" {
      if err := c.Bind(&pfq); err != nil {
        c.HTML(
          http.StatusOK,
          "a_plan.html",
          gin.H{
            "title": "Subscriber Plan",
            "flash": "Input error",
          },
        )
        return
      }
      err := pfcontroller.CreatePfqueue(db, &pfq)
      if err == nil {
        c.Redirect(http.StatusMovedPermanently, "/plans")
        return
      } else {
        c.HTML(
          http.StatusOK,
          "a_plan.html",
          gin.H{
            "title": "Subscriber Plan",
            "flash": err,
          },
        )
        return
      }
    } else {
      c.HTML(
        http.StatusOK,
        "a_plan.html",
        gin.H{
          "title": "Subscriber Plan",
          "flash": " ",
        },
      )
    }
  }

  return gin.HandlerFunc(fn)
}

func Setroutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/plans", common.BasicAuth, uiGetPfqueues(db))
	r.GET("/plan", common.BasicAuth, uiCreatePfqueue(db))
	r.POST("/plan", common.BasicAuth, uiCreatePfqueue(db))
}
