package handler

import (
	"github.com/pikachu0310/go-backend-template/internal/repository"
	"github.com/pikachu0310/go-backend-template/openapi/models"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
}

func (h *Handler) OauthCallback(ctx echo.Context, params models.OauthCallbackParams) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Login(ctx echo.Context, params models.LoginParams) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Logout(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEvents(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostEvent(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetCurrentEvent(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEvent(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PatchEvent(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEventCsv(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEventGames(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEventImage(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetEventTerms(ctx echo.Context, eventSlug models.EventSlugInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetGames(ctx echo.Context, params models.GetGamesParams) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostGame(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetGame(ctx echo.Context, gameId models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PatchGame(ctx echo.Context, gameId models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetGameIcon(ctx echo.Context, gameId models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetGameImage(ctx echo.Context, gameId models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PingServer(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetTerms(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostTerm(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetTerm(ctx echo.Context, termId models.TermIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PatchTerm(ctx echo.Context, termId models.TermIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetTermGames(ctx echo.Context, termId models.TermIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Test(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetMe(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetMeGames(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetUserGames(ctx echo.Context, userId models.UserIdInPath) error {
	//TODO implement me
	panic("implement me")
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
