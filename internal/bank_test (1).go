package main

import "testing"

func TestAccount_Deposit(t *testing.T) {
	type fields struct {
		ID      string
		Balance float64
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add tests for success and failure cases
		// Try to think of edge cases as well (might not be many for this one but with Withdraw there probably will be)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ID:      tt.fields.ID,
				Balance: tt.fields.Balance,
			}
			if got := a.Deposit(tt.args.amount); got != tt.want {
				t.Errorf("Account.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	// TODO: create table layout similar to Deposit and then add tests for success and failure cases
	// make sure you check that the balance inside the account is correct after each test
}
