package repository

import (
	"transactions-service/internal/ports/out/dto"

	"github.com/stretchr/testify/mock"
)

type AccreditationRepositoryMock struct {
	mock.Mock
}

func (m *AccreditationRepositoryMock) CreateUserBalanceCredit(accreditation dto.AccountCreditCreation) error {
	args := m.Called(accreditation)
	return args.Error(0)
}
