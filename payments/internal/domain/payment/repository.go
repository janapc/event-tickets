package payment

type IPaymentRepository interface {
	FindByID(ID string) (*Payment, error)
	Save(payment *Payment) error
	Update(payment *Payment) error
}
