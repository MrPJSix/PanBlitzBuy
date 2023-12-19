package logic

import (
	"context"
	"log"
	"net/http"
	"pan-blitz-buy/rpc/user/model"
	"strconv"

	"pan-blitz-buy/rpc/user/internal/svc"
	"pan-blitz-buy/rpc/user/user"
	"pan-blitz-buy/rpc/user/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterUserReq) (out *user.GeneralResp, err error) {
	out = &user.GeneralResp{}

	email := in.User.BasicInfo.Email
	trueCode, ok := l.svcCtx.CodeCache.Get(email)

	//验证码是否匹配
	if !ok || in.Code != trueCode {
		log.Println("错误的或失效的验证码")
		out.Msg = "错误的或失效验证码"
		out.Code = strconv.Itoa(http.StatusNotFound)
		return
	}

	if res := l.svcCtx.DB.Where("email = ?", email).First(&model.User{}); res.RowsAffected > 0 {
		log.Println("用户已存在", email)
		out.Msg = "用户已存在"
		out.Code = strconv.Itoa(http.StatusNotFound)
		return
	}

	hashedPass, err := utils.Encryption(in.User.BasicInfo.Password)
	if err != nil {
		log.Println("密码加密失败")
		out.Msg = "密码加密失败"
		out.Code = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	l.svcCtx.DB.Create(&model.User{
		Username:    in.User.Username,
		Email:       email,
		Password:    hashedPass,
		Description: in.User.Description,
		Status:      in.User.Status,
	})
	log.Println("注册成功", email)
	out.Msg = "恭喜您注册成功"
	out.Code = strconv.Itoa(http.StatusOK)
	return
}
