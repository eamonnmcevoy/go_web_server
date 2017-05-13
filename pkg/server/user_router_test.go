package server

import (
	"bytes"
	"encoding/json"
	"go_rest_api/pkg"
	"go_rest_api/pkg/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_UserRouter(t *testing.T) {
	t.Run("happy path", createUserHandler_should_pass_User_object_to_UserService_CreateUser)
	t.Run("invalid payload", get_user_handler_should_call_UserService_with_username_from_url)
}

func createUserHandler_should_pass_User_object_to_UserService_CreateUser(t *testing.T) {
	// Arrange
	us := mock.UserService{}
	testUserRouter := NewUserRouter(&us, mux.NewRouter())
	var result *root.User
	us.CreateUserFn = func(u *root.User) error {
		result = u
		return nil
	}

	testUsername := "test_username"
	testPassword := "test_password"

	values := map[string]string{"username": testUsername, "password": testPassword}
	jsonValue, _ := json.Marshal(values)
	payload := bytes.NewBuffer(jsonValue)

	// Act
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/", payload)
	r.Header.Set("Content-Type", "application/json")
	testUserRouter.ServeHTTP(w, r)

	// Assert
	if !us.CreateUserInvoked {
		t.Fatal("expected CreateUser() to be invoked")
	}
	if result.Username != testUsername {
		t.Fatalf("expected username to be: `%s`, got: `%s`", testUsername, result.Username)
	}
	if result.Password != testPassword {
		t.Fatalf("expected username to be: `%s`, got: `%s`", testPassword, result.Password)
	}
}

func get_user_handler_should_call_UserService_with_username_from_url(t *testing.T) {
	// Arrange
	us := mock.UserService{}
	testUserRouter := NewUserRouter(&us, mux.NewRouter())
	testUsername := "test_username"
	result := root.User{Username: testUsername}
	us.GetUserByUsernameFn = func(u string) (error, root.User) {
		return nil, result
	}

	// Act
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/"+testUsername, nil)
	r.Header.Set("Content-Type", "application/json")
	testUserRouter.ServeHTTP(w, r)

	// Assert
	if !us.GetUserByUsernameInvoked {
		t.Fatal("expected CreateUser() to be invoked")
	}

	expected, _ := json.Marshal(result)
	if string(expected) != w.Body.String() {
		t.Fatalf("expected response body to be: `%s`, got: `%s`", expected, w.Body.String())
	}
}
