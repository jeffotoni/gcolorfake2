package gcolorfake2

import "testing"

func TestYellow(t *testing.T) {

     type args struct {
	msg string
     }

     tests := []struct {
	     name string
	     args args
	     want string
     }{
	//test here... code
	{"yellow_", args{"nossa cor da nossa lib yellow"}, "\033[0;33mnossa cor da nossa lib yellow\033[0m"},
     }

     for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
	   if got := Yellow(tt.args.msg); got != tt.want {
		  t.Errorf("Yellow = %v, want %v", got, tt.want)
	   }
	})
     }
}


