package libs

import (
    "fmt"
)

func HelloForEachNums(nums int) {
    for i :=0; i < nums; i++ {
        fmt.Println("Hello! - ", i)
    }
}
