# randnum

Weight random algorithm with runtime or pre-built data structure\
Time complexity:\
Runtime -> O(n)\
Pre-Built -> O(1)

But pre-built methods needs 1 more (2 finally random numbers) to finish this procedure

Input just data just like:\
input -> weight1,choice1,weight2,choice2,etc...\
rand -> some random int value, link rand.Int

Benchmark result:

```shell
// goos: linux
// goarch: amd64
// pkg: github.com/caskeep/randnum
// cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
// BenchmarkRunTimeWeightedRandomLen3-8            49886893                21.53 ns/op
// BenchmarkStaticTimeWeightedRandomLen3-8         72011685                14.29 ns/op
// BenchmarkRunTimeWeightedRandomLen6-8            47187850                25.12 ns/op
// BenchmarkStaticTimeWeightedRandomLen6-8         71503300                14.40 ns/op
// BenchmarkRunTimeWeightedRandomLen12-8           28984764                41.09 ns/op
// BenchmarkStaticTimeWeightedRandomLen12-8        83873653                14.01 ns/op
// BenchmarkRunTimeWeightedRandomLen24-8           19592414                60.89 ns/op
// BenchmarkStaticTimeWeightedRandomLen24-8        71411912                14.29 ns/op
// PASS
// ok      github.com/caskeep/randnum       12.856s
```

Seen as len(input/data) grow, runtime cpu time grows, but pre-built is always stable complexity.

Further reading:
[Darts, Dice, and Coins: Sampling from a Discrete Distribution](https://www.keithschwarz.com/darts-dice-coins/)
