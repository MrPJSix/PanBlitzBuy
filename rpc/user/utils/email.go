package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

// GenEmailCode 生成指定位数随机验证码
func GenEmailCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {

		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendEmail(smtpHost, smtpPort, smtpUser, smtpPassword, toEmail, subject, content string) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	to := []string{toEmail}
	message := []byte(
		"From: " + smtpUser + "\r\n" +
			"To: " + strings.Join(to, ", ") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" + content + "\r\n")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, to, message)
	return err
}
