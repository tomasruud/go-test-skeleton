# Go Test Skeleton

## Install
`go install github.com/tomasruud/go-test-skeleton`

## Run
`go-test-skeleton <file>`

## Example
Given the source file looks like this:

```golang
// file: foo.go
package foo

type bar struct {}

func (_ bar) hi() {}
func (_ bar) Bye() {}

func doFoo() {}
func DoBar() {}
```

The output should look something like this:
```golang
func TestBarHi(t *testing.T) {
  t.Skip("not implemented")
}

func TestBarBye(t *testing.T) {
  t.Skip("not implemented")
}

func TestDoFoo(t *testing.T) {
  t.Skip("not implemented")
}

func TestDoBar(t *testing.T) {
  t.Skip("not implemented")
}
```
