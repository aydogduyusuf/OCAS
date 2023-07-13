package API

import "time"

type CreateStudentRequest struct {
	Username        string `json:"username" binding:"required,alphanum" `
	Password        string `json:"password" binding:"required,min=6" `
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6" `
	FullName        string `json:"full_name" binding:"required" `
	Email           string `json:"email" binding:"required,email" `
	UniversityName  string `json:"university_name"`
}

type CreateStudentResponse struct {
	Username       string `json:"username"`
	Fullname       string `json:"full_name"`
	Email          string `json:"email"`
	Description    string `json:"description"`
	Avatar         string `json:"avatar"`
	UniversityName string `json:"university_name"`
	Credit         int    `json:"credit"`
}

type LoginStudentRequest struct {
	Username string `json:"username" binding:"required,alphanum" `
	Password string `json:"password" binding:"required,min=6" `
}

type LoginStudentResponse struct {
	AccessToken          string                `json:"access_token"`
	AccessTokenExpiresAt time.Time             `json:"access_token_expires_at"`
	Student              CreateStudentResponse `json:"student"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

type PasswordResetVerifyRequest struct {
	Hash     string `json:"hash"`
	Password string `json:"password"`
}

type EditStudentProfilePageRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type EditStudentProfilePageResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type VerifyEmailRequest struct {
	Hash string `json:"hash"`
}

type StudentAddCourseRequest struct {
	CourseId string `json:"courseId"`
}

type StudentSubscriptionRequest struct {
	PlanType      string `json:"plan_type"`
	SubExpireTime string `json:"sub_expire_time"`
}

type StudentSubscriptionTypeChangeRequest struct {
	PlanType string `json:"plan_type"`
}

type StudentSubscriptionExtendRequest struct {
	ExtensionTime string `json:"extension_time"`
}

type SubscriptionResponse struct {
	PlanType      string `json:"plan_type"`
	RemainingTime int    `json:"remaining_time"`
}

type CreateMentorRequest struct {
	Username        string                 `json:"username" binding:"required,alphanum" `
	Password        string                 `json:"password" binding:"required,min=6" `
	ConfirmPassword string                 `json:"confirm_password" binding:"required,min=6" `
	FullName        string                 `json:"full_name" binding:"required" `
	UniversityName  string                 `json:"university_name"`
	Email           string                 `json:"email" binding:"required,email" `
	Description     string                 `json:"description"`
	Address         CreateAddressesRequest `json:"address"`
	CourseName      string                 `json:"course_name"`
}

type CreateMentorResponse struct {
	Username       string                 `json:"username"`
	Fullname       string                 `json:"full_name"`
	Email          string                 `json:"email"`
	UniversityName string                 `json:"university_name"`
	Description    string                 `json:"description"`
	Score          int                    `json:"score"`
	Balance        int                    `json:"balance"`
	Address        CreateAddressesRequest `json:"address"`
	CourseName     string                 `json:"course_name"`
}

type EvaluateMentorRequest struct {
	CID   string `json:"c_id"`
	Score int    `json:"score"`
}

type LoginMentorResponse struct {
	AccessToken          string               `json:"access_token"`
	AccessTokenExpiresAt time.Time            `json:"access_token_expires_at"`
	Mentor               CreateMentorResponse `json:"mentor"`
}

type CreateEventRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Title     string `json:"title"`
	Color     string `json:"color"`
}

type CreateUniversityRequest struct {
	UniversityName string                 `json:"university_name"`
	Abbreviation   string                 `json:"abbreviation"`
	EmailExtension string                 `json:"email_extension"`
	Address        CreateAddressesRequest `json:"address"`
}

type CreateAddressesRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type CreateCourseRequest struct {
	CourseName     string `json:"course_name"`
	Semester       string `json:"semester"`
	UniversityName string `json:"university_name"`
	Description    string `json:"description"`
}

type ReadCourseStudentResponse struct {
	CourseName  string `json:"course_name"`
	Semester    string `json:"semester"`
	Description string `json:"description"`
}
