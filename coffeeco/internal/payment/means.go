package payment

type Means string

const (
	MEANS_CASH = "cash"
	MEANS_CARD = "card"
	MEANS_COFFEEBUX = "coffeebux"
)

type CardDetails struct {
	cardToken string
}