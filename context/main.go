package main

import (
	"context"
	"fmt"
	"time"
)

// 1. context.Background() - на самом высоком уровне
// 2. context.TODO -когда не уверены, какой контекст использовать
// 3. context.Value - стоит использовать как можно реже. и передавать только необязательные параметры
// 4. ctx всегда передается первым аргументом в функцию

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	parse(ctx)
}

func parse(ctx context.Context) {

	for {
		select {
		case <-time.After(time.Millisecond * 100):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("Deadline exceeded")
			return
		}
	}
}
