package API

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	DB "group33/ocas/src/DB/sqlc"
	"net/http"
)

func (server *Server) CreateUniversity(ctx *gin.Context) {
	var req CreateUniversityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	addressGetArg := DB.GetAddressParams{
		Country: req.Address.Country,
		City:    req.Address.City,
		Street:  req.Address.Street,
	}

	address, err := server.store.GetAddress(ctx, addressGetArg)
	if err != nil {
		addressArg := DB.CreateAddressesParams{
			Country: req.Address.Country,
			City:    req.Address.City,
			Street:  req.Address.Street,
		}
		address, err = server.store.CreateAddresses(ctx, addressArg)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					ctx.JSON(http.StatusForbidden, errorResponse(err))
					return
				}
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	arg := DB.CreateUniversityParams{
		UniversityName: req.UniversityName,
		Abbreviation:   req.Abbreviation,
		EmailExtension: req.EmailExtension,
		Country:        address.Country,
		City:           address.City,
		Street:         address.Street,
	}

	university, err := server.store.CreateUniversity(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				ctx.Abort()
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, university)
}

func (server *Server) GetAllUniversities(ctx *gin.Context) {
	universities, err := server.store.GetAllUniversities(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, universities)
}
