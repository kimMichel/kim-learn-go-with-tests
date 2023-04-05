package counter

import (
	"sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("Incrementing the counter 3 times results in the value 3", func(t *testing.T) {
		counter := NewContador()

		counter.Incrementing()
		counter.Incrementing()
		counter.Incrementing()

		verifyContador(t, counter, 3)
	})

	t.Run("run safety with concurrency", func(t *testing.T) {
		expectCounting := 1000
		counter := NewContador()

		var wg sync.WaitGroup
		wg.Add(expectCounting)

		for i := 0; i < expectCounting; i++ {
			go func(w *sync.WaitGroup) {
				counter.Incrementing()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verifyContador(t, counter, expectCounting)
	})
}

func verifyContador(t *testing.T, result *Contador, expect int) {
	t.Helper()
	if result.Value() != expect {
		t.Errorf("result '%d', expect '%d'", result.Value(), expect)
	}
}

func NewContador() *Contador {
	return &Contador{}
}
