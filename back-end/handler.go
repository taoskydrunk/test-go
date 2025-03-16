package main

import (
	"back-end/spec"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Handler struct {
	commentService spec.CommentService
}

func (h *Handler) register(app *fiber.App) {
	//external
	external := app.Group("/v1/comments")

	external.Get("", h.find)
	external.Post("", h.insert)
}

func (h *Handler) find(ctx *fiber.Ctx) error {
	results, err := h.commentService.Find()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "OK",
		"message": "Accept request is successful",
		"data":    results,
	})
}

func (h *Handler) insert(ctx *fiber.Ctx) error {
	body := new(spec.CommentBody)
	_ = ctx.BodyParser(body)

	model := spec.CommentModel{
		ID:          primitive.NewObjectID(),
		Text:        body.Text,
		CreatedBy:   body.CreatedBy,
		CreatedDate: time.Now(),
	}

	if err := h.commentService.Insert(model); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(map[string]string{
		"status":  "OK",
		"message": "Accept request is successful",
	})
}
