package main

import "fmt"

type printer interface {
	print()
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// NOTE 'admin' embeds 'user'
type admin struct {
	user
	level string
}

func (ad admin) print() {
	fmt.Printf("Admin: %s<%s>, level: %s\n", ad.user.name, ad.user.email, ad.level)
}

func (u user) notify() {
	fmt.Printf("User %s has been notified.\n", u.name)
}

func main() {
	sep("Example to demonstrate how embedding exposes the inner types methods directly")
	typeEmbedding()
}

func typeEmbedding() {
	ad := admin{
		user: user{
			name:  "Bob",
			email: "bob@admin.noreply.com",
		},
		level: "superuser",
	}
	fmt.Println("admin implements printer. Calling its print method via admin value")
	ad.print()
	fmt.Println("admin embeds user. Calling user method from admin (Indirect invocation):")
	ad.user.notify()
	fmt.Println("admin embeds user. Calling user method from admin (Direct invocation):")
	ad.notify()
}

func sep(heading string) {
	var heading_len int
	heading_len = len(heading)
	if len(heading) == 0 {
		heading_len = 25
	}

	for i := 0; i < heading_len; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Println(heading)
	for i := 0; i < heading_len; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
