package __3_constructors

type UsersRepo interface {
	CreateUser(id string, user any) (any, error)
	UpdateUser(id string, user any) (any, error)
	ReadUser(id string) (any, error)
	DeleteUser(id string) (any, error)
}

//[idea tip] implement interface: cmd + I
//[idea tip] generate constructor: option + enter
type userRepoImpl struct {
}

func newUserRepoImpl() UsersRepo {
	return &userRepoImpl{} // notice this method returns UsersRepo and not return *userRepoImpl
}

// try deleting one of the methods below
// idea tip [implement missing methods]

func (u userRepoImpl) CreateUser(id string, user any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoImpl) UpdateUser(id string, user any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoImpl) ReadUser(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoImpl) DeleteUser(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}
