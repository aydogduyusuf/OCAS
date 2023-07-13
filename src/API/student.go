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

func (server *Server) CreateStudent(ctx *gin.Context) {
	var req CreateStudentRequest
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

	university, err := server.store.GetUniversityByName(ctx, req.UniversityName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
		ctx.Abort()
		return
	}

	arg := DB.CreateStudentParams{
		UserName:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		UID:            university.UID,
		Credit:         0,
	}

	student, err := server.store.CreateStudent(ctx, arg)
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

	rsp := CreateStudentResponse{
		Username:       student.UserName,
		Fullname:       student.FullName,
		Email:          student.Email,
		UniversityName: req.UniversityName,
		Credit:         int(student.Credit),
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) StudentLogin(ctx *gin.Context) {
	var req LoginStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	student, err := server.store.GetStudentByUserName(ctx, req.Username)
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

	err = Helpers.CheckPassword(req.Password, student.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	accessToken, err := generateJWT(student.SID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	university, err := server.store.GetUniversity(ctx, student.UID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
		ctx.Abort()
		return
	}
	rsp := LoginStudentResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Now().Add(server.config.AccessTokenDuration),
		Student: CreateStudentResponse{
			Username:       student.UserName,
			Fullname:       student.FullName,
			Email:          student.Email,
			Description:    student.Description.String,
			Avatar:         student.Description.String,
			UniversityName: university.UniversityName,
			Credit:         int(student.Credit),
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) EditStudentProfilePage(ctx *gin.Context) {
	var req EditStudentProfilePageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	_, err = server.store.GetStudent(ctx, sID)
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

	arg := DB.UpdateStudentProfileParams{
		SID:      sID,
		FullName: req.Name,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		Avatar: sql.NullString{
			String: req.Avatar,
			Valid:  true,
		},
	}
	server.store.UpdateStudentProfile(ctx, arg)

}

func (server *Server) PasswordResetRequest(ctx *gin.Context) {
	var req PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	student, err := server.store.GetStudent(ctx, sID)
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

	if student.Email == req.Email {
		Helpers.SendMail(student.Email, "PASSWORD RESET", "hash", "")
	}
}

func (server *Server) PasswordResetVerify(ctx *gin.Context) {
	var req PasswordResetVerifyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
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

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	arg := DB.UpdateStudentPasswordParams{
		SID:            sID,
		HashedPassword: hashedPassword,
	}
	server.store.UpdateStudentPassword(ctx, arg)
}

func (server *Server) VerifyEmailController(ctx *gin.Context) {
	var req VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	// TO-DO: add email status to student database table
}

func (server *Server) StudentAddCourse(ctx *gin.Context) {
	var req StudentAddCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}
	_, err = server.store.GetStudent(ctx, sID)
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

	cID, err := strconv.ParseInt(req.CourseId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("context error")))
		ctx.Abort()
		return
	}

	server.store.CreateStudentCourses(ctx, DB.CreateStudentCoursesParams{
		SID: sID,
		CID: cID,
	})
}

func (server *Server) EvaluateMentor(ctx *gin.Context) {
	var req EvaluateMentorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}
	_, err = server.store.GetStudent(ctx, sID)
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

	cID, err := strconv.ParseInt(req.CID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("context error")))
		ctx.Abort()
		return
	}

	mentor, err := server.store.GetMentorByCourseID(ctx, cID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't get mentor")))
		ctx.Abort()
		return
	}
	arg := DB.UpdateScoreOfMentorParams{
		MID:   mentor.MID,
		Score: int32(req.Score),
	}
	server.store.UpdateScoreOfMentor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't update score")))
		ctx.Abort()
		return
	}
}

func (server *Server) StudentSubscribe(ctx *gin.Context) {
	var req StudentSubscriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}
	_, err = server.store.GetSubscriptionBySID(ctx, sID)
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("student already has a subscription")))
		ctx.Abort()
		return
	}

	expTime, err := strconv.Atoi(req.SubExpireTime)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	currTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't parse time")))
		ctx.Abort()
		return
	}
	arg := DB.CreateSubscriptionParams{
		SID:           sID,
		PlanType:      req.PlanType,
		SubExpireTime: currTime.AddDate(0, expTime, 0).Format("2006-01-02 15:04:05"),
		SubStartTime:  timeStr,
	}

	subscription, err := server.store.CreateSubscription(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't subscribe student")))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}

func (server *Server) StudentUpdateSubscribe(ctx *gin.Context) {
	var req StudentSubscriptionTypeChangeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	arg := DB.UpdateSubscriptionTypeParams{
		SID:      sID,
		PlanType: req.PlanType,
	}

	subscription, err := server.store.UpdateSubscriptionType(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't update subscription type of the student")))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}

func (server *Server) StudentUpdateSubscribeTime(ctx *gin.Context) {
	var req StudentSubscriptionExtendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	extensionTime, err := strconv.Atoi(req.ExtensionTime)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}

	subscription, err := server.store.GetSubscriptionBySID(ctx, sID)
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
	theTime, err := time.Parse("2006-01-02 15:04:05", subscription.SubExpireTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't parse time")))
		ctx.Abort()
		return
	}

	arg := DB.UpdateSubscriptionTimeParams{
		SID:           sID,
		SubExpireTime: theTime.AddDate(0, extensionTime, 0).Format("2006-01-02 15:04:05"),
	}

	subscription, err = server.store.UpdateSubscriptionTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't update subscription expire time of the student")))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}

func (server *Server) StudentRemainingSubscribeTime(ctx *gin.Context) {
	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	subscription, err := server.store.GetSubscriptionBySID(ctx, sID)
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

	endTime, err := time.Parse("2006-01-02 15:04:05", subscription.SubExpireTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't parse time")))
		ctx.Abort()
		return
	}

	remainingTime := time.Until(endTime)

	ctx.JSON(http.StatusOK, SubscriptionResponse{
		PlanType:      subscription.PlanType,
		RemainingTime: int(remainingTime.Hours() / 24),
	})
}

func (server *Server) StudentCancelSubscribe(ctx *gin.Context) {
	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	err = server.store.DeleteSubscriptionOfAStudent(ctx, sID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("couldn't cancel subscription")))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
