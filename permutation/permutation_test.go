package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPermutationConcurrency(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name  string
		args1 args
		want  [][]byte
	}{
		// TODO: Add test cases.
		{"3 args", args{s: []byte{'1', '2', '3'}}, [][]byte{{'1', '2', '3'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermutationConcurrency(tt.args1.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermutationConcurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
		close(ch1)

	}()
	fmt.Println("abcd")
	time.Sleep(5 * time.Second)
	<-ch1
}
