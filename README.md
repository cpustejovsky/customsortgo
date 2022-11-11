# Custom Sort Go (CS:GO)
![csgo300](https://user-images.githubusercontent.com/24833119/201412479-de8f5718-0b42-48d3-94d5-6b849597a654.png)

Aside from my punny acronym, this library is to have a collection of custom applications of the [sort package](https://golang.org/pkg/sort)

I'm not a big fan of Leetcode style coding challenges becuase it just doesn't seem to reflect how I provide value as a developer. 

Having pre-written code to allow me to provide value more effectively is, however.

Plan on using this library either by importing the module or copying and pasting for various code challenges to see how they do.

## Use

```go
import (
  "fmt"

  csgo "github.com/cpustejovsky/customsortgo"
)

w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
sorted, err := csgo.Sort(w)

fmt.Println(sorted) // [aardvark albatross bee cat dolphin zebra]
```
