package user

import (
	"context"
	"log"
	"net/http"
	"pan-blitz-buy/api/rpcclients/userclient/user"
	"pan-blitz-buy/api/rpcclients/userclient/userservice"
	"pan-blitz-buy/api/utils"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"
)

type UserEmailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEmailSendLogic {
	return &UserEmailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEmailSendLogic) UserEmailSend(req *types.UserEmailSendReq) (resp *types.UserEmailSendResp, err error) {
	resp = &types.UserEmailSendResp{}

	//邮件地址格式验证
	if !utils.VerifyEmailFormat(req.Email) {
		log.Println("邮箱格式不正确", req.Email)
		resp.Code = http.StatusNotFound
		resp.Msg = "邮箱格式不正确"
		return
	}

	r, err := l.svcCtx.UserService.SendEmail(l.ctx, &userservice.GeneralReq{
		User: &user.BasicUserInfo{
			Email: req.Email,
		},
	})

	resp.Code, _ = strconv.Atoi(r.Code)
	resp.Msg = r.Msg

	if err != nil {
		log.Println("邮箱发送失败", err)
		return
	}

	log.Println("发送邮件成功")

	return
}
