package entity

type User interface {
	GetId() string
	GetUserName() string
	GetPassword() string
	GetHashedPassword() string
	SetHashedPassword(string)
	GetRole() string
}

func NewEmptyUser() *user {
	return &user{}
}

func NewUserFromEntity(
	entity User,
) *user {
	return &user{
		ID:             entity.GetId(),
		Username:       entity.GetUserName(),
		Password:       entity.GetHashedPassword(),
		HashedPassword: entity.GetHashedPassword(),
		Role:           entity.GetRole(),
	}
}

type user struct {
	ID             string
	Username       string
	Password       string
	HashedPassword string
	Role           string
}

func (s *user) GetId() string {
	return s.ID
}

func (s *user) GetUserName() string {
	return s.Username
}

func (s *user) SetUserName(name string) {
	s.Username = name
}

func (s *user) GetPassword() string {
	return s.Password
}

func (s *user) GetHashedPassword() string {
	return s.HashedPassword
}

func (s *user) SetHashedPassword(password string) {
	s.HashedPassword = password
}

func (s *user) SetDescription(password string) {
	s.Password = password
}

func (s *user) GetRole() string {
	return s.Role
}

func (s *user) SetRole(role string) {
	s.Role = role
}
