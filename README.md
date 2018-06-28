# goseq
Sequence id generator in go, can be used to build a stand alone service.

## Usage
`$ go get github.com/FrontMage/goseq`

```go
import "github.com/FrontMage/goseq"

seq := &goseq.MemSequencer{}

for i:=0; i<10; i++{
    fmt.Println(seq.Next)
}
```