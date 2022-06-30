package models

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}