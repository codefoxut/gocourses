package builder

import (
	"fmt"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func SendEmailImpl(email *email) {
	fmt.Println(email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	eb := EmailBuilder{}
	action(&eb)
	SendEmailImpl(&eb.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.
			From("abc@xyz.com").
			To("efg@xyz.com").
			Subject("[Ask] What's up?").
			Body("Hi, When can we meet to discuss further details?")
	})
}
