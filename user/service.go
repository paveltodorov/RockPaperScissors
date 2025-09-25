package user

const (
	DefaultBalance = 100
)

type Repository interface {
	AddUser(u *User) int
	FindUserByUsername(username string) (*User, bool)
	FindUserByID(id int) (*User, bool)
	ListUsers() []*User
	ListAIUsers() []*User
	UpdateUser(u *User) error
}

type Service struct {
	repo Repository
}

func (s *Service) ListAIUsers() []*User {
	return s.repo.ListAIUsers()
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Login(username, password string) (*User, error) {
	if err := s.validateCredentials(username, password); err != nil {
		return nil, err
	}

	if u, ok := s.repo.FindUserByUsername(username); ok {
		if u.Password == password {
			return u, nil
		}
		return nil, InvalidCredentialsError{Username: username}
	}

	// Create new user
	u := &User{Username: username, Password: password, Balance: DefaultBalance, Strategy: "", IsAI: false}
	userID := s.repo.AddUser(u)
	u.ID = userID
	return u, nil
}

func (s *Service) AddUser(u *User) int {
	return s.repo.AddUser(u)
}

func (s *Service) validateCredentials(username, password string) error {
	if len(username) < 3 || len(username) > 20 {
		return InvalidUsernameError{Username: username}
	}

	// Check for alphanumeric characters only
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return InvalidUsernameError{Username: username}
		}
	}

	// Validate password
	if len(password) < 4 {
		return InvalidPasswordError{Message: "password must be at least 4 characters"}
	}
	if len(password) > 50 {
		return InvalidPasswordError{Message: "password cannot exceed 50 characters"}
	}

	return nil
}

func (s *Service) GetByID(id int) (*User, error) {
	if id <= 0 {
		return nil, UserNotFoundError{UserID: id}
	}

	if u, ok := s.repo.FindUserByID(id); ok {
		return u, nil
	}
	return nil, UserNotFoundError{UserID: id}
}

func (s *Service) List() []*User {
	return s.repo.ListUsers()
}

func (s *Service) Update(u *User) error {
	return s.repo.UpdateUser(u)
}
