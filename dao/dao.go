package dao

import (
	"blog/model"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface{
	Register(user *model.User)
	Login(username string) model.User

	//博客操作
	AddPost(post *model.Post)//添加一篇文章
	GetAllPost() []model.Post//获得所有文章，返回切片
	getPost(pid int) model.Post
}

type manager struct{
	db *gorm.DB
}

var Mgr *manager

func init(){
	dsn:="root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("Failed to init db",err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func(mgr *manager) Register(user *model.User){
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User{
	var user model.User
	mgr.db.Where("username=?",username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *model.Post){
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post{
	var posts = make([]model.Post,10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) model.Post{
	var post model.Post
	mgr.db.First(&post,pid)
	return post
}
