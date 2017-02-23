// Package user provides functionalities to handle Users in the context of kontainer.io
package user

// The Service interface describes the function necessary for kontainer.io user handling
type Service interface {
	// CreateUser creates a new User and returns its id
	CreateUser(username string, cfg *Config, adr *Address) (uint, error)

	// EditUser is used to alter user information by id
	EditUser(id int, cfg *Config) error

	// ChangeUsername is used to change a users username by id
	ChangeUsername(id int, username string) error

	// DeleteUser is used to remove a user by id
	DeleteUser(id int) error

	// ResetPassword is used to reset a users password and issue a reset Mail
	ResetPassword(email string) error

	// GetUser is used to gather a users data set by id
	GetUser(id int, user *User) error
}

type dbAdapter interface {
	GetValue() interface{}
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
}

type service struct {
	db dbAdapter
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Address{}, &User{}, &Customer{})
}

func (s *service) CreateUser(username string, cfg *Config, adr *Address) (uint, error) {
	// TODO: implement functionality
	return 0, nil
}

func (s *service) EditUser(id int, cfg *Config) error {
	// TODO: implement functionality
	return nil
}

func (s *service) ChangeUsername(id int, username string) error {
	// TODO: implement functionality
	return nil
}

func (s *service) DeleteUser(id int) error {
	// TODO: implement functionality
	return nil
}

func (s *service) ResetPassword(email string) error {
	// TODO: implement functionality
	return nil
}

func (s *service) GetUser(id int, user *User) error {
	// TODO: implement functionality
	return nil
}

// NewService creates a UserService with necessary dependencies.
func NewService(db dbAdapter) (Service, error) {
	s := &service{
		db: db,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return nil, err
	}

	return s, nil
}