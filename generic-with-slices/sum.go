package main

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Sum(numbers []int) (sum int) {

	add := func(acc, curr int) int {
		acc += curr
		return acc
	}

	return Reduce(numbers, add, 0)
}

func SumAllTails(numbers ...[]int) []int {

	sumTail := func(acc, curr []int) []int {
		if len(curr) == 0 {
			return append(acc, 0)
		} else {
			tail := curr[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
