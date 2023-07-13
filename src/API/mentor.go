package API

import (
	"database/sql"
	"errors"
	DB "group33/ocas/src/DB/sqlc"
	"group33/ocas/src/Helpers"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) MentorCreate(ctx *gin.Context) {
	var req CreateMentorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give same password")))
		ctx.Abort()
		return
	}
	passwordIsValid := Helpers.VerifyPassword(req.Password)
	if !passwordIsValid {
		ctx.JSON(400, errorResponse(errors.New("password must be at least 8 characters long and contain at least one uppercase letter, lowercase letter, digit and special character")))
		return
	}
	hashedPassword, err := Helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	course, err := server.store.GetCourseByName(ctx, req.CourseName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct course name")))
		ctx.Abort()
		return
	}

	university, err := server.store.GetUniversityByName(ctx, req.UniversityName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
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
	arg := DB.CreateMentorParams{
		Username:        req.Username,
		CID:             course.CID,
		FullName:        req.FullName,
		HashedPassword:  hashedPassword,
		Email:           req.Email,
		Description:     req.Description,
		EvaluationCount: 0,
		Score:           0,
		Balance:         0,
		UID:             university.UID,
		Country:         address.Country,
		City:            address.City,
		Street:          address.Street,
	}

	mentor, err := server.store.CreateMentor(ctx, arg)
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

	rsp := CreateMentorResponse{
		Username:       mentor.Username,
		Fullname:       mentor.FullName,
		Email:          mentor.Email,
		UniversityName: university.UniversityName,
		Description:    mentor.Description,
		Score:          int(mentor.Score),
		Balance:        int(mentor.Balance),
		Address: CreateAddressesRequest{
			Country: mentor.Country,
			City:    mentor.City,
			Street:  mentor.Street,
		},
		CourseName: course.CourseName,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) MentorLogin(ctx *gin.Context) {
	var req LoginStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	mentor, err := server.store.GetMentorByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	err = Helpers.CheckPassword(req.Password, mentor.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	accessToken, err := generateJWTMentor(mentor.MID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	university, err := server.store.GetUniversity(ctx, mentor.UID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
		ctx.Abort()
		return
	}
	course, err := server.store.GetCourse(ctx, mentor.CID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
		ctx.Abort()
		return
	}
	rsp := LoginMentorResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Now().Add(server.config.AccessTokenDuration),
		Mentor: CreateMentorResponse{
			Username:       mentor.Username,
			Fullname:       mentor.FullName,
			Email:          mentor.Email,
			UniversityName: university.UniversityName,
			Description:    mentor.Description,
			Score:          int(mentor.Score),
			Balance:        int(mentor.Balance),
			Address: CreateAddressesRequest{
				Country: mentor.Country,
				City:    mentor.City,
				Street:  mentor.Street,
			},
			CourseName: course.CourseName,
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) MentorGetScore(ctx *gin.Context) {
	courseID := ctx.Query("id")
	cID, err := strconv.ParseInt(courseID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}
	mentor, err := server.store.GetMentorByCourseID(ctx, cID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("wrong course id")))
		ctx.Abort()
		return
	}
	if mentor.EvaluationCount == 0 {
		ctx.JSON(http.StatusOK, 0)
	} else {
		ctx.JSON(http.StatusOK, float32(mentor.Score)/float32(mentor.EvaluationCount))
	}

}

func (server *Server) MentorCreateEvent(ctx *gin.Context) {
	var req CreateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	mID, err := Helpers.GetMentorFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	mentor, err := server.store.GetMentor(ctx, mID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("can't get mentor")))
		ctx.Abort()
		return
	}

	arg := DB.CreateEventParams{
		CID:       mentor.CID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Title:     req.Title,
		Color:     req.Color,
	}
	event, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func (server *Server) MentorUpdateEvent(ctx *gin.Context) {
	eventID := ctx.Query("id")
	startTime := ctx.Query("start_time")
	endTime := ctx.Query("end_time")
	title := ctx.Query("title")
	color := ctx.Query("color")

	eID, err := strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}
	_, err = Helpers.GetMentorFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	event, err := server.store.GetEventByEventID(ctx, eID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("can't get event")))
		ctx.Abort()
		return
	}

	if startTime != "" {
		event, err = server.store.UpdateEventStartTime(ctx, DB.UpdateEventStartTimeParams{
			EID:       eID,
			StartTime: startTime,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("can't update event start time")))
			ctx.Abort()
			return
		}
	}
	if endTime != "" {
		event, err = server.store.UpdateEventEndTime(ctx, DB.UpdateEventEndTimeParams{
			EID:     eID,
			EndTime: endTime,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("can't update event end time")))
			ctx.Abort()
			return
		}
	}
	if title != "" {
		event, err = server.store.UpdateEventTitle(ctx, DB.UpdateEventTitleParams{
			EID:   eID,
			Title: title,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("can't update event title")))
			ctx.Abort()
			return
		}
	}
	if color != "" {
		event, err = server.store.UpdateEventColor(ctx, DB.UpdateEventColorParams{
			EID:   eID,
			Color: color,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("can't update event color")))
			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, event)
}

func (server *Server) MentorDeleteEvent(ctx *gin.Context) {
	_, err := Helpers.GetMentorFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	eventID := ctx.Query("id")
	eID, err := strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}

	err = server.store.DeleteEventOfACourse(ctx, eID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't delete event")))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
