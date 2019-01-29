// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	User
	Request
	IsNewDB
	ResponseUser
	ResponseRole
	ResponseMenu
	ResponseRule
	ResponseForm
	ResponseFormSchema
	ResponseToken
	Token
	Error
	Role
	Rules
	Conditions
	Menu
	Badge
	Wrapper
	Atributes
	Form
	FormSchema
	Buttons
	Class
	SelectOptions
	Values
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error)
	GetAllUsersRole(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseToken, error)
	UpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	DeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	GetUserMenus(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseMenu, error)
	GetUserRules(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseRule, error)
	Login(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseToken, error)
	UserToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseUser, error)
	CreateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	UpdateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	GetAllRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseRole, error)
	DeleteRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	CreateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error)
	UpdateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error)
	GetMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error)
	GetAllMenues(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseMenu, error)
	CreateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error)
	GetForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error)
	UpdateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error)
	DeleteForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error)
	GetAllForms(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseForm, error)
	DeleteFields(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error)
	DeleteTabs(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error)
	CreateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error)
	GetSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error)
	UpdateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error)
	DeleteSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*Error, error)
	GetAllSchemas(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseFormSchema, error)
	NewDB(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error)
	GetIsNewDB(ctx context.Context, in *Request, opts ...client.CallOption) (*IsNewDB, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.Create", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.Get", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAll", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAllUsersRole(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAllUsersRole", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.name, "Auth.Auth", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetUserMenus(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.name, "Auth.GetUserMenus", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetUserRules(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseRule, error) {
	req := c.c.NewRequest(c.name, "Auth.GetUserRules", in)
	out := new(ResponseRule)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Login(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.Login", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateToken", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UserToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.UserToken", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.name, "Auth.updateRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.name, "Auth.GetRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAllRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAllRoles", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.name, "Auth.deleteRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.name, "Auth.GetMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAllMenues(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAllMenues", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.name, "Auth.GetForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAllForms(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAllForms", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteFields(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteFields", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteTabs(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteTabs", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.name, "Auth.GetSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteSchema", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAllSchemas(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.name, "Auth.GetAllSchemas", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) NewDB(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Auth.NewDB", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetIsNewDB(ctx context.Context, in *Request, opts ...client.CallOption) (*IsNewDB, error) {
	req := c.c.NewRequest(c.name, "Auth.GetIsNewDB", in)
	out := new(IsNewDB)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Create(context.Context, *User, *ResponseUser) error
	Get(context.Context, *User, *ResponseUser) error
	GetAll(context.Context, *Request, *ResponseUser) error
	GetAllUsersRole(context.Context, *Request, *ResponseUser) error
	Auth(context.Context, *User, *ResponseToken) error
	UpdateUser(context.Context, *User, *ResponseUser) error
	DeleteUser(context.Context, *User, *ResponseUser) error
	GetUserMenus(context.Context, *User, *ResponseMenu) error
	GetUserRules(context.Context, *User, *ResponseRule) error
	Login(context.Context, *User, *ResponseUser) error
	ValidateToken(context.Context, *Token, *ResponseToken) error
	UserToken(context.Context, *Token, *ResponseUser) error
	CreateRole(context.Context, *Role, *ResponseRole) error
	UpdateRole(context.Context, *Role, *ResponseRole) error
	GetRole(context.Context, *Role, *ResponseRole) error
	GetAllRoles(context.Context, *Request, *ResponseRole) error
	DeleteRole(context.Context, *Role, *ResponseRole) error
	CreateMenu(context.Context, *Menu, *ResponseMenu) error
	UpdateMenu(context.Context, *Menu, *ResponseMenu) error
	GetMenu(context.Context, *Menu, *ResponseMenu) error
	GetAllMenues(context.Context, *Request, *ResponseMenu) error
	CreateForm(context.Context, *Form, *ResponseForm) error
	GetForm(context.Context, *Form, *ResponseForm) error
	UpdateForm(context.Context, *Form, *ResponseForm) error
	DeleteForm(context.Context, *Form, *ResponseForm) error
	GetAllForms(context.Context, *Request, *ResponseForm) error
	DeleteFields(context.Context, *Form, *Error) error
	DeleteTabs(context.Context, *Form, *Error) error
	CreateSchema(context.Context, *FormSchema, *ResponseFormSchema) error
	GetSchema(context.Context, *FormSchema, *ResponseFormSchema) error
	UpdateSchema(context.Context, *FormSchema, *ResponseFormSchema) error
	DeleteSchema(context.Context, *FormSchema, *Error) error
	GetAllSchemas(context.Context, *Request, *ResponseFormSchema) error
	NewDB(context.Context, *Request, *ResponseUser) error
	GetIsNewDB(context.Context, *Request, *IsNewDB) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Create(ctx context.Context, in *User, out *ResponseUser) error
		Get(ctx context.Context, in *User, out *ResponseUser) error
		GetAll(ctx context.Context, in *Request, out *ResponseUser) error
		GetAllUsersRole(ctx context.Context, in *Request, out *ResponseUser) error
		Auth(ctx context.Context, in *User, out *ResponseToken) error
		UpdateUser(ctx context.Context, in *User, out *ResponseUser) error
		DeleteUser(ctx context.Context, in *User, out *ResponseUser) error
		GetUserMenus(ctx context.Context, in *User, out *ResponseMenu) error
		GetUserRules(ctx context.Context, in *User, out *ResponseRule) error
		Login(ctx context.Context, in *User, out *ResponseUser) error
		ValidateToken(ctx context.Context, in *Token, out *ResponseToken) error
		UserToken(ctx context.Context, in *Token, out *ResponseUser) error
		CreateRole(ctx context.Context, in *Role, out *ResponseRole) error
		UpdateRole(ctx context.Context, in *Role, out *ResponseRole) error
		GetRole(ctx context.Context, in *Role, out *ResponseRole) error
		GetAllRoles(ctx context.Context, in *Request, out *ResponseRole) error
		DeleteRole(ctx context.Context, in *Role, out *ResponseRole) error
		CreateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error
		UpdateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error
		GetMenu(ctx context.Context, in *Menu, out *ResponseMenu) error
		GetAllMenues(ctx context.Context, in *Request, out *ResponseMenu) error
		CreateForm(ctx context.Context, in *Form, out *ResponseForm) error
		GetForm(ctx context.Context, in *Form, out *ResponseForm) error
		UpdateForm(ctx context.Context, in *Form, out *ResponseForm) error
		DeleteForm(ctx context.Context, in *Form, out *ResponseForm) error
		GetAllForms(ctx context.Context, in *Request, out *ResponseForm) error
		DeleteFields(ctx context.Context, in *Form, out *Error) error
		DeleteTabs(ctx context.Context, in *Form, out *Error) error
		CreateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error
		GetSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error
		UpdateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error
		DeleteSchema(ctx context.Context, in *FormSchema, out *Error) error
		GetAllSchemas(ctx context.Context, in *Request, out *ResponseFormSchema) error
		NewDB(ctx context.Context, in *Request, out *ResponseUser) error
		GetIsNewDB(ctx context.Context, in *Request, out *IsNewDB) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Create(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *authHandler) Get(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Get(ctx, in, out)
}

func (h *authHandler) GetAll(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.GetAll(ctx, in, out)
}

func (h *authHandler) GetAllUsersRole(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.GetAllUsersRole(ctx, in, out)
}

func (h *authHandler) Auth(ctx context.Context, in *User, out *ResponseToken) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *authHandler) UpdateUser(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.UpdateUser(ctx, in, out)
}

func (h *authHandler) DeleteUser(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.DeleteUser(ctx, in, out)
}

func (h *authHandler) GetUserMenus(ctx context.Context, in *User, out *ResponseMenu) error {
	return h.AuthHandler.GetUserMenus(ctx, in, out)
}

func (h *authHandler) GetUserRules(ctx context.Context, in *User, out *ResponseRule) error {
	return h.AuthHandler.GetUserRules(ctx, in, out)
}

func (h *authHandler) Login(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *authHandler) ValidateToken(ctx context.Context, in *Token, out *ResponseToken) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *authHandler) UserToken(ctx context.Context, in *Token, out *ResponseUser) error {
	return h.AuthHandler.UserToken(ctx, in, out)
}

func (h *authHandler) CreateRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.CreateRole(ctx, in, out)
}

func (h *authHandler) UpdateRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.UpdateRole(ctx, in, out)
}

func (h *authHandler) GetRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.GetRole(ctx, in, out)
}

func (h *authHandler) GetAllRoles(ctx context.Context, in *Request, out *ResponseRole) error {
	return h.AuthHandler.GetAllRoles(ctx, in, out)
}

func (h *authHandler) DeleteRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.DeleteRole(ctx, in, out)
}

func (h *authHandler) CreateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.CreateMenu(ctx, in, out)
}

func (h *authHandler) UpdateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.UpdateMenu(ctx, in, out)
}

func (h *authHandler) GetMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.GetMenu(ctx, in, out)
}

func (h *authHandler) GetAllMenues(ctx context.Context, in *Request, out *ResponseMenu) error {
	return h.AuthHandler.GetAllMenues(ctx, in, out)
}

func (h *authHandler) CreateForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.CreateForm(ctx, in, out)
}

func (h *authHandler) GetForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.GetForm(ctx, in, out)
}

func (h *authHandler) UpdateForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.UpdateForm(ctx, in, out)
}

func (h *authHandler) DeleteForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.DeleteForm(ctx, in, out)
}

func (h *authHandler) GetAllForms(ctx context.Context, in *Request, out *ResponseForm) error {
	return h.AuthHandler.GetAllForms(ctx, in, out)
}

func (h *authHandler) DeleteFields(ctx context.Context, in *Form, out *Error) error {
	return h.AuthHandler.DeleteFields(ctx, in, out)
}

func (h *authHandler) DeleteTabs(ctx context.Context, in *Form, out *Error) error {
	return h.AuthHandler.DeleteTabs(ctx, in, out)
}

func (h *authHandler) CreateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.CreateSchema(ctx, in, out)
}

func (h *authHandler) GetSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.GetSchema(ctx, in, out)
}

func (h *authHandler) UpdateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.UpdateSchema(ctx, in, out)
}

func (h *authHandler) DeleteSchema(ctx context.Context, in *FormSchema, out *Error) error {
	return h.AuthHandler.DeleteSchema(ctx, in, out)
}

func (h *authHandler) GetAllSchemas(ctx context.Context, in *Request, out *ResponseFormSchema) error {
	return h.AuthHandler.GetAllSchemas(ctx, in, out)
}

func (h *authHandler) NewDB(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.NewDB(ctx, in, out)
}

func (h *authHandler) GetIsNewDB(ctx context.Context, in *Request, out *IsNewDB) error {
	return h.AuthHandler.GetIsNewDB(ctx, in, out)
}
