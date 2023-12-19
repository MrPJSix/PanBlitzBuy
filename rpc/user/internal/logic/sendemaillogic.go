package logic

import (
	"context"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"pan-blitz-buy/rpc/user/utils"
	"strconv"

	"pan-blitz-buy/rpc/user/internal/svc"
	"pan-blitz-buy/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *user.GeneralReq) (out *user.GeneralResp, err error) {
	out = &user.GeneralResp{}

	randCode := utils.GenEmailCode(4)
	email := in.User.Email

	l.svcCtx.CodeCache.Set(email, randCode, cache.DefaultExpiration)

	subject := "Pan-BliztBuy平台注册验证码"
	content := "您请求的验证码是" + randCode + "，有效期3分钟。"
	err = utils.SendEmail(
		l.svcCtx.Config.Email.Host,
		strconv.Itoa(l.svcCtx.Config.Email.Port),
		l.svcCtx.Config.Email.Username,
		l.svcCtx.Config.Email.Password,
		in.User.Email,
		subject,
		content,
	)
	if err != nil {
		log.Println("发送邮件失败", err)
		out.Msg = "发送邮件失败"
		out.Code = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	log.Println("发送邮件成功", in.User.Email, randCode)

	out.Code = strconv.Itoa(http.StatusOK)
	out.Msg = "发送邮件成功，请查收！"
	return
}
