

func Promise(executor func() interface{}) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve, reject := args[0], args[1]
		go executor()
		return nil
	})
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}
