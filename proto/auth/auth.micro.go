// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	User
	Request
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
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

type AuthClient interface {
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
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Create", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Get", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAll", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllUsersRole(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllUsersRole", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Auth", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.UpdateUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.DeleteUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserMenus(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetUserMenus", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserRules(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseRule, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetUserRules", in)
	out := new(ResponseRule)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Login", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ValidateToken", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.UserToken", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.updateRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllRoles", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.deleteRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.UpdateMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllMenues(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllMenues", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.UpdateForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteForm(ctx context.Context, in *Form, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.DeleteForm", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllForms(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseForm, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllForms", in)
	out := new(ResponseForm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteFields(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.DeleteFields", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteTabs(ctx context.Context, in *Form, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.DeleteTabs", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.UpdateSchema", in)
	out := new(ResponseFormSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteSchema(ctx context.Context, in *FormSchema, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.DeleteSchema", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllSchemas(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseFormSchema, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllSchemas", in)
	out := new(ResponseFormSchema)
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
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Create(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *Auth) Get(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Get(ctx, in, out)
}

func (h *Auth) GetAll(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.GetAll(ctx, in, out)
}

func (h *Auth) GetAllUsersRole(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.GetAllUsersRole(ctx, in, out)
}

func (h *Auth) Auth(ctx context.Context, in *User, out *ResponseToken) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *Auth) UpdateUser(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.UpdateUser(ctx, in, out)
}

func (h *Auth) DeleteUser(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.DeleteUser(ctx, in, out)
}

func (h *Auth) GetUserMenus(ctx context.Context, in *User, out *ResponseMenu) error {
	return h.AuthHandler.GetUserMenus(ctx, in, out)
}

func (h *Auth) GetUserRules(ctx context.Context, in *User, out *ResponseRule) error {
	return h.AuthHandler.GetUserRules(ctx, in, out)
}

func (h *Auth) Login(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *Auth) ValidateToken(ctx context.Context, in *Token, out *ResponseToken) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *Auth) UserToken(ctx context.Context, in *Token, out *ResponseUser) error {
	return h.AuthHandler.UserToken(ctx, in, out)
}

func (h *Auth) CreateRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.CreateRole(ctx, in, out)
}

func (h *Auth) UpdateRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.UpdateRole(ctx, in, out)
}

func (h *Auth) GetRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.GetRole(ctx, in, out)
}

func (h *Auth) GetAllRoles(ctx context.Context, in *Request, out *ResponseRole) error {
	return h.AuthHandler.GetAllRoles(ctx, in, out)
}

func (h *Auth) DeleteRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.DeleteRole(ctx, in, out)
}

func (h *Auth) CreateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.CreateMenu(ctx, in, out)
}

func (h *Auth) UpdateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.UpdateMenu(ctx, in, out)
}

func (h *Auth) GetMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.GetMenu(ctx, in, out)
}

func (h *Auth) GetAllMenues(ctx context.Context, in *Request, out *ResponseMenu) error {
	return h.AuthHandler.GetAllMenues(ctx, in, out)
}

func (h *Auth) CreateForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.CreateForm(ctx, in, out)
}

func (h *Auth) GetForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.GetForm(ctx, in, out)
}

func (h *Auth) UpdateForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.UpdateForm(ctx, in, out)
}

func (h *Auth) DeleteForm(ctx context.Context, in *Form, out *ResponseForm) error {
	return h.AuthHandler.DeleteForm(ctx, in, out)
}

func (h *Auth) GetAllForms(ctx context.Context, in *Request, out *ResponseForm) error {
	return h.AuthHandler.GetAllForms(ctx, in, out)
}

func (h *Auth) DeleteFields(ctx context.Context, in *Form, out *Error) error {
	return h.AuthHandler.DeleteFields(ctx, in, out)
}

func (h *Auth) DeleteTabs(ctx context.Context, in *Form, out *Error) error {
	return h.AuthHandler.DeleteTabs(ctx, in, out)
}

func (h *Auth) CreateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.CreateSchema(ctx, in, out)
}

func (h *Auth) GetSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.GetSchema(ctx, in, out)
}

func (h *Auth) UpdateSchema(ctx context.Context, in *FormSchema, out *ResponseFormSchema) error {
	return h.AuthHandler.UpdateSchema(ctx, in, out)
}

func (h *Auth) DeleteSchema(ctx context.Context, in *FormSchema, out *Error) error {
	return h.AuthHandler.DeleteSchema(ctx, in, out)
}

func (h *Auth) GetAllSchemas(ctx context.Context, in *Request, out *ResponseFormSchema) error {
	return h.AuthHandler.GetAllSchemas(ctx, in, out)
}