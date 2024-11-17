package partyrobot

import "testing"

func TestWelcome(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Welcome(tt.args.name); got != tt.want {
				t.Errorf("Welcome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHappyBirthday(t *testing.T) {
	type args struct {
		name string
		age  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HappyBirthday(tt.args.name, tt.args.age); got != tt.want {
				t.Errorf("HappyBirthday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignTable(t *testing.T) {
	type args struct {
		name      string
		table     int
		neighbor  string
		direction string
		distance  float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssignTable(tt.args.name, tt.args.table, tt.args.neighbor, tt.args.direction, tt.args.distance); got != tt.want {
				t.Errorf("AssignTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
