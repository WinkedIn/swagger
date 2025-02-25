// Package winkedIn provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package winkedIn

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get list of miscellaneous items
	// (POST /v1/getItems/{itemType})
	GetItems(w http.ResponseWriter, r *http.Request, itemType ItemType)
	// insert miscellaneous items
	// (POST /v1/insertItems)
	InsertItems(w http.ResponseWriter, r *http.Request)
	// Creates a new user
	// (POST /v1/users)
	CreateUsers(w http.ResponseWriter, r *http.Request)
	// fetch user's personal information
	// (GET /v1/users/{userID}/personalInfo)
	GetUsersPersonal(w http.ResponseWriter, r *http.Request, userID UserID)
	// Update user's personal information
	// (PUT /v1/users/{userID}/personalInfo)
	UpdateUsersPersonal(w http.ResponseWriter, r *http.Request, userID UserID)
	// Upload user photos
	// (POST /v1/users/{userID}/photos)
	UploadUserPhotos(w http.ResponseWriter, r *http.Request, userID UserID)
	// fetch user's preferences
	// (GET /v1/users/{userID}/preferences)
	GetUsersPreferences(w http.ResponseWriter, r *http.Request, userID UserID, params GetUsersPreferencesParams)
	// Add/Update user's preferences
	// (POST /v1/users/{userID}/preferences)
	UpdateUsersPreferences(w http.ResponseWriter, r *http.Request, userID UserID)
	// fetch user's professional information
	// (GET /v1/users/{userID}/professionalInfo)
	GetUsersProfessional(w http.ResponseWriter, r *http.Request, userID UserID)
	// Update user's professional information
	// (PUT /v1/users/{userID}/professionalInfo)
	UpdateUsersProfessional(w http.ResponseWriter, r *http.Request, userID UserID)
	// Send OTP to user
	// (POST /v1/users/{userID}/sendOTP)
	SendOTP(w http.ResponseWriter, r *http.Request, userID UserID)
	// Verify OTP
	// (POST /v1/users/{userID}/verifyOTP)
	VerifyOTP(w http.ResponseWriter, r *http.Request, userID UserID)
	// Verify users photos
	// (POST /v1/users/{userID}/verifyPhotos)
	VerifyUserPhotos(w http.ResponseWriter, r *http.Request, userID UserID)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get list of miscellaneous items
// (POST /v1/getItems/{itemType})
func (_ Unimplemented) GetItems(w http.ResponseWriter, r *http.Request, itemType ItemType) {
	w.WriteHeader(http.StatusNotImplemented)
}

// insert miscellaneous items
// (POST /v1/insertItems)
func (_ Unimplemented) InsertItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Creates a new user
// (POST /v1/users)
func (_ Unimplemented) CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// fetch user's personal information
// (GET /v1/users/{userID}/personalInfo)
func (_ Unimplemented) GetUsersPersonal(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user's personal information
// (PUT /v1/users/{userID}/personalInfo)
func (_ Unimplemented) UpdateUsersPersonal(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Upload user photos
// (POST /v1/users/{userID}/photos)
func (_ Unimplemented) UploadUserPhotos(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// fetch user's preferences
// (GET /v1/users/{userID}/preferences)
func (_ Unimplemented) GetUsersPreferences(w http.ResponseWriter, r *http.Request, userID UserID, params GetUsersPreferencesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add/Update user's preferences
// (POST /v1/users/{userID}/preferences)
func (_ Unimplemented) UpdateUsersPreferences(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// fetch user's professional information
// (GET /v1/users/{userID}/professionalInfo)
func (_ Unimplemented) GetUsersProfessional(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user's professional information
// (PUT /v1/users/{userID}/professionalInfo)
func (_ Unimplemented) UpdateUsersProfessional(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Send OTP to user
// (POST /v1/users/{userID}/sendOTP)
func (_ Unimplemented) SendOTP(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Verify OTP
// (POST /v1/users/{userID}/verifyOTP)
func (_ Unimplemented) VerifyOTP(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Verify users photos
// (POST /v1/users/{userID}/verifyPhotos)
func (_ Unimplemented) VerifyUserPhotos(w http.ResponseWriter, r *http.Request, userID UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetItems operation middleware
func (siw *ServerInterfaceWrapper) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "itemType" -------------
	var itemType ItemType

	err = runtime.BindStyledParameterWithOptions("simple", "itemType", chi.URLParam(r, "itemType"), &itemType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemType", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItems(w, r, itemType)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// InsertItems operation middleware
func (siw *ServerInterfaceWrapper) InsertItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.InsertItems(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateUsers operation middleware
func (siw *ServerInterfaceWrapper) CreateUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersPersonal operation middleware
func (siw *ServerInterfaceWrapper) GetUsersPersonal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersPersonal(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUsersPersonal operation middleware
func (siw *ServerInterfaceWrapper) UpdateUsersPersonal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUsersPersonal(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UploadUserPhotos operation middleware
func (siw *ServerInterfaceWrapper) UploadUserPhotos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UploadUserPhotos(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersPreferences operation middleware
func (siw *ServerInterfaceWrapper) GetUsersPreferences(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersPreferencesParams

	// ------------- Optional query parameter "preferenceType" -------------

	err = runtime.BindQueryParameter("form", true, false, "preferenceType", r.URL.Query(), &params.PreferenceType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "preferenceType", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersPreferences(w, r, userID, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUsersPreferences operation middleware
func (siw *ServerInterfaceWrapper) UpdateUsersPreferences(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUsersPreferences(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersProfessional operation middleware
func (siw *ServerInterfaceWrapper) GetUsersProfessional(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersProfessional(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUsersProfessional operation middleware
func (siw *ServerInterfaceWrapper) UpdateUsersProfessional(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUsersProfessional(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SendOTP operation middleware
func (siw *ServerInterfaceWrapper) SendOTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SendOTP(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// VerifyOTP operation middleware
func (siw *ServerInterfaceWrapper) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.VerifyOTP(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// VerifyUserPhotos operation middleware
func (siw *ServerInterfaceWrapper) VerifyUserPhotos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.VerifyUserPhotos(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/getItems/{itemType}", wrapper.GetItems)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/insertItems", wrapper.InsertItems)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users", wrapper.CreateUsers)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/users/{userID}/personalInfo", wrapper.GetUsersPersonal)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/v1/users/{userID}/personalInfo", wrapper.UpdateUsersPersonal)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/{userID}/photos", wrapper.UploadUserPhotos)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/users/{userID}/preferences", wrapper.GetUsersPreferences)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/{userID}/preferences", wrapper.UpdateUsersPreferences)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/users/{userID}/professionalInfo", wrapper.GetUsersProfessional)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/v1/users/{userID}/professionalInfo", wrapper.UpdateUsersProfessional)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/{userID}/sendOTP", wrapper.SendOTP)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/{userID}/verifyOTP", wrapper.VerifyOTP)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/users/{userID}/verifyPhotos", wrapper.VerifyUserPhotos)
	})

	return r
}
