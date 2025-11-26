package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rafaelchr/task-management/models"
	"github.com/rafaelchr/task-management/services"
	"github.com/rafaelchr/task-management/utils"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Register Failed", err.Error())
	}

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return utils.Success(ctx, "Register Success", userResponse)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Login Failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return utils.Success(ctx, "Login Success", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          userResponse,
	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "User not Found", err.Error())
	}

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return utils.Success(ctx, "Success getting user", userResponse)
}

func (c *UserController) GetUserPagination(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	users, total, err := c.service.GetAllPagination(filter, sort, limit, offset)
	if err != nil {
		return utils.BadRequest(ctx, "Gagal Mengambil Data", err.Error())
	}

	var userResponse []models.UserResponse
	_ = copier.Copy(&userResponse, &users)

	meta := utils.PaginationMeta{
		Page:      page,
		Limit:     limit,
		Total:     int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter:    filter,
		Sort:      sort,
	}

	if total == 0 {
		return utils.NotFoundPagination(ctx, "Data pengguna tidak ditemukan", userResponse, meta)
	}

	return utils.SuccessPagination(ctx, "Data ditemukan", userResponse, meta)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	publicID, err := uuid.Parse(id)
	if err != nil {
		return utils.BadRequest(ctx, "Invalid ID format", err.Error())
	}

	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}
	user.PublicID = publicID

	if err := c.service.Update(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal update data", err.Error())
	}

	userUpdated, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.InternalServerError(ctx, "Gagal mengambil data", err.Error())
	}

	var userResponse models.UserResponse
	err = copier.Copy(&userResponse, &userUpdated)
	if err != nil {
		return utils.InternalServerError(ctx, "Error parsing data", err.Error())
	}

	return utils.Success(ctx, "Berhasil update data user", userResponse)
}
