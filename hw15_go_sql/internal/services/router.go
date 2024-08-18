package services

import (
	"net/http"

	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	switch r.URL.Path {
	case "/api/getOrdersByUser":
		GetOrdersByUser(w, r)
	case "/api/getUsersAndProducts":
		GetUsersAndProducts(w, r)
	case "/api/getUserStatistics":
		GetUserStatistics(w, r)
	default:
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "Method is not supported"}`))
	}
}
