// net/smtpパッケージとnet/mailパッケージによってメールの送信を行うことができる
// enmimeで添付ファイルなども送信可能であるが、UTF-8エンコーディングで送信されてしまうため、文字化けする可能性もある
package main

import (
	"log"
	"net/smtp"

	"github.com/jhillyerd/enmime"
)

func main() {
	smtpHost := "my-mail-server:25"
	smtpAuth := smtp.PlainAuth(
		"example.com",
		"example-user",
		"example-password",
		"auth.example.com",
	)

	sender := enmime.NewSMTP(smtpHost, smtpAuth)

	master := enmime.Builder().
		From("送信太郎", "taro@example.com").
		Subject("メールのタイトル").
		Text([]byte("テキストメールの本文")).
		HTML([]byte("<p>HTML メール<b>本文</b></p>")).
		AddFileAttachment("document.pdf")

	msg := master.To("宛先花子", "hanako@example.com")
	err := msg.Send(sender)
	if err != nil {
		log.Fatal(err)
	}
}
