package main

import "github.com/golang-jwt/jwt/v5"

type Role struct {
	Id   uint   `json:"id" gorm:"column:id" gorm:"primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

type User struct {
	Id       uint   `json:"id" gorm:"column:id" gorm:"primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	FullName string `json:"full_name" gorm:"column:full_name"`
	Age      uint   `json:"age" gorm:"column:age"`
	RoleId   uint   `json:"role_id" gorm:"column:role_id" gorm:"foreignkey:role_id;references:role_id"`
	Role     Role   `json:"role"`
}

type Post struct {
	Id       uint   `json:"id" gorm:"column:id" gorm:"primary_key"`
	Title    string `json:"title" gorm:"column:title"`
	Content  string `json:"content" gorm:"column:content"`
	PostedAt string `json:"posted_at" gorm:"column:posted_at"`
	AuthorId uint   `json:"author_id" gorm:"column:author_id" gorm:"foreignkey=author_id;references:id"`
	Author   User   `json:"author"`
}

type Claims struct {
	UserId           uint   `json:"user_id"`
	GrantedAuthority string `json:"granted_authority"`
	jwt.RegisteredClaims
}
