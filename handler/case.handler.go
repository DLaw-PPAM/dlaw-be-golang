package handler

import (
	"example.com/m/v2/database"
	"example.com/m/v2/model/dto"
	"example.com/m/v2/model/entity"
	"github.com/gofiber/fiber/v2"
)

func AddNewCase(c *fiber.Ctx) error {
	newCase := new(dto.AddCaseRequestDTO)

	if err := c.BodyParser(newCase); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new case",
			"error":   err.Error(),
		})
	}

	var Lawyer entity.LawyerUser
	database.DB.Where("id = ?", newCase.LawyerID).First(&Lawyer)

	//if clientid is the same as lawyerid return
	if Lawyer.ClientID == newCase.LawyerID {
		return c.Status(400).JSON(fiber.Map{
			"message": "Client and lawyer cannot be the same",
		})
	}

	// Add case to database
	newCaseReq := entity.Case{
		Subject:     newCase.Subject,
		Media:       newCase.Media,
		Notes:       newCase.Notes,
		Status:      newCase.Status,
		Hour:        newCase.Hour,
		AdditionFee: newCase.AdditionFee,
		ClientID:    newCase.ClientID,
		LawyerID:    newCase.LawyerID,
	}

	newCaseRes := database.DB.Create(&newCaseReq)

	if newCaseRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new case",
			"error":   newCaseRes.Error.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "New case created",
		"case":    newCase,
	})
}

func GetAllCase(c *fiber.Ctx) error {
	var cases []entity.Case
	database.DB.Find(&cases)

	return c.Status(200).JSON(fiber.Map{
		"message": "All cases",
		"cases":   cases,
	})
}

func GetCaseByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var cases entity.Case
	database.DB.Where("id = ?", id).First(&cases)

	return c.Status(200).JSON(fiber.Map{
		"message": "Case by ID",
		"case":    cases,
	})
}
