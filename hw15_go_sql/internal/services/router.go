package services

import (
	"net/http"

	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
)

func Router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/shop/getOrdersByUser":
		GetOrdersByUser(w, r)
	case "/api/shop/getUsersAndProducts":
		GetPageableEntities(w, r, UsersAndProducts)
	case "/api/shop/getUserStatistics":
		GetUserStatistics(w, r)

	case "/api/users/getAll":
		GetPageableEntities(w, r, Users)
	case "/api/users/getByID":
		GetUserByID(w, r)
	case "/api/users/update":
		UpdateUser(w, r)
	case "/api/users/create":
		CreateUser(w, r)
	case "/api/users/deleteByID":
		DeleteUser(w, r)
	default:
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "Method is not supported"}`))
	}
}
