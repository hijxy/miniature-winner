package controller

import (
	"fmt"
	/*"strconv"*/
	"blog/dao"
	"blog/model"
	"blog/jwt"
	"strings"
	/*"html/template"
	"github.com/russross/blackfriday"*/
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	password2:=c.PostForm("password2")
	if password!=password2{
		c.HTML(200,"register.html","两次密码输入不一致，请再次输入")
	}else{
		if len(password)<5{
		c.HTML(200,"register.html","密码长度小于5，请重新输入")}
	}
	user:=model.User{
		Username:username,
		Password:password,
	}
	dao.Mgr.Register(&user)
	c.Redirect(301 ,"/login")
}

func GoRegister(c *gin.Context){
	c.HTML(200,"register.html",nil)
}

func GoLogin(c *gin.Context){
	c.HTML(200,"login.html",nil)
}

func Login(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	u:=dao.Mgr.Login(username)
	if u.Username==""{
		c.HTML(200,"login.html","用户名不存在，再试一次吧")
	}else{
		if u.Password!=password{
			fmt.Println("密码错误")
			c.HTML(200,"login.html","密码错误,再试一次吧")
		}
		uid:="8"
			payload:=jwt.JwtPayload{UserDefined:map[string]any{"uid":uid}}
			token,_:=jwt.GenJWT(jwt.DefaultHeader,payload,jwt.JWT_SECRET)
			fmt.Println("jwt token",token)
			c.SetCookie(
				jwt.COOKIE_NAME,
				token,
				86400*1,
				"/",
				"localhost",
				false,
				true,
			)
			fmt.Println("登录成功")
			c.Redirect(301,"/")
		
	    }
    }

func getUidfromCookie1(c *gin.Context) string{
	for _,cookie:=range strings.Split(c.Request.Header.Get("cookie"),";"){
		arr:=strings.Split(cookie,"=")
		key:=strings.TrimSpace(arr[0])
		value:=strings.TrimSpace(arr[1])
		if key ==jwt.COOKIE_NAME{
			if _,payload,err:=jwt.Verifyjwt(value,jwt.JWT_SECRET);err==nil{
				if uid,exists:=payload.UserDefined["uid"];exists{
					return uid.(string)
				}
			}
		}
	}
	return ""
}



func Index(c *gin.Context){
	c.HTML(200,"index.html",nil)
}

func ListUser(c *gin.Context){
	c.HTML(200,"userlist.html",nil)
}

//操作博客

//博客列表
func GetPostIndex(c *gin.Context){
	posts:=dao.Mgr.GetAllPost()
	c.HTML(200,"postIndex.html",posts)
}

//添加博客
func AddPost(c *gin.Context){
	title:=c.PostForm("title")
	tag:=c.PostForm("tag")
	content:=c.PostForm("content")
	post:=model.Post{
		Title: title,
		Tag: tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)
	c.Redirect(302,"/post_index")
}

//跳转到添加博客
func GoAddPost(c *gin.Context){
	c.HTML(200,"post.html",nil)
}

func GoPricing(c *gin.Context){
	c.HTML(200,"pricing.html",nil)
}


/*func PostDetail(c *gin.Context){
	s:=c.Query("pid")
	pid,_:=strconv.Atoi(s)
	p:=dao.Mgr.GetPost(pid)
	md := blackfriday.New()
	content:=md.Run([]byte(p.Content))
	
	c.HTML(200,"detail.html",gin.H{
		"Title":p.Title,
		"Content":template.HTML(content),
	})
}
	blackfriday导入失败，待解决*/



 