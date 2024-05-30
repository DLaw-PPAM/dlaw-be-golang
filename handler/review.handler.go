package handler

import (
	"example.com/m/v2/database"
	"example.com/m/v2/model/dto"
	"example.com/m/v2/model/entity"
	"github.com/gofiber/fiber/v2"
)

func AddReview(c *fiber.Ctx) error {
	newReview := new(dto.AddReviewRequestDTO)

	if err := c.BodyParser(newReview); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new review",
			"error":   err.Error(),
		})
	}

	// Add review to database
	newReviewReq := entity.Review{
		Rating:      newReview.Rating,
		Description: newReview.Description,
		ClientID:    newReview.ClientID,
		LawyerID:    newReview.LawyerID,
		ClientName:  newReview.ClientName,
	}

	newReviewRes := database.DB.Create(&newReviewReq)

	if newReviewRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new review",
			"error":   newReviewRes.Error.Error(),
		})
	}

	var lawyer entity.LawyerUser
	database.DB.Where("id = ?", newReview.LawyerID).First(&lawyer)
	lawyer.Reviews = append(lawyer.Reviews, newReviewReq)
	database.DB.Save(&lawyer)

	return c.Status(201).JSON(fiber.Map{
		"message": "New review created",
		"review":  newReview,
	})

}
