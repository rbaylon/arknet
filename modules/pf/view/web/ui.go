package pfui

import (
	"arknet/modules/common/model"
	"arknet/modules/pf/controller"
	"arknet/modules/pf/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
  "strconv"
  "github.com/gin-contrib/sessions"
)

func uiGetIplans(db *gorm.DB) gin.HandlerFunc {
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

func uiCreateIplan(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var iplan pfmodel.InternetPlan
    csrftoken, _ := common.GenCsrfToken()
    session := sessions.Default(c)
		if c.Request.Method == "POST" {
      t, _ := session.Get("csrftoken").(string)
			if err := c.Bind(&iplan); err != nil || common.CheckCsrfToken(t, c.Request.PostFormValue("csrftoken")) == false {
        session.Set("csrftoken", csrftoken)
        session.Save()
				c.HTML(
          400,
					"a_plan.html",
					gin.H{
						"title": "Internet Plan",
						"flash": err,
            "CsrfToken": csrftoken,
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
        session.Set("csrftoken", csrftoken)
        session.Save()
				c.HTML(
					400,
					"a_plan.html",
					gin.H{
						"title": "Internet Plan",
						"flash": err,
            "CsrfToken": csrftoken,
					},
				)
				return
			}
		} else {
      session.Set("csrftoken", csrftoken)
      session.Save()
      id, err := strconv.Atoi(c.Query("id"))
      if err == nil {
        iplan, err := pfcontroller.GetIplanByID(db, id)
        if err != nil {
          c.HTML(
            400,
            "a_plan.html",
            gin.H{
              "title": "Internet Plan",
              "flash": err,
              "CsrfToken": csrftoken,
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
              "CsrfToken": csrftoken,
            },
          )
          return
        }
      } else {
        c.HTML(
          http.StatusOK,
          "a_plan.html",
          gin.H{
            "title": "Internet Plan",
            "flash": " ",
            "CsrfToken": csrftoken,
          },
        )
      }
		}
	}

	return gin.HandlerFunc(fn)
}

func uiCreateTelcoConfig(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var tc pfmodel.TelcoConfig
    csrftoken, _ := common.GenCsrfToken()
    session := sessions.Default(c)
    session.Set("csrftoken", csrftoken)
    session.Save()
		if c.Request.Method == "POST" {
      t, _ := session.Get("csrftoken").(string)
			if err := c.Bind(&tc); err != nil || common.CheckCsrfToken(t, c.Request.PostFormValue("csrftoken")) == false {
				c.HTML(
          400,
					"a_telco.html",
					gin.H{
						"title": "Telco Config",
						"flash": err,
            "CsrfToken": csrftoken,
					},
				)
				return
			}
      var err error
      id, err := strconv.Atoi(c.Request.PostFormValue("id"))
      if err == nil {
        tc.ID = uint(id)
        err = pfcontroller.UpdateTelcoConfig(db, &tc)
      } else {
			  err = pfcontroller.CreateTelcoConfig(db, &tc)
      }
			if err == nil {
				c.Redirect(http.StatusMovedPermanently, "/plans")
				return
			} else {
				c.HTML(
					400,
					"a_telco.html",
					gin.H{
						"title": "Internet Plan",
						"flash": err,
            "CsrfToken": csrftoken,
					},
				)
				return
			}
		} else {
      id, err := strconv.Atoi(c.Query("id"))
      if err == nil {
        tc, err := pfcontroller.GetTelcoConfigByID(db, id)
        if err != nil {
          c.HTML(
            400,
            "a_telco.html",
            gin.H{
              "title": "Internet Plan",
              "flash": err,
              "CsrfToken": csrftoken,
            },
          )
          return
        } else {
          c.HTML(
            http.StatusOK,
            "a_telco.html",
            gin.H{
              "title": "Internet Plan",
              "flash": " ",
              "data": tc,
              "CsrfToken": csrftoken,
            },
          )
        }
      } else {
        c.HTML(
          http.StatusOK,
          "a_telco.html",
          gin.H{
            "title": "Internet Plan",
            "flash": " ",
            "CsrfToken": csrftoken,
          },
        )
      }
		}
	}

	return gin.HandlerFunc(fn)
}

func Setroutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/plans", common.BasicAuth, uiGetIplans(db))
	r.GET("/plan", common.BasicAuth, uiCreateIplan(db))
	r.POST("/plan", common.BasicAuth, uiCreateIplan(db))
  r.GET("/telco", common.BasicAuth, uiCreateTelcoConfig(db))
  r.POST("/telco", common.BasicAuth, uiCreateTelcoConfig(db))
}
