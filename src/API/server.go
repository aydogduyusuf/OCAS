package API

import (
	"github.com/gin-gonic/gin"
	db "group33/ocas/src/DB/sqlc"
	"group33/ocas/src/Helpers"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	config Helpers.Config
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config Helpers.Config, store db.Store) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	authRoutes := router.Group("/").Use(IsAuthorized(server.config))

	router.POST("/student/sign", server.CreateStudent)
	router.POST("/student/login", server.StudentLogin)

	router.POST("/mentor/sign", server.MentorCreate)
	router.POST("/mentor/login", server.MentorLogin)
	router.GET("/mentor/getScore", server.MentorGetScore)
	authRoutes.POST("/mentor/course/event", server.MentorCreateEvent)
	authRoutes.POST("/mentor/course/eventUpdate", server.MentorUpdateEvent)
	authRoutes.DELETE("/mentor/course/eventDelete", server.MentorDeleteEvent)

	router.POST("/university", server.CreateUniversity)
	router.GET("/allschools", server.GetAllUniversities)

	router.POST("/course", server.CreateCourse)
	router.GET("/allcourses", server.GetAllCourses)
	router.GET("/course/events", server.GetCourseEvents)

	authRoutes.GET("/student/courses", server.ReadCourseStudent)

	authRoutes.POST("/student/passwordResetRequest", server.PasswordResetRequest)
	authRoutes.POST("/student/passwordResetVerify", server.PasswordResetVerify)
	authRoutes.POST("/student/verifyEmail", server.VerifyEmailController)
	authRoutes.POST("/student/editProfile", server.EditStudentProfilePage)
	authRoutes.POST("/student/addCourse", server.StudentAddCourse)
	authRoutes.POST("/student/evaluateMentor", server.EvaluateMentor)

	authRoutes.POST("/student/subscribe", server.StudentSubscribe)
	authRoutes.POST("/student/subscribe/typeChange", server.StudentUpdateSubscribe)
	authRoutes.POST("/student/subscribe/extend", server.StudentUpdateSubscribeTime)
	authRoutes.GET("/student/subscription", server.StudentRemainingSubscribeTime)
	authRoutes.DELETE("/student/subscribe/cancel", server.StudentCancelSubscribe)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
