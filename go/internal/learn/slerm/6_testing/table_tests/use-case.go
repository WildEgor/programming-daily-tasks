package table_tests

//go:generate mockgen -source=use-case.go -destination=mocks.go -package=table_tests

type Command[T interface{}] struct {
	payload T
}

type IUseCase[T any, R any] interface {
	Handle(cmd Command[T]) (R, error)
}

type UserModel struct {
	Id   string
	Name string
}

type IUserRepository interface {
	FindUserById(id string) (*UserModel, error)
}

type INotificationService interface {
	NotifyUser(id string)
}

type FetchUserPayload struct {
	Id string
}

type FetchUserResult struct {
	us *UserModel
}

type FetchUsersUseCase struct {
	ur IUserRepository
	ns INotificationService
}

func (uc *FetchUsersUseCase) Handle(cmd Command[FetchUserPayload]) (*FetchUserResult, error) {
	us, err := uc.ur.FindUserById(cmd.payload.Id)
	if err != nil {
		return nil, err
	}

	uc.ns.NotifyUser(us.Id)

	return &FetchUserResult{
		us,
	}, nil
}
