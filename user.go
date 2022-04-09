package papergirl

const (
	RoleAdmin UserRole = "Admin"
	RoleUser  UserRole = "User"
)

var users Users

type (
	UserRole string
	Username string
	Password string
)

type User struct {
	id       ID
	Username Username
	Password Password
	Role     UserRole
}

func NewUser(username Username, password Password, role UserRole) User {
	return User{
		Username: username,
		Password: password,
		Role:     role,
	}
}

func (user User) ID() ID {
	return user.id
}

type Users interface {
	Repository[User]
	Count() (int, error)
}

func CreateUser(username Username, password Password, role UserRole) error {
	hashedPassword, err := passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	user := NewUser(username, hashedPassword, role)
	return users.Save(user)
}

func CountUsers() (int, error) {
	return users.Count()
}

type InMemoryUsers struct {
	*InMemoryRepository[User]
}

func NewInMemoryUsers() Users {
	repository := NewInMemoryRepository[User]()
	return &InMemoryUsers{repository}
}

func (users *InMemoryUsers) Count() (int, error) {
	return len(users.store), nil
}
