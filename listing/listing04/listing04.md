Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
print: 0-9
fatal error: all goroutines are asleep - deadlock!
```

Range продолжает ждать сообщения из незакрытого канала. 