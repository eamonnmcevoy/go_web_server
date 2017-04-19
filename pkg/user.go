package root

type User struct {
  Id           string  `json:"id"`
  Username     string  `json:"username"`
  Password     string  `json:"password"`
}

type UserService interface {
  Create(u *User) error
  GetByUsername(username string) (*User,error)
}