# Running go tests from `main()` (cursed)

```sh
$ go run . -test.v
=== RUN   TestHello
hello
=== RUN   TestHello/subtest
subtest
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/subtest (0.00s)
PASS
```
