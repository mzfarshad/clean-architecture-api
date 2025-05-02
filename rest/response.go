package rest

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/pkg/errs"
)

type Handler interface {
	// Handle sets the *fiber.Ctx status code using the error of the response
	//and marshals the response to Json, then calls *fiber.Ctx 's Next() method.
	// It has handled log internally.
	Handle(ctx *fiber.Ctx) error
}

type Success interface {
	Paginate(pagination Pagination) Handler
	Handler
}

func NewSuccess(data idto) Success {
	return &response{Data: data, Succeed: true}
}

func NewFailed(err error) Handler {
	if err == nil {
		panic("trying to create new failed rest response using nil error!")
	}

	var customErr errs.Error
	if !errors.As(err, &customErr) {
		customErr = errs.New(errs.Internal, "something went wrong").CausedBy(err).(errs.Error)
	}
	resp := new(response)
	resp.Error = new(responseErr).from(customErr)
	return resp
}

type response struct {
	Data       idto         `json:"data"`
	Pagination any          `json:"pagination,omitempty"`
	Error      *responseErr `json:"error,omitempty"`
	Succeed    bool         `json:"succeed"`
}

func (r *response) Paginate(p Pagination) Handler {
	if p != nil {
		r.Pagination = &pagination{
			PageSize:   p.Size(),
			Page:       p.Page(),
			TotalCount: p.Total(),
			Filters:    p.Filters(),
		}
	}
	return r
}

func (r *response) Handle(ctx *fiber.Ctx) error {
	// TODO: Set Layer to log context
	//logCtx := logs.SetLayer(ctx.Context(), logs.LayerHandler)

	// TODO: Set Url to log context
	//pathUrl := ctx.Path()
	//logCtx = logs.SetUrl(logCtx, pathUrl)

	// TODO: set client ip
	//ip := ctx.Get("Cf-Connecting-Ip", "")
	//logCtx = logs.SetClientIP(logCtx, ip)

	// TODO: Set UserId to log context
	//claims, ok := ctx.Context().UserValue(auth.CtxKeyAuthUser).(*auth.UserClaims)
	//if ok && claims != nil {
	//	logCtx = logs.SetUserId(logCtx, claims.UserId)
	//}

	// TODO: Set StatusCode to log context
	statusCode := r.Error.httpStatus()
	//logCtx = logs.SetHttpStatus(logCtx, statusCode)
	// Setting up ctx status code
	ctx.Status(statusCode)
	// Handle response
	switch r.Data.(type) {
	case File:
		file := r.Data.(File)
		// TODO: Set http response to log context
		//logCtx = logs.SetHttpResponse(logCtx, file.Name)
		//logs.Info(logCtx, fmt.Sprintf("[API][Succeed] %s", pathUrl))
		ctx.Status(fiber.StatusOK)
		ctx.Set(fiber.HeaderAccessControlExposeHeaders, fiber.HeaderContentDisposition)
		ctx.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", file.Name))
		ctx.Set(fiber.HeaderContentType, file.Type.Extension)
		return ctx.Send(file.Bytes)

	default: // DTO, List, Map
		// TODO: Set http response to log context
		//logCtx = logs.SetHttpResponse(logCtx, *r)
		if r.Succeed {
			//logs.Info(logCtx, fmt.Sprintf("[API][Succeed] %s", pathUrl))
			//return ctx.Next() TODO: uncomment when add a middleware which will be called after handlers
		} else { // Failed
			//msg := strings.Join(r.Error.Messages, ",")
			//logs.Error(logCtx, fmt.Sprintf("[API][Failed] %s >>> %q error: %s", pathUrl, r.Error.Type, msg))
			//return err // TODO: uncomment when add a middleware which will be called after handlers
		}
		// TODO: Remove/comment when add a middleware which will be called after handlers
		return ctx.JSON(r)
	}
}
