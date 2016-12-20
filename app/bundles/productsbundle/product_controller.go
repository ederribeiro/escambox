package productsbundle

import (
	"net/http"
	"github.com/ederribeiro/escambox/app/common"
)

type ProductController struct {
	common.Controller
}

func (c *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		[]*Product{New("Tênis", "Lindo Tênis")},
	    http.StatusOK,
	)
}