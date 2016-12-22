package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

/*
 *
 * Multiplexed handlers.
 *
 */

func home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "index.gohtml", nil)
}

/*
 *
 * Administrative and internal handlers.
 *
func admin(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "admin.gohtml", nil)
}
*/

func mailHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		//Send email using appengine email api
		ctx := appengine.NewContext(req)
		req.ParseForm()
		name := req.Form.Get("name")
		email := req.Form.Get("email")
		subject := req.Form.Get("subject")
		clientMessage := req.Form.Get("message")
		msg := &mail.Message{
			Sender: name + "<" + email + ">",
			// TODO Change this email address
			To:      []string{"developers.comeback.kids@gmail.com"},
			Subject: subject,
			Body:    fmt.Sprintf(clientMessage + "\n\n" + name + "\n" + email),
		}
		if err := mail.Send(ctx, msg); err != nil {
			//An error has happened. Your email could not be sent.
			log.Errorf(ctx, "Couldn't send email: %v", err)
		}
	}
}
