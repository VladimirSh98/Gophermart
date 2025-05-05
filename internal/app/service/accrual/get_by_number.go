package accrual

import "github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"

func (s Service) GetByNumber(number string) (*accrual.Calculations, error) {
	res, err := s.Client.GetByNumber(number)
	if err != nil {
		return nil, err
	}
	return res, nil
}
