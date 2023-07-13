package Helpers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func GetStudentFromContext(ctx *gin.Context) (int64, error) {
	studentId, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("context error")))
		ctx.Abort()
		return 0, errors.New("context error")
	}
	sIDstr, ok := studentId.(string)
	if !ok {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return 0, errors.New("type cast error")
	}
	sID, err := strconv.ParseInt(sIDstr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return 0, errors.New("type cast error")
	}
	return sID, nil
}

func GetMentorFromContext(ctx *gin.Context) (int64, error) {
	mentorId, ok := ctx.Get("mentor")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("context error")))
		ctx.Abort()
		return 0, errors.New("context error")
	}
	mIDstr, ok := mentorId.(string)
	if !ok {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return 0, errors.New("type cast error")
	}
	mID, err := strconv.ParseInt(mIDstr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return 0, errors.New("type cast error")
	}
	return mID, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
