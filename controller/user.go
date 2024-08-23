package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	repoUser repositories.UserInterface
}

// NewUserController will create an object that represent the User.Article interface
func NewUserController(repoUser repositories.UserInterface) *User {
	return &User{repoUser}
}

// Register will create a new user
// @Summary      Register new user
// @Description  Register new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request  body  model.UserCreate true "User data"
// @Success      201  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/register [post]
func (u *User) Register(ctx *gin.Context) {
	user := &model.UserCreate{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	newUser, err := u.repoUser.Create(user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.MailClient().SendMailAsync(
		[]string{newUser.Email},
		[]string{},
		"OTP",
		utils.CreateHTMLOTP(newUser.Name, newUser.Province, newUser.City, newUser.OTP),
	)

	utils.SuccessResponse(ctx, http.StatusCreated, newUser)
}

// Login will login user
// @Summary      Login user
// @Description  Login user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request  body  model.UserLogin true "User data"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/login [post]
func (u *User) Login(ctx *gin.Context) {
	var user model.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(utils.NewError(utils.ErrValidation, "email atau password tidak valid"))
		ctx.Next()
		return
	}

	userData, err := u.repoUser.Login(user.Email, user.Password)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	token, err := utils.GenerateToken(userData.ID.String(), userData.Email, userData.Role)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, token)
}

// Update will update user data
// @Summary      Update user data
// @Description  Update user data
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Param 		 request  body  model.UserUpdate true "User data"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user [put]
func (u *User) Update(ctx *gin.Context) {
	user := &model.UserUpdate{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	// Ambil uuid yang disimpan di context dari middleware JWT
	claimsData := ctx.MustGet("id").(string)
	// convert claimsData["id"] ke string dan ke uuid
	userID := uuid.Must(uuid.Parse(claimsData))

	userData, err := u.repoUser.Update(userID, user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, userData)
}

// FindOne will find user by id
// @Summary      Find user by id
// @Description  Find user by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user [get]
func (u *User) FindOne(ctx *gin.Context) {
	id := uuid.Must(uuid.Parse(ctx.MustGet("id").(string)))
	user, err := u.repoUser.FindOne(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, user)
}

// VerifyOTP will verify otp
// @Summary      Verify OTP
// @Description  Verify OTP
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 otp  path  string true "OTP"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/verify/{otp} [get]
func (u *User) VerifyOTP(ctx *gin.Context) {
	otp := ctx.Param("otp")
	user, err := u.repoUser.VerifyOTP(otp)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, user)
}

// ResendEmailOTP will resend email OTP
// @Summary      Resend email OTP
// @Description  Resend email OTP
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 email  path  string true "User email"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/resend/{email} [get]
func (u *User) ResendEmailOTP(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := u.repoUser.ResendEmailOTP(email)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	err = utils.MailClient().SendMailSync(
		[]string{user.Email},
		[]string{},
		"OTP Islamind",
		utils.CreateHTMLOTP(user.Name, user.Province, user.City, user.OTP),
	)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, user)
}

// Logout will logout user
// @Summary      Logout user
// @Description  Logout user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/logout [get]
func (u *User) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		_ = ctx.Error(utils.NewError(utils.ErrUnauthorized, "Token is required"))
		ctx.Next()
		return
	}

	durationToken, err := utils.GetExpiredToken(token)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	if err := u.repoUser.Logout(token, durationToken); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Logout success")
}
