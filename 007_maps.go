package main

import "fmt"

type userinfo struct {
	name  string
	email string
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

func main() {
	sep("Map Declaration with var and make")
	mapDeclaration()
	sep("Map Literal Declaration")
	mapLiteral()
	sep("Map: Delete, Get, Get with boolean check")
	mapRemove()
}

func mapDeclaration() {
	// NOTE Map is a reference type just like slice
	// the below declaration is not usable. This points
	// to nil (0x0). The value needs to be allocated via
	// make
	var contacts map[string]userinfo
	fmt.Printf("contacts: %p\n", contacts)

	contacts = make(map[string]userinfo)
	contacts["A"] = userinfo{"A", "a@mail.com"}
	contacts["B"] = userinfo{"B", "b@mail.com"}
	contacts["C"] = userinfo{"C", "c@mail.com"}
	contacts["D"] = userinfo{"D", "d@mail.com"}

	fmt.Println("Ranging over with keys and values")
	for key, value := range contacts {
		fmt.Printf("%s: %+v\n", key, value)
	}

	fmt.Println("Ranging over for just keys")
	for key := range contacts {
		fmt.Printf("%+v\n", key)
	}
}

func mapLiteral() {

	contacts := map[string]userinfo{
		"Roy":  {"Roy", "roy@mail.com"},
		"Ford": {"Ford", "ford@mail.com"},
		"Jack": {"Jack", "jack@mail.com"},
	}
	fmt.Println("Ranging over with keys and values")
	for key, value := range contacts {
		fmt.Printf("%s: %+v\n", key, value)
	}

	fmt.Println("Ranging over for just keys")
	for key := range contacts {
		fmt.Printf("%+v\n", key)
	}
}

func mapRemove() {
	contacts := map[string]userinfo{
		"Roy":  {"Roy", "roy@mail.com"},
		"Ford": {"Ford", "ford@mail.com"},
		"Jack": {"Jack", "jack@mail.com"},
	}
	fmt.Println("Initial Data")
	for key, value := range contacts {
		fmt.Printf("%s: %+v\n", key, value)
	}

	delKey := "Ford"
	fmt.Println("Delete map entry with key:", delKey)
	delete(contacts, delKey)

	fmt.Printf("Check if key %v exists\n", delKey)
	u := contacts[delKey]
	fmt.Printf("Get operation on key %v resulted in value %v of type %T\n", delKey, u, u)

	u, found := contacts[delKey]
	fmt.Printf("Get operation (with boolean check) on key %v resulted found = %v of value %v\n", delKey, found, u)
}
