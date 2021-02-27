package user

type User struct {
	Name string // Exported
	ID   int    // Exported

	password string // un-exported
}
