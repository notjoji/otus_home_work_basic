package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/notjoji/otus_home_work_basic/hw16_docker/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw16_docker/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw16_docker/pkg/pgdb"
	"net/http"
)

type MethodAPI int

const (
	Shop     MethodAPI = 1
	Users    MethodAPI = 2
	Orders   MethodAPI = 3
	Products MethodAPI = 4
)

func (method MethodAPI) GetQueryParameter() string {
	switch method {
	case Shop, Users:
		return "userID"
	case Orders:
		return "orderID"
	case Products:
		return "productID"
	}
	return ""
}

func Router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/shop/getOrdersByUser":
		GetEntityByID(w, r, Shop)
	case "/api/shop/getUsersAndProducts":
		GetPageableEntities(w, r, Shop)
	case "/api/shop/getUserStatistics":
		GetUserStatistics(w, r)

	case "/api/users/getAll":
		GetPageableEntities(w, r, Users)
	case "/api/users/getByID":
		GetEntityByID(w, r, Users)
	case "/api/users/update":
		UpdateEntity(w, r, Users)
	case "/api/users/create":
		CreateEntity(w, r, Users)
	case "/api/users/deleteByID":
		DeleteEntity(w, r, Users)

	case "/api/orders/getAll":
		GetPageableEntities(w, r, Orders)
	case "/api/orders/getByID":
		GetEntityByID(w, r, Orders)
	case "/api/orders/update":
		UpdateEntity(w, r, Orders)
	case "/api/orders/create":
		CreateEntity(w, r, Orders)
	case "/api/orders/deleteByID":
		DeleteEntity(w, r, Orders)

	case "/api/products/getAll":
		GetPageableEntities(w, r, Products)
	case "/api/products/getByID":
		GetEntityByID(w, r, Products)
	case "/api/products/update":
		UpdateEntity(w, r, Products)
	case "/api/products/create":
		CreateEntity(w, r, Products)
	case "/api/products/deleteByID":
		DeleteEntity(w, r, Products)

	default:
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "Method is not supported"}`))
	}
}

func GetUserStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	ctx := context.Background()

	repo := repository.New(pgdb.DB.Conn())

	result, err := repo.GetUserStatistics(ctx)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	utils.ResponseJSON(w, resultByte)
}
