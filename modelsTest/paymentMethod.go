package modelsTest

type PaymentModel interface {
	Get() []Payment
	Insert(Payment) error
}
