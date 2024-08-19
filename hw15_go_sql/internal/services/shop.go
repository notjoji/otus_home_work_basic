package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

func GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	userID, err := utils.GetNumericQueryParam(r, "userID")
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

	result, err := repo.GetOrdersByUserId(ctx, int64(userID))
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

func GetUsersAndProducts(w http.ResponseWriter, r *http.Request) {
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

	params := repository.GetUsersAndProductsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetUsersAndProducts(ctx, params)
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

func GetUserStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

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
