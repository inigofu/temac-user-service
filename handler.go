package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	pb "github.com/inigofu/temac-user-service/proto/auth"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) NewDB(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	// Create new greeter client
	//client := pb.NewAuthService("temac.auth", microclient.DefaultClient)

	var user pb.User
	ruser := &pb.ResponseUser{}
	configFile, err := os.Open("user.json")
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&user)
	log.Print("user", user)

	err = srv.Create(ctx, &user, ruser)
	if err != nil {
		log.Println("Could not create: %v", err)
	} else {
		log.Printf("Created: %s", ruser.User.Idcode)
	}
	rauth := &pb.ResponseToken{}
	log.Print("Auth user:", user)
	err = srv.Auth(ctx, &user, rauth)
	if err != nil {
		log.Fatalf("Could not auth: %v", err)
	}
	// let's just exit because
	log.Println("autg with token", rauth.Token.Token)

	var menu pb.Menu
	rmenu := &pb.ResponseMenu{}
	configFile, err = os.Open("menu.json")
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser = json.NewDecoder(configFile)
	jsonParser.Decode(&menu)
	log.Println("menu", menu)
	ctx_new := metadata.NewContext(ctx, map[string]string{
		"Authorization": rauth.Token.Token,
	})
	err = srv.CreateMenu(ctx_new, &menu, rmenu)
	if err != nil {
		log.Println("Could not create: %v", err)
		return err
	} else {
		log.Printf("Created menu: %s", rmenu)
	}

	var role pb.Role
	rrole := &pb.ResponseRole{}
	configFile, err = os.Open("role.json")
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	jsonParser = json.NewDecoder(configFile)
	jsonParser.Decode(&role)
	log.Println("role", role)

	err = srv.CreateRole(ctx_new, &role, rrole)
	if err != nil {
		log.Println("Could not create: %v", err)
		return err
	} else {
		log.Printf("Created role: %s", rrole)
	}
	temprole := make([]*pb.Role, 1)
	temprole[0] = &pb.Role{Idcode: rrole.Role.Idcode}
	user = *ruser.User
	user.Roles = temprole
	log.Printf("Updating user: %s", user)
	err = srv.UpdateUser(ctx_new, &user, ruser)
	if err != nil {
		log.Println("Could not update: %v", err)
		return err
	} else {
		log.Printf("Created use: %s", ruser)
	}
	var form []pb.Form
	rform := &pb.ResponseForm{}
	configFile, err = os.Open("form.json")
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	jsonParser = json.NewDecoder(configFile)
	jsonParser.Decode(&form)
	log.Println("form", form)
	for _, element := range form {
		err = srv.CreateForm(ctx_new, &element, rform)
		if err != nil {
			log.Println("Could not create form: %v", err)
			return err
		} else {
			log.Printf("Created form: %s", rform)
		}
	}
	log.Printf("Procedure finished")
	return nil
}
func (srv *service) GetIsNewDB(ctx context.Context, req *pb.Request, res *pb.IsNewDB) error {
	isnewdb, err := srv.repo.NewDB()
	if err != nil {
		return err
	}
	log.Println("NewDB response:", isnewdb)
	res.Isnew = isnewdb.Isnew

	return nil
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	user, err := srv.repo.Get(req.Idcode)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetUserRules(ctx context.Context, req *pb.User, res *pb.ResponseRule) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	rules, err := srv.repo.GetUserRules(req.Email)
	if err != nil {
		return err
	}
	res.Rules = rules
	return nil
}

func (srv *service) GetUserMenus(ctx context.Context, req *pb.User, res *pb.ResponseMenu) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	menues, err := srv.repo.GetUserMenus(req.Email)
	if err != nil {
		return err
	}
	res.Menues = menues
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *service) GetAllUsersRole(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	users, err := srv.repo.GetAllUsersRole()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *service) UpdateUser(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.UpdateUser(req)
	if err != nil {
		return err
	}
	res.User = req
	return nil
}
func (srv *service) DeleteUser(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteUser(req)
	if err != nil {
		return err
	}
	res.User = nil
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.ResponseToken) error {
	log.Println("Auth in with:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user, err)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = &pb.Token{Token: token}
	res.User = user
	return nil
}
func (srv *service) Login(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	token := &pb.ResponseToken{}
	err := srv.Auth(ctx, req, token)
	if err != nil {
		log.Println("User loging error:", err)
		return errors.BadRequest("temac.auth", fmt.Sprintf("User loging error:: %v", err))
	}
	md := make(metadata.Metadata)
	md["Authorization"] = token.Token.Token
	ctx = metadata.NewContext(ctx, md)
	menu := &pb.ResponseMenu{}
	err = srv.GetUserMenus(ctx, req, menu)
	if err != nil {
		return err
	}
	rule := &pb.ResponseRule{}
	err = srv.GetUserRules(ctx, req, rule)
	if err != nil {
		return err
	}
	res.Menues = menu.Menues
	res.User = token.User
	res.Token = token.Token
	res.Rules = rule.Rules
	log.Println("User loged: ", res)
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {

	log.Println("Creating user: ", req)

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error hashing password: %v", err))
	}

	tmpUser := &pb.User{}
	*tmpUser = *req
	tmpUser.Password = string(hashedPass)
	if err := srv.repo.Create(tmpUser); err != nil {
		log.Println("error creating user: %v", err)
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.tokenService.Encode(tmpUser)
	if err != nil {
		return err
	}
	res.User = tmpUser
	res.Token = &pb.Token{Token: token}

	/*
		if err := srv.Publisher.Publish(ctx, req); err != nil {
			return BadRequest("go.micro.api.example",fmt.Sprintf("error publishing event: %v", err))
		}*/

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.ResponseToken) error {
	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Idcode == "" {
		return errors.BadRequest("go.micro.api.example", "invalid user")
	}

	res.Token = &pb.Token{Valid: true}
	res.User = &pb.User{Email: claims.User.Email}

	return nil
}
func (srv *service) UserToken(ctx context.Context, req *pb.Token, res *pb.ResponseUser) error {
	log.Println("validating token") // Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Idcode == "" {
		return errors.BadRequest("go.micro.api.example", "invalid user")
	}
	md := make(metadata.Metadata)
	md["Authorization"] = req.Token
	ctx = metadata.NewContext(ctx, md)
	menu := &pb.ResponseMenu{}
	err = srv.GetUserMenus(ctx, claims.User, menu)
	if err != nil {
		return err
	}
	rule := &pb.ResponseRule{}
	err = srv.GetUserRules(ctx, claims.User, rule)
	if err != nil {
		return err
	}
	res.Menues = menu.Menues
	res.User = claims.User
	res.Token = &pb.Token{Valid: true, Token: req.Token}
	res.Rules = rule.Rules

	return nil
}
func (srv *service) CreateRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	log.Println("Creating role: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	log.Println("Creating role 1: ", req)
	for i, element := range req.Menues {
		if element.Idcode == "" {
			menu, err := srv.repo.GetMenubyName(element.Name)
			if err != nil {
				return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating role: %v", err))
			}
			req.Menues[i] = &pb.Menu{Idcode: menu.Idcode}
		} else {
			req.Menues[i] = &pb.Menu{Idcode: element.Idcode}
		}
	}
	log.Println("Creating role 2: ", req)
	if req.Rules == nil {
		temprule := make([]*pb.Rules, 1)
		temprule[0] = &pb.Rules{Actions: "read", Subject: "all"}
		req.Rules = temprule
	}
	log.Println("Creating role 3: ", req)
	if err := srv.repo.CreateRole(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating role: %v", err))
	}
	res.Role = req
	return nil
}
func (srv *service) UpdateRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	log.Println("Creating role: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.UpdateRole(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating role: %v", err))
	}

	res.Role = req
	return nil
}
func (srv *service) GetRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	role, err := srv.repo.GetRole(req.Idcode)
	if err != nil {
		return err
	}
	res.Role = role
	return nil
}
func (srv *service) DeleteRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteRole(req)
	if err != nil {
		return err
	}
	res.Role = nil
	return nil
}
func (srv *service) GetAllRoles(ctx context.Context, req *pb.Request, res *pb.ResponseRole) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	roles, err := srv.repo.GetAllRoles()
	if err != nil {
		return err
	}
	res.Roles = roles
	return nil
}

func (srv *service) CreateMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	log.Println("Creating menu: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.CreateMenu(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating menu: %v", err))
	}

	res.Menu = req
	return nil
}
func (srv *service) UpdateMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	log.Println("Updating menu: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.UpdateMenu(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error updating menu: %v", err))
	}

	res.Menu = req
	return nil
}
func (srv *service) GetMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	menu, err := srv.repo.GetMenu(req.Idcode)
	if err != nil {
		return err
	}
	res.Menu = menu
	return nil
}
func (srv *service) GetAllMenues(ctx context.Context, req *pb.Request, res *pb.ResponseMenu) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	menues, err := srv.repo.GetAllMenues()
	if err != nil {
		return err
	}
	res.Menues = menues
	return nil
}

func (srv *service) CreateForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Creating form: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.CreateForm(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating form: %v", err))
	}

	res.Form = req
	return nil
}
func (srv *service) GetForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Getting form: ", req, "with name:", req.Name)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	form, err := srv.repo.GetForm(req.Name)
	if err != nil {
		return err
	}
	res.Form = form
	return nil
}
func (srv *service) DeleteForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Getting form: ", req, "with idcode:", req.Idcode)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteForm(req)
	if err != nil {
		return err
	}
	res.Form = nil
	return nil
}
func (srv *service) UpdateForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Updating form: ", req, "with idcode:", req.Idcode)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	form, err := srv.repo.UpdateForm(req)
	if err != nil {
		return err
	}
	res.Form = form
	return nil
}
func (srv *service) GetAllForms(ctx context.Context, req *pb.Request, res *pb.ResponseForm) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	forms, err := srv.repo.GetAllForms()
	if err != nil {
		return err
	}
	res.Forms = forms
	return nil
}
func (srv *service) DeleteFields(ctx context.Context, req *pb.Form, res *pb.Error) error {
	log.Println("Deleting fields")
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteFields(req)
	if err != nil {
		return err
	}
	res = nil
	return nil
}
func (srv *service) DeleteTabs(ctx context.Context, req *pb.Form, res *pb.Error) error {
	log.Println("Deleting tabs")
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteTabs(req)
	if err != nil {
		return err
	}
	res = nil
	return nil
}

func (srv *service) CreateSchema(ctx context.Context, req *pb.FormSchema, res *pb.ResponseFormSchema) error {
	log.Println("Creating schema: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.CreateSchema(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error creating schema: %v", err))
	}

	res.Formschema = req
	return nil
}
func (srv *service) UpdateSchema(ctx context.Context, req *pb.FormSchema, res *pb.ResponseFormSchema) error {
	log.Println("Updating schema: ", req)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	if err := srv.repo.UpdateSchema(req); err != nil {
		return errors.BadRequest("go.micro.api.example", fmt.Sprintf("error updating schema: %v", err))
	}

	res.Formschema = req
	return nil
}
func (srv *service) GetSchema(ctx context.Context, req *pb.FormSchema, res *pb.ResponseFormSchema) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	schema, err := srv.repo.GetSchema(req.Idcode)
	if err != nil {
		return err
	}
	res.Formschema = schema
	return nil
}
func (srv *service) GetAllSchemas(ctx context.Context, req *pb.Request, res *pb.ResponseFormSchema) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	schemas, err := srv.repo.GetAllSchemas()
	if err != nil {
		return err
	}
	res.Formschemas = schemas
	return nil
}
func (srv *service) DeleteSchema(ctx context.Context, req *pb.FormSchema, res *pb.Error) error {
	log.Println("Deleting FormSchema")
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.BadRequest("go.micro.api.example", "no auth meta-data found in request")
	}

	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	err = srv.repo.DeleteSchema(req)
	if err != nil {
		return err
	}
	res = nil
	return nil
}
