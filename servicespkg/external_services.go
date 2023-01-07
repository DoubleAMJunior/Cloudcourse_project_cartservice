package servicespkg

type ExternalServices interface {
	GetUser(string) (*User, error)
	GetProduct(int32) (*Product, error)
	HasAccess(string, string, string) (bool, error)
}

type User struct {
	Id string
}

type Product struct {
	Title string
	Count int32
}
