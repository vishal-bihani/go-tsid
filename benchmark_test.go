package tsid

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkGenerate(b *testing.B) {

	b.Run("one goroutine", func(b *testing.B) {

		tsidFactory, err := TsidFactoryBuilder().
			WithNode(1).
			Build()

		if err != nil {
			b.FailNow()
		}

		for i := 0; i < b.N; i++ {
			tsidFactory.Generate()
		}
	})

	b.Run("multiple goroutines", func(b *testing.B) {

		goroutineCount := 10
		wg := &sync.WaitGroup{}

		for i := 0; i < goroutineCount; i++ {

			wg.Add(1)
			go func(iterationCount int, wg *sync.WaitGroup) {
				defer wg.Done()
				tsidFactory, err := TsidFactoryBuilder().
					WithNode(1).
					Build()

				if err != nil {
					fmt.Errorf("Failed to instantiate tsid factory with error: %s", err)
					return
				}

				for i := 0; i < b.N; i++ {
					tsidFactory.Generate()
				}

			}(b.N, wg)
		}

		wg.Wait()
	})

}
