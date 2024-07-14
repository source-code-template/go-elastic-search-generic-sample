package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
	"github.com/gorilla/mux"

	"go-service/internal/user/model"
	"go-service/internal/user/service"
)

const InternalServerError = "Internal Server Error"

func NewUserHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service service.UserService, validate func(context.Context, interface{}) ([]core.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *UserHandler {
	filterType := reflect.TypeOf(model.UserFilter{})
	modelType := reflect.TypeOf(model.User{})
	_, jsonMap, _ := core.BuildMapField(modelType)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, nil)
	return &UserHandler{service: service, SearchHandler: searchHandler, Validate: validate, jsonMap: jsonMap}
}

type UserHandler struct {
	service service.UserService
	*search.SearchHandler
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
	JSON(w, http.StatusOK, users)
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
		JSON(w, http.StatusNotFound, nil)
	} else {
		JSON(w, http.StatusOK, user)
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
		h.LogError(r.Context(), er2.Error(), MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Create(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error(), MakeMap(user))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if res > 0 {
		JSON(w, http.StatusCreated, user)
	} else {
		JSON(w, http.StatusConflict, res)
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
		h.LogError(r.Context(), er2.Error(), MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Update(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error(), MakeMap(user))
		http.Error(w, er3.Error(), http.StatusInternalServerError)
		return
	}
	if res > 0 {
		JSON(w, http.StatusOK, user)
	} else if res == 0 {
		JSON(w, http.StatusNotFound, res)
	} else {
		JSON(w, http.StatusConflict, res)
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
		h.LogError(r.Context(), er3.Error(), MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er4 := h.service.Patch(r.Context(), jsonUser)
	if er4 != nil {
		h.LogError(r.Context(), er4.Error(), MakeMap(user))
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if res > 0 {
		JSON(w, http.StatusOK, jsonUser)
	} else if res == 0 {
		JSON(w, http.StatusNotFound, res)
	} else {
		JSON(w, http.StatusConflict, res)
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
		JSON(w, http.StatusOK, res)
	} else {
		JSON(w, http.StatusNotFound, res)
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
func IsFound(res interface{}) int {
	if isNil(res) {
		return http.StatusNotFound
	}
	return http.StatusOK
}
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
func MakeMap(res interface{}, opts ...string) map[string]interface{} {
	key := "request"
	if len(opts) > 0 && len(opts[0]) > 0 {
		key = opts[0]
	}
	m := make(map[string]interface{})
	b, err := json.Marshal(res)
	if err != nil {
		return m
	}
	m[key] = string(b)
	return m
}
