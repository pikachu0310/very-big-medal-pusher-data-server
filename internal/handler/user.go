package handler

//// スキーマ定義
//type (
//	GetUsersResponse []GetUserResponse
//
//	GetUserResponse struct {
//		ID    uuid.UUID `json:"id"`
//		Name  string    `json:"name"`
//		Email string    `json:"email"`
//	}
//
//	CreateUserRequest struct {
//		Name  string `json:"name"`
//		Email string `json:"email"`
//	}
//
//	CreateUserResponse struct {
//		ID uuid.UUID `json:"id"`
//	}
//)
//
//// GET /api/v1/users
//func (h *Handler) GetUsers(c echo.Context) error {
//	users, err := h.repo.GetUsers(c.Request().Context())
//	if err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
//	}
//
//	res := make(GetUsersResponse, len(users))
//	for i, user := range users {
//		res[i] = GetUserResponse{
//			ID:    user.ID,
//			Name:  user.Name,
//			Email: user.Email,
//		}
//	}
//
//	return c.JSON(http.StatusOK, res)
//}
//
//// POST /api/v1/users
//func (h *Handler) CreateUser(c echo.Context) error {
//	req := new(CreateUserRequest)
//	if err := c.Bind(req); err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
//	}
//
//	err := vd.ValidateStruct(
//		req,
//		vd.Field(&req.Name, vd.Required),
//		vd.Field(&req.Email, vd.Required, is.Email),
//	)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err)).SetInternal(err)
//	}
//
//	userID, err := h.repo.CreateUser(c.Request().Context(), repository.CreateUserParams{
//		Name:  req.Name,
//		Email: req.Email,
//	})
//	if err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
//	}
//
//	res := CreateUserResponse{
//		ID: userID,
//	}
//
//	return c.JSON(http.StatusOK, res)
//}
