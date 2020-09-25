	package main
  
  func main(){
  
    a := []int{2,4,1}
    
    arrays := make([][]int, 1, len(a))

    for i := range a {

      for _, r := range arrays {
              
        arrays = append(arrays, append([]int{a[i]}, r...))
      }
      
    }
    fmt.Println("SubArrays in given set" , arrays)
  }
