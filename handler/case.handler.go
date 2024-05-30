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

func GetCaseByUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	var cases []entity.Case
	database.DB.Where("client_id = ?", id).Find(&cases)

	return c.Status(200).JSON(fiber.Map{
		"message": "Case by user ID",
		"case":    cases,
	})
}

func UpdateCaseByID(c *fiber.Ctx) error {
	updateCase := new(dto.UpdateCaseByIDRequestDTO)

	if err := c.BodyParser(updateCase); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing update case",
			"error":   err.Error(),
		})
	}

	var cases entity.Case
	database.DB.Where("id = ?", updateCase.ID).First(&cases)

	if cases.ID == (entity.Case{}).ID {
		return c.Status(404).JSON(fiber.Map{
			"message": "Case not found",
		})
	}

	cases.Subject = updateCase.Subject
	cases.Media = updateCase.Media
	cases.Notes = updateCase.Notes
	cases.Status = updateCase.Status
	cases.Hour = updateCase.Hour
	cases.AdditionFee = updateCase.AdditionFee
	cases.ClientID = updateCase.ClientID
	cases.LawyerID = updateCase.LawyerID

	database.DB.Save(&cases)

	return c.Status(200).JSON(fiber.Map{
		"message": "Case updated",
		"case":    cases,
	})

}

func DeleteCaseByID(c *fiber.Ctx) error {
	deleteCase := new(dto.DeleteCaseRequestDTO)

	if err := c.BodyParser(deleteCase); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing delete case",
			"error":   err.Error(),
		})
	}

	var cases entity.Case
	database.DB.Where("id = ?", deleteCase.ID).First(&cases)

	if cases.ID == (entity.Case{}).ID {
		return c.Status(404).JSON(fiber.Map{
			"message": "Case not found",
		})
	}

	database.DB.Delete(&cases)

	return c.Status(200).JSON(fiber.Map{
		"message": "Case deleted",
		"case":    cases,
	})
}
