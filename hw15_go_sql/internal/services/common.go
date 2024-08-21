package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/jackc/pgx/v5"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

type CreateOrderWithProducts struct {
	UserID      *int64             `db:"user_id" json:"userId"`
	OrderDate   pgtype.Timestamptz `db:"order_date" json:"orderDate"`
	TotalAmount int64              `db:"total_amount" json:"totalAmount"`
	ProductIDs  []*int64           `json:"productIds"`
}

func (req CreateOrderWithProducts) Map() repository.CreateOrderParams {
	return repository.CreateOrderParams{
		UserID:      req.UserID,
		OrderDate:   req.OrderDate,
		TotalAmount: req.TotalAmount,
	}
}

func CreateOrderProductParams(orderID *int64, productIDs []*int64) []repository.CreateOrderProductParams {
	result := make([]repository.CreateOrderProductParams, len(productIDs))
	for i, id := range productIDs {
		result[i] = repository.CreateOrderProductParams{
			ProductID: id,
			OrderID:   orderID,
		}
	}
	return result
}

func GetPageableEntities(w http.ResponseWriter, r *http.Request, method MethodAPI) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	limit, offset, err := utils.GetPageableParams(r)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	repo := repository.New(pgdb.DB.Conn())

	var result any
	var methodError error
	switch method {
	case Shop:
		result, methodError = repo.GetUsersAndProducts(ctx, repository.GetUsersAndProductsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case Users:
		result, methodError = repo.GetUsers(ctx, repository.GetUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case Orders:
		result, methodError = repo.GetOrders(ctx, repository.GetOrdersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case Products:
		result, methodError = repo.GetProducts(ctx, repository.GetProductsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	default:
		methodError = errors.New("invalid method")
	}
	if methodError != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, methodError.Error()),
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

func GetEntityByID(w http.ResponseWriter, r *http.Request, method MethodAPI) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	ID, err := utils.GetNumericQueryParam(r, method.GetQueryParameter())
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	repo := repository.New(pgdb.DB.Conn())

	var result any
	var methodError error
	switch method {
	case Shop:
		result, methodError = repo.GetOrdersByUserId(ctx, int64(ID))
	case Users:
		result, methodError = repo.GetUserById(ctx, int64(ID))
	case Orders:
		result, methodError = repo.GetOrderById(ctx, int64(ID))
	case Products:
		result, methodError = repo.GetProductById(ctx, int64(ID))
	default:
		methodError = errors.New("invalid method")
	}
	if methodError != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, methodError.Error()),
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

func UpdateEntity(w http.ResponseWriter, r *http.Request, method MethodAPI) {
	if r.Method != http.MethodPut {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "PUT method is required"}`))
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot read request"}`))
		return
	}

	ctx := context.Background()
	msg := ""

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	switch method {
	case Shop:
		err = errors.New("invalid method")
	case Users:
		{
			var params repository.UpdateUserParams
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			_, err = repo.UpdateUser(ctx, params)
			msg = fmt.Sprintf(`"User updated, id=%d"`, params.ID)
		}
	case Orders:
		{
			var params repository.UpdateOrderParams
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			_, err = repo.UpdateOrder(ctx, params)
			msg = fmt.Sprintf(`"Order updated, id=%d"`, params.ID)
		}
	case Products:
		{
			var params repository.UpdateProductParams
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			_, err = repo.UpdateProduct(ctx, params)
			msg = fmt.Sprintf(`"Product updated, id=%d"`, params.ID)
		}
	default:
		err = errors.New("invalid method")
	}

	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	_ = tx.Commit(ctx)

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "%s"}`, msg),
	))
}

func CreateEntity(w http.ResponseWriter, r *http.Request, method MethodAPI) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "POST method is required"}`))
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot read request"}`))
		return
	}

	ctx := context.Background()
	msg := ""
	var id int64
	var methodError error

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	switch method {
	case Shop:
		methodError = errors.New("invalid method")
	case Users:
		{
			var params repository.CreateUserParams
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			id, methodError = repo.CreateUser(ctx, params)
			msg = fmt.Sprintf(`"User created, id=%d"`, id)
		}
	case Orders:
		{
			var params CreateOrderWithProducts
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			id, methodError = repo.CreateOrder(ctx, params.Map())

			if methodError != nil {
				utils.ResponseJSON(w, []byte(
					fmt.Sprintf(`{"success": false,"msg": "%s"}`, methodError.Error()),
				))
				return
			}

			orderProducts := CreateOrderProductParams(&id, params.ProductIDs)
			for _, orderProduct := range orderProducts {
				_, methodError = repo.CreateOrderProduct(ctx, orderProduct)
				if methodError != nil {
					break
				}
			}

			msg = fmt.Sprintf(`"Order created, id=%d"`, id)
		}
	case Products:
		{
			var params repository.CreateProductParams
			err = json.Unmarshal(request, &params)
			if err != nil {
				utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
				return
			}
			id, methodError = repo.CreateProduct(ctx, params)
			msg = fmt.Sprintf(`"Product created, id=%d"`, id)
		}
	default:
		methodError = errors.New("invalid method")
	}

	if methodError != nil {
		_ = tx.Rollback(ctx)
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, methodError.Error()),
		))
		return
	}
	_ = tx.Commit(ctx)

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "%s"}`, msg),
	))
}

func DeleteEntity(w http.ResponseWriter, r *http.Request, method MethodAPI) {
	if r.Method != http.MethodDelete {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "POST method is required"}`))
		return
	}

	ID, err := utils.GetNumericQueryParam(r, method.GetQueryParameter())
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	switch method {
	case Shop:
		err = errors.New("invalid method")
	case Users:
		_, err = repo.DeleteUser(ctx, int64(ID))
	case Orders:
		_, err = repo.DeleteOrder(ctx, int64(ID))
	case Products:
		_, err = repo.DeleteProduct(ctx, int64(ID))
	default:
		err = errors.New("invalid method")
	}
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	_ = tx.Commit(ctx)

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "Entity deleted, id=%d"}`, ID),
	))
}
