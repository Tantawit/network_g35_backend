package router

import (
	"bytes"
	"context"
	"fmt"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

type FiberRouter struct {
	App  *fiber.App
	Chat fiber.Router
	Ws   fiber.Router
}

func (r *FiberRouter) Shutdown() error {
	return r.App.Shutdown()
}

type FiberCtx struct {
	*fiber.Ctx
}

func (c *FiberCtx) Context() context.Context {
	return c.Ctx.Context()
}

func (c *FiberCtx) UserContext() context.Context {
	return c.Ctx.UserContext()
}

func (c *FiberCtx) UserType() constant.UserType {
	userType, ok := c.Ctx.Locals("UserType").(string)
	if !ok {
		return constant.UnknownType
	}

	result, _ := strconv.Atoi(userType)
	return constant.UserType(result)
}

func (c *FiberCtx) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	return c.QueryParser(query)
}

func (c *FiberCtx) QueryParam(query any) error {
	return c.QueryParser(query)
}

func (c *FiberCtx) Query(s string) string {
	return c.Query(s)
}

func (c *FiberCtx) UserID() string {
	userId, ok := c.Ctx.Locals("UserId").(string)
	if !ok {
		return ""
	}

	return userId
}

func (c *FiberCtx) Bind(v any) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberCtx) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberCtx) ID() (id string, err error) {
	id = c.Params("id")

	_, err = uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *FiberCtx) Param(key string) (value string, err error) {
	value = c.Params(key)
	if value == "" {
		return "", errors.New("No param")
	}

	return value, nil
}

func (c *FiberCtx) Token() string {
	return c.Ctx.Get(fiber.HeaderAuthorization, "")
}

func (c *FiberCtx) Method() string {
	return c.Ctx.Method()
}

func (c *FiberCtx) Path() string {
	return c.Ctx.Path()
}

func (c *FiberCtx) StoreValue(k string, v string) {
	c.Locals(k, v)
}

func (c *FiberCtx) File(key string, allowContent map[string]struct{}, maxSize int64) (*dto.DecomposedFile, error) {
	file, err := c.Ctx.FormFile(key)
	if err != nil {
		return nil, err
	}

	if !gosdk.IsExisted(allowContent, file.Header["Content-Type"][0]) {
		return nil, errors.New("Not allow content")
	}

	if file.Size > maxSize {
		return nil, errors.New(fmt.Sprintf("Max file size is %v", maxSize))
	}
	content, err := file.Open()
	if err != nil {
		return nil, errors.New("Cannot read file")
	}

	defer content.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, content); err != nil {
		return nil, err
	}

	return &dto.DecomposedFile{
		Filename: file.Filename,
		Data:     buf.Bytes(),
	}, nil
}

func (c *FiberCtx) GetFormData(key string) string {
	return c.Ctx.FormValue(key)
}

func (c *FiberCtx) AuthInfo() *dto.AuthInfo {
	headers := c.GetReqHeaders()

	return &dto.AuthInfo{
		Hostname:  headers["Host"],
		UserAgent: headers["User-Agent"],
		IPAddress: c.IP(),
	}
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{Ctx: c}
}
