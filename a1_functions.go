package a1
import ("io/ioutil"
        "strings"
        "fmt"
        "errors"
        "reflect"
      )

type Time24 struct {
  hour, minute, second uint8

}
//returns number of primes <=n
func CountPrimes(n int) int{
  //negatives return 0
  if n > 0 {
    //n=1 or n=2 returns 1
    if n > 2 {
      numVar := 1;    //counter

      //only checking odd numbers (no even is prime)
      //starts from 3 as 1 is already counted
      for i := 3; i <= n; i+=2 {
        //check if numbers previous to it will divide it
        //j is less than half of i
        j:= (i+1)/2
        for j>2 && i%j != 0  {    //does not need to check for %2
          j-=1
        }
        if j==2 {     //i was not divisible by any value greater than 2
          numVar+=1
        }    //no numbers divided it
      }
      return numVar
    }
    return 1;
  }
return 0
}

//func countStrings(filename)

func CountStrings(filename string) map[string]int {
  //read file and convert to string
  content, err := ioutil.ReadFile(filename)
  //var check = true //flag for repeats
  if err != nil {
    fmt.Printf("Error with file.")
   }

    m := make(map[string]int)
   str := string(content) // convert content to string
   //create splice with the substrings divided by whitespace
   arr := strings.Fields(str)
   if len(arr)==0 { return m}
   m[arr[0]]=1
   if(len(arr) ==1) {return m}

   //loop through array to search for equal strings
   //each string compares itself to those after it
   for i:=1; i<len(arr); i++ {
       if m[arr[i]]<1 {
            m[arr[i]]=1
       } else {
         m[arr[i]]+=1
       }
     }
   return m
}


func EqualsTime24(a Time24, b Time24) bool {
  if a == b {
    return true
  }
 return false
}

func LessthanTime24(a Time24, b Time24) bool {
  if a.hour < b.hour {return true}

  if a.hour == b.hour {
    if a.minute < b.minute {
      return true
    }
    if a.minute == b.minute {
      if a.second < b.second {
        return true
      }
    }
  }
  return false
}

//converts a time24 to human-readable
func (t Time24)String() string {
  padStr := fmt.Sprintf("%02d:%02d:%02d",t.hour, t.minute,t.second)
return padStr
}

func (t Time24)ValidTime24() bool {

  if 0 <= t.hour && t.hour < 24{
    if 0 <= t.minute && t.minute < 60 {
      if 0 <= t.second && t.second < 60 {
        return true
      }
    }
  }
  return false
}


func MinTime24(times []Time24) (Time24,error) {
  m := Time24{0,0,0}
  e := errors.New("")
  if len(times)==0 {
    e = errors.New("function-minTime24; parameter passed is empty")
    return m,e
  }

  tmp := times[0]
  for i:=0; i < len(times); i++ {
    if times[i].ValidTime24() == false {
      e = errors.New("Slice passed is not valid")
      return m,e
    }
    if LessthanTime24(times[i], tmp) == true {
      tmp = times[i]
    }

    }
   m = tmp
  return m,e
}

//useslinear search to return first index of location x in slice lst
func linearSearch(x interface{}, lst interface{}) (int, error) {
  //returns new value initiazied to the interfac
  valx := reflect.ValueOf(x)
  lstVal := reflect.ValueOf(lst)
  //xType := valx.Type()
//  lstType := lstVal.Type()
  //lstElem := reflect.TypeOf(lst).Elem()

//  if lstType.Kind() != reflect.Slice {
//    panic("list is not of appropriate type: slices")
//  }
  err := errors.New("value is not in list")
  rInt := -1

  //if x cannot be converted to lsttype
  //if xType != lstElem {
  //  panic("x is not of compatible type to list elements")
  //}
  for i:=0; i < lstVal.Len(); i++ {
    if valx == lstVal.Field(i) {
      err = errors.New("")
      return i, err
    }
  }

return rInt,err
}
