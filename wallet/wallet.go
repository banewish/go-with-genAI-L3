// Package wallet реалізує Розділ 3, Крок 1 домашньої роботи:
// діагностику та виправлення класичної пастки Go — змішування
// приймача-значення і приймача-вказівника на одному типі.
package wallet

// SecureWallet — той самий тип, що й у прикладі з домашнього
// завдання. Balance() навмисно залишено з приймачем-значенням, щоб
// ви могли побачити (і виправити) саме ту невідповідність, яку ми
// розбирали на занятті.
type SecureWallet struct {
	balance float64
}

// Balance повертає поточний баланс. Використовує приймач-вказівник
// для консистентності з Deposit() — обидва методи мають той же приймач.
func (w *SecureWallet) Balance() float64 {
	return w.balance
}

func (w *SecureWallet) Deposit(amt float64) {
	w.balance += amt
}

// ApplyDeposits додає amt до кожного гаманця у зрізі wallets.
//
// Це — саме той сценарій із заняття: wallets це []SecureWallet
// (зріс ЗНАЧЕНЬ, а не вказівників). Наївний `for _, w := range
// wallets { w.Deposit(amt) }` НЕ подіє, бо w — копія елемента циклу.
//
// Рішення: індексний for-цикл (for i := range wallets) дозволяє
// звертатися до wallets[i] напряму. Go автоматично перетворює
// wallets[i].Deposit(amt) в (&wallets[i]).Deposit(amt), беручи адресу
// справжнього елемента зрізу, а не копії.
func ApplyDeposits(wallets []SecureWallet, amt float64) {
	for i := range wallets {
		wallets[i].Deposit(amt)
	}
}
