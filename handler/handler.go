package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var HandlerModule = fx.Provide(NewHandler)
var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FistName string `db:"first_name"`
	LastName string `db:"last_name"`
	Email    string
}

type Handler struct {
	db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Hello(c *fiber.Ctx) error {
	h.db.MustExec(schema)
	tx := h.db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.Commit()

	people := []Person{}
	h.db.Select(&people, "select * from person order by first_name asc")

	return c.JSON(people)
}
