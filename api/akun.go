package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/terajari/idompet/database/sqlc"
)

type createAkunReq struct {
	Nama  string `json:"nama" binding:"required"`
	Saldo int64  `json:"saldo" binding:"required"`
}

func (server *Server) createAkun(ctx *gin.Context) {
	var req createAkunReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	akun, err := server.Store.CreateAkun(ctx, database.CreateAkunParams{
		Nama:  req.Nama,
		Saldo: req.Saldo,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)
}

type getAkunReq struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAkun(ctx *gin.Context) {
	var req getAkunReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	akun, err := server.Store.GetAkun(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)
}

type listAkunReq struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAkun(ctx *gin.Context) {
	var req listAkunReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}
	akun, err := server.Store.ListAkun(ctx, database.ListAkunParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)
}
