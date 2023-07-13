package API

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	DB "group33/ocas/src/DB/sqlc"
	"group33/ocas/src/Helpers"
	"net/http"
	"strconv"
)

func (server *Server) CreateCourse(ctx *gin.Context) {
	var req CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	//addressGetArg := server.store.GetCourse(ctx, )
	university, err := server.store.GetUniversityByName(ctx, req.UniversityName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("please give correct university name")))
		ctx.Abort()
		return
	}

	arg := DB.CreateCourseParams{
		CourseName:  req.CourseName,
		Semester:    req.Semester,
		UID:         university.UID,
		Description: req.Description,
	}

	course, err := server.store.CreateCourse(ctx, arg)
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

	ctx.JSON(http.StatusOK, course)
}

func (server *Server) GetAllCourses(ctx *gin.Context) {
	courses, err := server.store.GetAllCourses(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

func (server *Server) ReadCourseStudent(ctx *gin.Context) {
	sID, err := Helpers.GetStudentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("you need to login to see the courses")))
		return
	}

	studentCourses, err := server.store.GetCoursesOfAStudent(ctx, sID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	var res []ReadCourseStudentResponse
	for i := 0; i < len(studentCourses); i++ {
		course, err := server.store.GetCourse(ctx, studentCourses[i].CID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			ctx.Abort()
			return
		}
		/*university, err := server.store.GetUniversity(ctx, course.UID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			ctx.Abort()
			return
		}*/
		c := ReadCourseStudentResponse{
			CourseName:  course.CourseName,
			Semester:    course.Semester,
			Description: course.Description,
		}
		res = append(res, c)
	}
	ctx.JSON(http.StatusOK, res)
	return
}

func (server *Server) GetCourseEvents(ctx *gin.Context) {
	cIDstr := ctx.Query("id")
	cID, err := strconv.ParseInt(cIDstr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("type cast error")))
		ctx.Abort()
		return
	}
	events, err := server.store.GetEventsByCourseID(ctx, cID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, events)
}
