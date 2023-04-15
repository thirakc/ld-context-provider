package ldcontext

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"ld-context-provider/pkg/controller/service/ldcontext"
	"ld-context-provider/pkg/httpserver"
	"ld-context-provider/pkg/internal/common"
	mockldcontext "ld-context-provider/pkg/mock/ldcontext"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_NewHandler(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		handler := NewHandler(&mockldcontext.MockService{})
		require.NotNil(t, handler)
		require.Equal(t, 1, len(handler.GetRESTHandlers()))
	})
}

func TestHanlder_CreateLDContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		handler := NewHandler(&mockldcontext.MockService{})
		require.NotNil(t, handler)

		reqBytes, err := json.Marshal(&ldcontext.LDContextArg{Url: "mockUrl"})
		require.NoError(t, err)

		h := lookupHandler(t, handler, CreateLDContextPath, http.MethodPost)
		_, code := sendRequestToHandler(t, h, bytes.NewBuffer(reqBytes), CreateLDContextPath)

		require.Equal(t, http.StatusOK, code)
	})

	t.Run("should fail when service return error", func(t *testing.T) {
		handler := NewHandler(&mockldcontext.MockService{ErrCreateLDContext: common.NewError(1000, errors.New("mock error"))})
		require.NotNil(t, handler)

		reqBytes, err := json.Marshal(CreateLDContextRequest{ldcontext.LDContextArg{Url: "mockUrl"}})
		require.NoError(t, err)

		h := lookupHandler(t, handler, CreateLDContextPath, http.MethodPost)
		_, code := sendRequestToHandler(t, h, bytes.NewBuffer(reqBytes), CreateLDContextPath)

		require.Equal(t, http.StatusInternalServerError, code)
	})

	t.Run("should fail when wrong request", func(t *testing.T) {
		handler := NewHandler(&mockldcontext.MockService{})
		require.NotNil(t, handler)

		h := lookupHandler(t, handler, CreateLDContextPath, http.MethodPost)
		_, code := sendRequestToHandler(t, h, nil, CreateLDContextPath)

		require.Equal(t, http.StatusBadRequest, code)
	})
}

func lookupHandler(t *testing.T, handler *Handler, path, method string) httpserver.HTTPHandler {
	t.Helper()

	handlers := handler.GetRESTHandlers()
	require.NotEmpty(t, handlers)

	for _, h := range handlers {
		if h.Path() == path && h.Method() == method {
			return h
		}
	}

	require.Fail(t, "unable to find handler")

	return nil
}

func sendRequestToHandler(t *testing.T, handler httpserver.HTTPHandler, requestBody io.Reader, path string) (*bytes.Buffer, int) {
	t.Helper()

	req, err := http.NewRequestWithContext(context.Background(), handler.Method(), path, requestBody)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Handle(handler.Method(), handler.Path(), gin.HandlerFunc(handler.Handler()))

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Body, rr.Code
}
