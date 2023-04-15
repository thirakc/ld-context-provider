package ldcontext

import (
	ldcontextsvc "ld-context-provider/pkg/controller/service/ldcontext"
	"ld-context-provider/pkg/httpserver"
	"ld-context-provider/pkg/internal/common"
	"ld-context-provider/pkg/internal/envelop"
	"ld-context-provider/pkg/logz"
	"ld-context-provider/pkg/utils/restutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Servicer interface {
	CreateLDContext(*ldcontextsvc.LDContextArg) common.Error
}

var logger = logz.NewLogger()

const (
	CreateLDContextPath = "/ldcontext"
)

const (
	CreateLDContextErrorDesc = "CreateLDContext Error"
)

type Handler struct {
	service Servicer
}

func NewHandler(s Servicer) *Handler {
	return &Handler{s}
}

func (h *Handler) GetRESTHandlers() []httpserver.HTTPHandler {
	return []httpserver.HTTPHandler{
		restutil.NewHTTPHandler(CreateLDContextPath, http.MethodPost, h.CreateLDContext),
	}
}

func (h *Handler) CreateLDContext(c *gin.Context) {
	var req CreateLDContextRequest
	if err := c.Bind(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, envelop.NewValidateError(ldcontextsvc.InvalidRequestErrorCode))
		return
	}

	if err := h.service.CreateLDContext(&req.LDContextArg); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, envelop.NewResponseError(err.Code(), CreateLDContextErrorDesc))
		return
	}

	c.JSON(http.StatusOK, envelop.NewResponseSuccess(nil))
}
