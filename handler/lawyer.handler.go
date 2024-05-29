package handler

import (
	"example.com/m/v2/database"
	"example.com/m/v2/model/dto"
	"example.com/m/v2/model/entity"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func AddLawyer(c *fiber.Ctx) error {
	lawyer := new(dto.AddLawyerRequestDTO)

	if err := c.BodyParser(lawyer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new lawyer",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(lawyer)

	var User entity.User
	database.DB.Where("id = ?", lawyer.ClientID).First(&User)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new lawyer",
			"error":   errValidate.Error(),
		})
	}

	// Add lawyer to database
	newLawyer := entity.LawyerUser{
		ClientID:     lawyer.ClientID,
		PricePerHour: lawyer.PricePerHour,
		Rating:       lawyer.Rating,
		User:         User,
	}

	newLawyerRes := database.DB.Create(&newLawyer)

	if newLawyerRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new lawyer",
			"error":   newLawyerRes.Error.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "New lawyer created",
		"lawyer":  newLawyer,
	})
}

func GetAllLawyers(c *fiber.Ctx) error {
	var lawyers []entity.LawyerUser
	database.DB.Preload("Specialties").Find(&lawyers)

	return c.Status(200).JSON(fiber.Map{
		"message": "All lawyers",
		"lawyers": lawyers,
	})
}

func GetLawyerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var lawyer entity.LawyerUser
	database.DB.Preload("Specialties").Where("id = ?", id).First(&lawyer)

	return c.Status(200).JSON(fiber.Map{
		"message": "Lawyer found",
		"lawyer":  lawyer,
	})
}

func AddSpecialtiestoLawyer(c *fiber.Ctx) error {
	lawyerSpecialties := new(dto.AddSpecialtiesToLawyerRequestDTO)

	if err := c.BodyParser(lawyerSpecialties); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing new lawyer specialties",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(lawyerSpecialties)

	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validating new lawyer specialties",
			"error":   errValidate.Error(),
		})
	}

	// Retrieve the lawyer by ID
	var lawyer entity.LawyerUser
	if err := database.DB.Preload("Specialties").Where("id = ?", lawyerSpecialties.LawyerID).First(&lawyer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Lawyer not found",
			"error":   err.Error(),
		})
	}

	// Retrieve the specialty by ID
	var specialty entity.Specialties
	if err := database.DB.Where("id = ?", lawyerSpecialties.SpecialtiesID).First(&specialty).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Specialty not found",
			"error":   err.Error(),
		})
	}

	//check if the lawyer already has the specialty
	for _, s := range lawyer.Specialties {
		if s.ID == specialty.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Lawyer already has this specialty",
			})
		}
	}

	// Append the new specialty to the lawyer's specialties
	lawyer.Specialties = append(lawyer.Specialties, specialty)

	// Save the updated lawyer with the new list of specialties
	if err := database.DB.Save(&lawyer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating lawyer specialties",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Specialties added to lawyer",
		"lawyer":  lawyer,
	})
}
