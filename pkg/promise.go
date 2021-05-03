

import "syscall/js"

func Promise(executor func() interface{}) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve, reject := args[0], args[1]
		go executor()
		return nil
	})
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func fib(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}