package recovery

func SafeGoroutine(f func(), ctx string, values ...interface{}) {
	defer Recovery(ctx, values...)
	f()
}
