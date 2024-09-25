package jwt

import(
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const(
	JWT_SECRET="123456"
	COOKIE_NAME="Auth"
)

type JwtHeader struct{
	Algo  string  `json:"alg"`
	Type  string  `json:"typ"`
}

var (
	DefaultHeader=JwtHeader{
		Algo:"HS256",
		Type:"JWT",
	}
)

type JwtPayload struct{
	ID          string         `json:"jti"`
	Issue       string         `json:"iss"`
	Audience    string         `json:"aud"`
    Subject     string         `json:"sub"`
	IssueAt     int64          `json:"iat"`
	NotBefore   int64          `json:"nbf"`
	Expiration  int64          `json:"exp"`
	UserDefined map[string]any `json:"ud"`
}

func GenJWT(header JwtHeader,payload JwtPayload,secret string)(string,error){
	var part1,part2,signature string
	//header转成json，进行Base64编码
	if bs1,err:=json.Marshal(header);err!=nil{
		return "",err
	}else{
		part1=base64.RawURLEncoding.EncodeToString(bs1)
	}
	if bs2,err:=json.Marshal(payload);err!=nil{
		return "",err
	}else{
		part2=base64.RawURLEncoding.EncodeToString(bs2)
	}
	//基于sha256的哈希认证算法。任意长度的字符串，经过sha256之后长度都变成256bits
	h:=hmac.New(sha256.New,[]byte(secret))
	h.Write([]byte(part1+"."+part2))
	signature=base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return part1+"."+part2+"."+signature,nil
}

func Verifyjwt(token string,secret string)(*JwtHeader,*JwtPayload,error){
	parts:=strings.Split(token,".")
	if len(parts)!=3{
		return nil,nil,fmt.Errorf("token是%d部分",len(parts))
	}
	h:=hmac.New(sha256.New,[]byte(secret))
	h.Write([]byte(parts[0]+"."+parts[1]))
	signature:=base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature!=parts[2]{
		return nil,nil,fmt.Errorf("验证失败")
	}


	var err error
	var parts1,parts2 []byte
	if parts1,err=base64.RawURLEncoding.DecodeString(parts[0]);err!=nil{
		return nil,nil,fmt.Errorf("header Base64反解失败")
	}
	if parts2,err=base64.RawURLEncoding.DecodeString(parts[0]);err!=nil{
		return nil,nil,fmt.Errorf("payload Base64反解失败")
	}

	var header JwtHeader
	var payload JwtPayload
	if err=json.Unmarshal(parts1,&header);err!=nil{
		return nil,nil,fmt.Errorf("header json反解失败")
	}
	if err=json.Unmarshal(parts2,&payload);err!=nil{
		return nil,nil,fmt.Errorf("header json反解失败")
	}
	return &header,&payload,nil
}
/*
func TestJWT(t *testing.T){
	secret:="123456"
	header:=DefaultHeader
	payload:=JwtPayload{

	}

	if token,err:=GenJWT(header,payload,secret);err!=nil{
		fmt.Printf("生成json web token失败:%v",err)
	} else {
		fmt.Println(token)
		if -,p,err:=VerifyJwt(tokrn,secret);err!=nil{
			fmt.Println(err)
		}else{
			fmt.Printf("JWT验证通过。欢迎%s!\n",p.UserDefined["name"])
		}
	}
}*/