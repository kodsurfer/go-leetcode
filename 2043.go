type Bank []int64


func Constructor(balance []int64) Bank {
    return balance
}


func (this Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1>len(this) || account2>len(this) || this[account1-1]<money {
        return false
    }
    this[account1-1]-=money
    this[account2-1]+=money
    return true
    }

func (this Bank) Deposit(account int, money int64) bool {
    if account>len(this) {
        return false
    }
    this[account-1]+=money
    return true
}


func (this Bank) Withdraw(account int, money int64) bool {
    if account>len(this) || this[account-1]<money {
        return false
    }
    this[account-1]-=money
    return true
}

