package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search/handler"
	"github.com/gorilla/mux"

	"go-service/internal/user/model"
	"go-service/internal/user/service"
)

const InternalServerError = "Internal Server Error"

func NewUserHandler(find func(context.Context, *model.UserFilter, int64, int64) ([]model.User, int64, error), service service.UserService, validate func(context.Context, interface{}) ([]core.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *UserHandler {
	userType := reflect.TypeOf(model.User{})
	_, jsonMap, _ := core.BuildMapField(userType)
	searchHandler := search.NewSearchHandler[model.User, *model.UserFilter](find, logError, nil)
	return &UserHandler{service: service, SearchHandler: searchHandler, Validate: validate, jsonMap: jsonMap}
}

type UserHandler struct {
	service service.UserService
	*search.SearchHandler[model.User, *model.UserFilter]
	Validate func(context.Context, interface{}) ([]core.ErrorMessage, error)
	jsonMap  map[string]int
}

func (h *UserHandler) All(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.All(r.Context())
	if err != nil {
		h.LogError(r.Context(), fmt.Sprintf("Error: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, users)
}
func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.Load(r.Context(), id)
	if err != nil {
		h.LogError(r.Context(), fmt.Sprintf("Error to get user %s: %s", id, err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		core.JSON(w, http.StatusNotFound, nil)
	} else {
		core.JSON(w, http.StatusOK, user)
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	errors, er2 := h.Validate(r.Context(), &user)
	if er2 != nil {
		h.LogError(r.Context(), er2.Error(), core.MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		core.JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Create(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error(), core.MakeMap(user))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if res > 0 {
		core.JSON(w, http.StatusCreated, user)
	} else {
		core.JSON(w, http.StatusConflict, res)
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	errors, er2 := h.Validate(r.Context(), &user)
	if er2 != nil {
		h.LogError(r.Context(), er2.Error(), core.MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		core.JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Update(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error(), core.MakeMap(user))
		http.Error(w, er3.Error(), http.StatusInternalServerError)
		return
	}
	if res > 0 {
		core.JSON(w, http.StatusOK, user)
	} else if res == 0 {
		core.JSON(w, http.StatusNotFound, res)
	} else {
		core.JSON(w, http.StatusConflict, res)
	}
}
func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	var user model.User
	body, er1 := core.BuildMapAndStruct(r, &user)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	jsonUser, er2 := core.BodyToJsonMap(r, user, body, []string{"id"}, h.jsonMap)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "method", "patch"))
	errors, er3 := h.Validate(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error(), core.MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		core.JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er4 := h.service.Patch(r.Context(), jsonUser)
	if er4 != nil {
		h.LogError(r.Context(), er4.Error(), core.MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if res > 0 {
		core.JSON(w, http.StatusOK, jsonUser)
	} else if res == 0 {
		core.JSON(w, http.StatusNotFound, res)
	} else {
		core.JSON(w, http.StatusConflict, res)
	}
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.Delete(r.Context(), id)
	if err != nil {
		h.LogError(r.Context(), fmt.Sprintf("Error to delete user %s: %s", id, err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res > 0 {
		core.JSON(w, http.StatusOK, res)
	} else {
		core.JSON(w, http.StatusNotFound, res)
	}
}
