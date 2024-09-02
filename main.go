package main

import (
	"fmt"
	"golang-boilerplate/controllers"
	"golang-boilerplate/platform/mysql"
	"golang-boilerplate/platform/yaml"
	"golang-boilerplate/repository/auth"
	"golang-boilerplate/repository/user"
	authservice "golang-boilerplate/service/auth"
	reconcilliationservice "golang-boilerplate/service/reconcilliation"
	userservice "golang-boilerplate/service/user"
	"path/filepath"
)

func main() {
	filePath, _ := filepath.Abs("config.yaml")
	conf, err := yaml.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("error open yaml file: %s", err.Error()))
	}

	db, err := mysql.Open(mysql.ConfigDB{
		Host:     conf.DataSource.MySQL.Host,
		Username: conf.DataSource.MySQL.Username,
		Password: conf.DataSource.MySQL.Password,
		DBName:   conf.DataSource.MySQL.DBName,
		Port:     conf.DataSource.MySQL.Port,
	})

	if err != nil {
		panic(fmt.Sprintf("error open database: %s", err.Error()))
	}

	userRepo := user.InitRepo(db)
	authRepo := auth.InitRepo(conf.Auth)

	userService := userservice.UserService{
		UserRepo: userRepo,
	}

	authService := authservice.AuthService{
		UserService: userService,
		AuthRepo:    authRepo,
	}

	reconService := reconcilliationservice.ReconciliationService{}

	controller := controllers.Controller{
		AuthService:  authService,
		ReconService: reconService,
		AuthConfig:   conf.Auth,
	}

	controller.Serve()
}
