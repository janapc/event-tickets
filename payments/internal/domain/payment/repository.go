package payment

import "context"

type IPaymentRepository interface {
	FindByID(ctx context.Context, ID string) (*Payment, error)
	Save(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
}
