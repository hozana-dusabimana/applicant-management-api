package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"job-tracker/database"
	"job-tracker/models"
	"job-tracker/utils"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	// Get Redis configuration from environment variables
	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")

	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     getEnv("REDIS_PASSWORD", ""),
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Test Redis connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
	} else {
		log.Println("Redis connected successfully")
	}
}

// Helper function to get environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func CreateApplicant(c *fiber.Ctx) error {
	var applicant models.Applicant
	if err := c.BodyParser(&applicant); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Sanitize input
	applicant.Name = utils.SanitizeString(applicant.Name)
	applicant.Email = strings.ToLower(utils.SanitizeString(applicant.Email))
	applicant.Position = utils.SanitizeString(applicant.Position)
	applicant.Phone = utils.SanitizeString(applicant.Phone)
	applicant.Notes = utils.SanitizeString(applicant.Notes)

	// Validate required fields
	if applicant.Name == "" || applicant.Email == "" || applicant.Position == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name, email, and position are required"})
	}

	// Validate email format
	if !utils.ValidateEmail(applicant.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email format"})
	}

	// Validate phone if provided
	if applicant.Phone != "" && !utils.ValidatePhone(applicant.Phone) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid phone number format"})
	}

	// Set default status if not provided
	if applicant.Status == "" {
		applicant.Status = "pending"
	} else if !utils.ValidateStatus(applicant.Status) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid status value"})
	}

	// Check if email already exists
	var existingApplicant models.Applicant
	if err := database.DB.Where("email = ?", applicant.Email).First(&existingApplicant).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{"error": "Email already exists"})
	}

	if err := database.DB.Create(&applicant).Error; err != nil {
		log.Printf("Database error creating applicant: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create applicant"})
	}

	// Clear cache to ensure fresh data on next request
	// Note: In production, we'd use pattern matching to clear all paginated cache
	rdb.Del(ctx, "applicants_page_1_limit_10", "applicants_page_1_limit_20") // Clear common cache keys
	log.Printf("Created new applicant with ID: %d", applicant.ID)

	return c.Status(201).JSON(applicant)
}

func GetApplicants(c *fiber.Ctx) error {
	// Get query parameters for pagination
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Create cache key with pagination
	cacheKey := fmt.Sprintf("applicants_page_%d_limit_%d", pageInt, limitInt)

	val, err := rdb.Get(ctx, cacheKey).Result()

	if err == redis.Nil {
		// Cache miss - fetch from database
		var applicants []models.Applicant
		offset := (pageInt - 1) * limitInt

		if err := database.DB.Offset(offset).Limit(limitInt).Find(&applicants).Error; err != nil {
			log.Printf("Database error: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch applicants"})
		}

		// Cache the result for 3 minutes
		jsonData, _ := json.Marshal(applicants)
		rdb.Set(ctx, cacheKey, jsonData, time.Minute*3)

		log.Printf("Cache miss - fetched %d applicants from database", len(applicants))
		return c.JSON(fiber.Map{
			"data":  applicants,
			"page":  pageInt,
			"limit": limitInt,
		})

	} else if err != nil {
		log.Printf("Redis error: %v", err)
		// Fallback to database if Redis fails
		var applicants []models.Applicant
		offset := (pageInt - 1) * limitInt
		if err := database.DB.Offset(offset).Limit(limitInt).Find(&applicants).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch applicants"})
		}
		return c.JSON(fiber.Map{
			"data":  applicants,
			"page":  pageInt,
			"limit": limitInt,
		})
	}

	// Cache hit
	var applicants []models.Applicant
	json.Unmarshal([]byte(val), &applicants)
	log.Printf("Cache hit - returned %d applicants", len(applicants))

	return c.JSON(fiber.Map{
		"data":  applicants,
		"page":  pageInt,
		"limit": limitInt,
	})
}

func GetApplicant(c *fiber.Ctx) error {
	id := c.Params("id")
	var applicant models.Applicant

	if err := database.DB.First(&applicant, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Applicant not found"})
	}

	return c.JSON(applicant)
}

func UpdateApplicant(c *fiber.Ctx) error {
	id := c.Params("id")
	var applicant models.Applicant

	// Check if applicant exists
	if err := database.DB.First(&applicant, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Applicant not found"})
	}

	// Parse update data
	var updateData models.Applicant
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Update applicant
	if err := database.DB.Model(&applicant).Updates(updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update applicant"})
	}

	// Clear cache - TODO: implement proper cache invalidation
	rdb.Del(ctx, "applicants_page_1_limit_10", "applicants_page_1_limit_20")
	return c.JSON(applicant)
}

func DeleteApplicant(c *fiber.Ctx) error {
	id := c.Params("id")
	var applicant models.Applicant

	// Check if applicant exists
	if err := database.DB.First(&applicant, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Applicant not found"})
	}

	// Delete applicant
	if err := database.DB.Delete(&applicant).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete applicant"})
	}

	// Clear cache
	rdb.Del(ctx, "applicants_page_1_limit_10", "applicants_page_1_limit_20")
	return c.Status(200).JSON(fiber.Map{"message": "Applicant deleted successfully"})
}
