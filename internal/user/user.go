package user

import (
	"context"
	"net/http"

	v "github.com/core-go/core/v10"
	repo "github.com/core-go/elasticsearch/repository"
	"github.com/core-go/search"
	"github.com/core-go/search/elasticsearch/query"
	"github.com/elastic/go-elasticsearch/v8"

	"go-service/internal/user/handler"
	"go-service/internal/user/model"
	"go-service/internal/user/service"
)

type UserTransport interface {
	All(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(client *elasticsearch.Client, logError func(context.Context, string, ...map[string]interface{})) (UserTransport, error) {
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	userQueryBuilder := query.NewBuilder[model.User, *model.UserFilter]()
	userSearchBuilder := repo.NewSearchBuilder[model.User, *model.UserFilter](client, []string{"users"}, userQueryBuilder.BuildQuery, search.GetSort)
	userRepository := repo.NewRepository[model.User](client, "users")
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userSearchBuilder.Search, userService, validator.Validate, logError)
	return userHandler, nil
}
