package repository

import (
	models "go-admin/modules/primitive/model"
	"gorm.io/gorm"
)

const (
	IsActive   = iota + 1
	ParentMenu = "parent_menu"
	Menu       = "menu"
)

func GetListParentMenu(db *gorm.DB) ([]models.Menu, error) {
	var listParentMenu []models.Menu
	err := db.Raw("select * from web_menu where is_active = ? and menu_type = ? ",
		IsActive, ParentMenu).Scan(&listParentMenu).Error
	if err != nil {
		return nil, err
	}
	return listParentMenu, nil
}

func GetMenu(db *gorm.DB, parentID int) ([]models.Menu, error) {
	var menus []models.Menu
	err := db.Raw("select * from web_menu where parent_id = ? and is_active = ? ",
		parentID, IsActive).Scan(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func GetMenuWithParentID(db *gorm.DB, parentID int) ([]models.Menu, error) {
	var menus []models.Menu
	err := db.Raw("select * from web_menu where parent_id = ? and is_active = ? and menu_type = ? ",
		parentID, IsActive, Menu).Scan(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func GetMenuIsActive(db *gorm.DB) ([]models.Menu, error) {
	var menus []models.Menu
	err := db.Raw("select * from web_menu where is_active = ? ",
		IsActive).Scan(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
