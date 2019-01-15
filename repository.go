package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/inigofu/temac-user-service/proto/auth"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	GetAllUsersRole() ([]*pb.User, error)
	Get(idcode string) (*pb.User, error)
	UpdateUser(*pb.User) error
	DeleteUser(*pb.User) error
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
	GetAllRoles() ([]*pb.Role, error)
	GetRole(idcode string) (*pb.Role, error)
	CreateRole(role *pb.Role) error
	UpdateRole(role *pb.Role) error
	DeleteRole(role *pb.Role) error
	GetAllMenues() ([]*pb.Menu, error)
	GetMenu(idcode string) (*pb.Menu, error)
	GetMenubyName(name string) (*pb.Menu, error)
	CreateMenu(menu *pb.Menu) error
	UpdateMenu(menu *pb.Menu) error
	GetUserMenus(userid string) ([]*pb.Menu, error)
	GetUserRules(userid string) ([]*pb.Rules, error)
	GetForm(name string) (*pb.Form, error)
	DeleteForm(form *pb.Form) error
	UpdateForm(form *pb.Form) (*pb.Form, error)
	CreateForm(form *pb.Form) error
	GetAllForms() ([]*pb.Form, error)
	GetSchema(idcode string) (*pb.FormSchema, error)
	CreateSchema(schema *pb.FormSchema) error
	UpdateSchema(schema *pb.FormSchema) error
	GetAllSchemas() ([]*pb.FormSchema, error)
	DeleteFields(form *pb.Form) error
	DeleteTabs(form *pb.Form) error
	DeleteSchema(form *pb.FormSchema) error
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	log.Println("Entering GetAll")
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetAllUsersRole() ([]*pb.User, error) {
	log.Println("Entering GetAllUsersRole")
	var users []*pb.User
	if err := repo.db.Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(idcode string) (*pb.User, error) {
	var user *pb.User
	user = &pb.User{Idcode: idcode}
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	log.Println("GetbyMail:", user)
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		log.Println("GetbyMail error:", err)
		return nil, err
	}
	log.Println("GetbyMail succes:", user)
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(user).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) UpdateUser(user *pb.User) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Save(user).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) DeleteUser(user *pb.User) error {
	if err := repo.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetAllRoles() ([]*pb.Role, error) {
	var roles []*pb.Role
	if err := repo.db.Preload("Menues").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (repo *UserRepository) GetRole(idcode string) (*pb.Role, error) {
	var role *pb.Role
	role = &pb.Role{Idcode: idcode}
	if err := repo.db.Preload("Menues").First(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
func (repo *UserRepository) CreateRole(role *pb.Role) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(&role).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) UpdateRole(role *pb.Role) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Save(&role).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) DeleteRole(role *pb.Role) error {
	if err := repo.db.Delete(&role).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) GetAllMenues() ([]*pb.Menu, error) {
	var menues []*pb.Menu
	if err := repo.db.Find(&menues).Error; err != nil {
		return nil, err
	}
	return menues, nil
}

func (repo *UserRepository) GetMenu(idcode string) (*pb.Menu, error) {
	var menu *pb.Menu
	menu = &pb.Menu{Idcode: idcode}
	if err := repo.db.First(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
func (repo *UserRepository) GetMenubyName(name string) (*pb.Menu, error) {
	var menu *pb.Menu
	menu = &pb.Menu{}
	if err := repo.db.Where("name = ?", name).First(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
func (repo *UserRepository) CreateMenu(menu *pb.Menu) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(menu).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) UpdateMenu(menu *pb.Menu) error {
	if err := repo.db.Set("gorm:association_autoupdate", false).Save(menu).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserMenus(email string) ([]*pb.Menu, error) {
	user := &pb.User{}
	//var roles []*pb.Role
	var menues []*pb.Menu
	var rolmenuesall []*pb.Menu

	if err := repo.db.Preload("Roles.Menues").Select("idcode").Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	log.Println("Getting menues from:", user)
	for _, role := range user.Roles {
		rolmenuesall = append(rolmenuesall, role.Menues...)
	}
	var rolmenues []string
	for _, role := range rolmenuesall {
		rolmenues = append(rolmenues, role.Idcode)
	}
	type Result struct {
		Children_idcode string
	}

	var results []Result
	var childrenid []string
	if err := repo.db.Raw("SELECT children_idcode FROM menu_childrens").Scan(&results).Error; err != nil {
		return nil, err
	}
	for _, result := range results {
		childrenid = append(childrenid, result.Children_idcode)
	}
	// (*sql.Row)
	if rolmenues != nil {
		if err := repo.db.Not(childrenid).Where(rolmenues).Preload("Children", "idcode in (?)", rolmenues).Find(&menues).Error; err != nil {

			log.Println("error getting menues:", err)
			return nil, err
		}
	}
	log.Println("Menues:", menues)
	return menues, nil
}
func (repo *UserRepository) GetUserRules(email string) ([]*pb.Rules, error) {
	log.Println("Getting user rules")
	user := &pb.User{}
	var rolrulessall []*pb.Rules

	if err := repo.db.Preload("Roles").Preload("Roles.Rules").Where("email = ?", email).
		First(&user).Error; err != nil {
		if fmt.Sprintf("%v", err) == "can't preload field Rules for auth.Role" {
			log.Println("Record not found 33:")
		}
		return nil, err
	}
	log.Println("Getting rules from:", user)
	for _, rule := range user.Roles {
		rolrulessall = append(rolrulessall, rule.Rules...)
	}
	if user.Designer {
		temprule := make([]*pb.Rules, 1)
		temprule[0] = &pb.Rules{Actions: "update", Subject: "griddesigner"}
		rolrulessall = append(rolrulessall, temprule...)
	}
	return rolrulessall, nil
}
func (repo *UserRepository) GetAllForms() ([]*pb.Form, error) {
	var forms []*pb.Form
	if err := repo.db.Preload("Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Fields.Values").Preload("Fields.Selectoptions").Preload("Tabs").Preload("Tabs.Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Tabs.Fields.Values").Preload("Tabs.Fields.Selectoptions").
		Find(&forms).Error; err != nil {
		return nil, err
	}
	return forms, nil
}

func (repo *UserRepository) GetForm(name string) (*pb.Form, error) {
	form := &pb.Form{}
	log.Println("Getting form with name:", name)
	//form = &pb.Form{Name: name}
	if err := repo.db.Preload("Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Fields.Values").Preload("Fields.Selectoptions").Preload("Tabs").
	Preload("Tabs.Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Tabs.Fields.Values").Preload("Tabs.Fields.Selectoptions").Where("name = ?", name).
		First(&form).Error; err != nil {
		return nil, err
	}
	return form, nil
}
func (repo *UserRepository) DeleteForm(form *pb.Form) error {

	if err := repo.db.Delete(pb.FormSchema{}, "form_refer = ?", form.Idcode).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&form).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) DeleteFields(form *pb.Form) error {
	var field *pb.FormSchema
	field = form.Fields[0]
	if err := repo.db.Model(&form).Association("Fields").Delete(field).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) DeleteTabs(form *pb.Form) error {

	var tab *pb.Form
	tab = form.Tabs[0]
	if err := repo.db.Model(&form).Association("Tabs").Delete(tab).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) UpdateForm(form *pb.Form) (*pb.Form, error) {
	log.Println("Updating form", form)
	if err := repo.db.Save(&form).Error; err != nil {
		return nil, err
	}
	return form, nil
}
func (repo *UserRepository) CreateForm(form *pb.Form) error {
	if err := repo.db.Create(&form).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) GetAllSchemas() ([]*pb.FormSchema, error) {
	var formschemas []*pb.FormSchema
	if err := repo.db.Preload("Values").Preload("Selectoptions").Find(&formschemas).Error; err != nil {
		return nil, err
	}
	return formschemas, nil
}

func (repo *UserRepository) GetSchema(idcode string) (*pb.FormSchema, error) {
	var formschema *pb.FormSchema
	formschema = &pb.FormSchema{Idcode: idcode}
	if err := repo.db.Preload("Values").Preload("Selectoptions").First(&formschema).Error; err != nil {
		return nil, err
	}
	return formschema, nil
}
func (repo *UserRepository) CreateSchema(formschema *pb.FormSchema) error {
	if err := repo.db.Create(&formschema).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) UpdateSchema(formschema *pb.FormSchema) error {
	if err := repo.db.Save(&formschema).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteSchema(formschema *pb.FormSchema) error {

	if formschema.FormRefer != "" {
		var form *pb.Form
		form = &pb.Form{Idcode: formschema.FormRefer}
		if err := repo.db.Find(&form).Error; err == nil {
			errtext := errors.New("this field has asocciated fomrs")
			return errtext
		}
	}
	if err := repo.db.Delete(formschema).Error; err != nil {
		return err
	}
	return nil
}
