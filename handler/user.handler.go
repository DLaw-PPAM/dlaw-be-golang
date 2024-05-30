package handler

import (
	"time"

	"example.com/m/v2/database"
	"example.com/m/v2/model/dto"
	"example.com/m/v2/model/entity"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetHelloWorld(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Hello World",
	})
}

func UserRegister(c *fiber.Ctx) error {
	user := new(dto.UserRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new user",
			"error":   errValidate.Error(),
		})
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing BirthDate",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		FullName:    user.FullName,
		Username:    user.Username,
		Email:       user.Email,
		Password:    string(hashedPassword),
		BirthDate:   birthDate,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Bio:         user.Bio,
	}

	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	var existingUser2 entity.User
	res2 := database.DB.Where("username = ?", user.Username).First(&existingUser2)
	if res2.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)

	if newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new user",
			"error":   newUserRes.Error.Error(),
		})
	}

	responseDTO := dto.UserRegisterResponseDTO{
		Message:     "New user created successfully",
		ID:          newUser.ID,
		Username:    newUser.Username,
		FullName:    newUser.FullName,
		Email:       newUser.Email,
		BirthDate:   newUser.BirthDate.Format("2006-01-02"),
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
		Bio:         newUser.Bio,
	}

	return c.Status(201).JSON(responseDTO)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []entity.User
	database.DB.Find(&users)

	return c.Status(200).JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User
	database.DB.Where("id = ?", id).First(&user)

	return c.Status(200).JSON(user)
}

func AddSpecialties(c *fiber.Ctx) error {
	specialties := new(dto.AddSpecialtiesRequestDTO)

	if err := c.BodyParser(specialties); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new specialties",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(specialties)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new specialties",
			"error":   errValidate.Error(),
		})
	}

	newSpecialties := entity.Specialties{
		Name: specialties.Name,
	}

	newSpecialtiesRes := database.DB.Create(&newSpecialties)

	if newSpecialtiesRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new specialties",
			"error":   newSpecialtiesRes.Error.Error(),
		})
	}

	responseDTO := dto.AddSpecialtiesResponseDTO{
		Message: "New specialties created successfully",
		ID:      newSpecialties.ID,
		Name:    newSpecialties.Name,
	}

	return c.Status(201).JSON(responseDTO)
}

func GetAllSpecialties(c *fiber.Ctx) error {
	var specialties []entity.Specialties
	database.DB.Find(&specialties)

	return c.Status(200).JSON(specialties)
}
