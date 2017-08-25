package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan drawResult)

type drawResult struct {
	result chan bool
	amount int
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- drawResult{ch, amount}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withdraws:
			if balance < draw.amount {
				draw.result <- false
			} else {
				balance -= draw.amount
				draw.result <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
