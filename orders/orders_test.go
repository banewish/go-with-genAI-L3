package orders

import (
	"errors"
	"testing"
)

// fakeStore — просунутий фейк (fake), що записує кожен викликання
// Exec. Це саме той тип мок-об'єкта, який мав би згенерувати ШІ-
// асистент у Розділі 2.3 — тут він наданий готовим, щоб автотести
// працювали незалежно від вашого власного промпту.
type fakeStore struct {
	calls     int
	lastQuery string
	lastArgs  []any
	err       error
}

// fakeResult — мок-реалізація Result для тестування
type fakeResult struct {
	id  string
	err error
}

func (r *fakeResult) OrderID() string {
	return r.id
}

func (r *fakeResult) Error() error {
	return r.err
}

func (f *fakeStore) Exec(query string, args ...any) (Result, error) {
	f.calls++
	f.lastQuery = query
	f.lastArgs = args

	// Витягнути orderID з аргументів (перший аргумент)
	var orderID string
	if len(args) > 0 {
		if id, ok := args[0].(string); ok {
			orderID = id
		}
	}

	return &fakeResult{id: orderID, err: f.err}, f.err
}

func TestPlaceOrder_Success(t *testing.T) {
	store := &fakeStore{}
	svc := NewOrderService(store)

	result, err := svc.PlaceOrder("order-1", 42.5)
	if err != nil {
		t.Fatalf("PlaceOrder повернув неочікувану помилку: %v", err)
	}
	if result == nil {
		t.Fatal("PlaceOrder повернув nil Result")
	}
	if result.OrderID() != "order-1" {
		t.Fatalf("Result.OrderID() = %q, очікувалось 'order-1'", result.OrderID())
	}

	if store.calls != 1 {
		t.Fatalf("Exec викликано %d разів, очікувалось 1", store.calls)
	}
	if store.lastQuery == "" {
		t.Error("Exec отримав порожній запит (query)")
	}
	if len(store.lastArgs) < 2 {
		t.Fatalf("Exec отримав %d аргументів, очікувалось щонайменше 2 (orderID, amount)", len(store.lastArgs))
	}
}

func TestPlaceOrder_PropagatesError(t *testing.T) {
	store := &fakeStore{err: errors.New("db is down")}
	svc := NewOrderService(store)

	result, err := svc.PlaceOrder("order-2", 10)
	if err == nil {
		t.Fatal("PlaceOrder мав повернути помилку, коли store.Exec її повертає")
	}
	if result == nil {
		t.Fatal("PlaceOrder повернув nil Result навіть при помилці")
	}
	if result.Error() == nil {
		t.Fatal("Result.Error() повинен повернути помилку")
	}
}

// TestOrderServiceUsesInjectedStore перевіряє, що OrderService
// дійсно використовує ІНТЕРФЕЙС, а не конкретний тип — головна
// мета рефакторингу з Розділу 2.
func TestOrderServiceUsesInjectedStore(t *testing.T) {
	var _ OrderStore = &fakeStore{} // fakeStore має задовольняти OrderStore

	store := &fakeStore{}
	svc := NewOrderService(store)
	if svc == nil {
		t.Fatal("NewOrderService повернув nil")
	}

	result, err := svc.PlaceOrder("order-3", 5)
	if err != nil {
		t.Fatalf("PlaceOrder повернув неочікувану помилку: %v", err)
	}
	if result == nil {
		t.Fatal("PlaceOrder повернув nil Result")
	}
	if store.calls == 0 {
		t.Error("PlaceOrder не викликав store.Exec — перевірте, що сервіс справді використовує injected store")
	}
}
