package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"swaggo/internal/users/pkg/api"

	"net/http"
)

func RegisterRoutes(router chi.Router) {
	resource := usersResource{}

	router.Route("/users", func(r chi.Router) {
		r.Get("/list", resource.handleListCustomers)
	})
}

type usersResource struct{}

type customerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// @Summary      Some summary
// @Description  Some description
// @Tags         users
// @Security     Token
// @Produce      json
// @Success      200  {object}	api.PaginatedResponse[customerResponse]
// @Router       /users/list      [get]
func (c usersResource) handleListCustomers(w http.ResponseWriter, r *http.Request) {
	resp := api.PaginatedResponse[customerResponse]{
		PageNumber: 1,
		PageSize:   2,
		Total:      3,
		TotalPages: 2,
		Data: []customerResponse{
			{
				ID:   "id3",
				Name: "name3",
			},
		},
	}
	render.JSON(w, r, resp)
}
