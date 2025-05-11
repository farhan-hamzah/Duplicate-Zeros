package main
import "fmt"

func duplicateZeros(arr []int) {
    n := len(arr)
    
    countZeros := 0
    for i := 0; i < n; i++ {
        if arr[i] == 0 {
            countZeros++
        }
    }
    for i := n - 1; i >= 0; i-- {
        if arr[i] == 0 {
            if i+countZeros < n {
                arr[i+countZeros] = 0
            }
            countZeros--
            if i+countZeros < n {
                arr[i+countZeros] = 0
            }
        } else {
            if i+countZeros < n {
                arr[i+countZeros] = arr[i]
            }
        }
    }
}

func main() {
    arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
    duplicateZeros(arr)
    fmt.Println(arr)
}
