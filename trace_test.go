package funchain_test

import "github.com/EnhaoSun/funchain"

func a() {
	defer funchain.Trace()()
	b()
}

func b() {
	defer funchain.Trace()()
	c()
}

func c() {
	defer funchain.Trace()()
	d()
}

func d() {
	defer funchain.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// g[01]:	->github.com/EnhaoSun/funchain_test.a
	// g[01]:		->github.com/EnhaoSun/funchain_test.b
	// g[01]:			->github.com/EnhaoSun/funchain_test.c
	// g[01]:				->github.com/EnhaoSun/funchain_test.d
	// g[01]:				<-github.com/EnhaoSun/funchain_test.d
	// g[01]:			<-github.com/EnhaoSun/funchain_test.c
	// g[01]:		<-github.com/EnhaoSun/funchain_test.b
	// g[01]:	<-github.com/EnhaoSun/funchain_test.a
}
