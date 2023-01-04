package code290_example_send_email

import (
	"log"
	"net/smtp"
)

func SendMail() {
	username := "EMAIL_ADDR" // 로그인 이메일 주소 / 비밀번호
	password := "PASSWORD"
	from := "hyenmin@gmail.com"         // 발신자 이메일 주소
	to := []string{"hyenmin@gmail.com"} // 수신자 이메일 주소

	headerSubject := "Subject: [제목] 메일 테스트 발송\r\n" // 메일 제목 작성
	headerBlank := "\r\n"                          // 메일 형식 지정 [Content-Type: text/html;\n]
	body := "[본문] 메일 테스트 발송"                       // 메일 본문

	msg := []byte(headerSubject + headerBlank + body)
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Fatalln("Error")
		return
	}
	log.Fatalln("Success")
}
