package table_tests

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

var tsc = map[string]struct {
	n int
	r int
}{
	"test 1": {
		n: 2,
		r: 2,
	},
	"test 2": {
		n: 5,
		r: 120,
	},
}

func TestFactorial(t *testing.T) {
	for name, tc := range tsc {
		t.Run(name, func(t *testing.T) {
			r := factorial(tc.n)
			assert.Equal(t, r, tc.r)
		})
	}
}

func TestUseCase_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := UserModel{
		Id: "1",
	}

	mockedUserRepo := NewMockIUserRepository(ctrl)
	mockedNotificationService := NewMockINotificationService(ctrl)

	mockedUserRepo.EXPECT().FindUserById(mockUser.Id).Return(&mockUser, nil)
	mockedNotificationService.EXPECT().NotifyUser(mockUser.Id).MaxTimes(1)

	uc := &FetchUsersUseCase{
		ur: mockedUserRepo,
		ns: mockedNotificationService,
	}

	cmd := Command[FetchUserPayload]{
		payload: FetchUserPayload{
			Id: "1",
		},
	}

	result, err := uc.Handle(cmd)
	if err != nil {
		return
	}

	assert.Equal(t, result.us.Id, cmd.payload.Id)
}
