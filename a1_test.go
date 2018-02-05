package a1
import ("testing"
        "log"
        "strings"
        "io/ioutil"
       "fmt"
       "strconv"
  //     "errors"
      )

func TestCountPrimes(t *testing.T) {
  cases := []struct {
    in, want int
  }{
    {-6, 0},
    {5, 3},
    {10000, 1229},
  }
  for _, c := range cases {
    got := CountPrimes(c.in)
    if got != c.want {
      t.Errorf("CountPrimes(%d) == %d, want %d", c.in, got, c.want)
    }
  }
}

func TestCountStrings(t *testing.T) {
  //checking input file
  c,err := ioutil.ReadFile("filename.txt")
  if err != nil {
       log.Fatal(err)
   }

  got := CountStrings("filename.txt")
  //if c holds 2 or less bytes, it is empty
  if len(got)==0 && len(c)>2 {
    t.Errorf("Function returns empty map for a non-empty file")
  }
  if len(got)!=0 && len(c)<=2{
     t.Errorf("Function returns non-empty map for empty file")
  }

//test for null values within got
 testerc := string(c) // convert content to string
 //create splice with the substrings divided by whitespace
 arr := strings.Fields(testerc)
 for i:=0; i<len(arr); i++ {
   if got[arr[i]]<1 {
     t.Errorf("All key strings have not been mapped to a value")
   }
 }
}

func TestEqualsTime24(t *testing.T) {

  cases := []struct {
    in1, in2 Time24
    want bool
  }{
    {Time24{0,8,5}, Time24{0,8,6}, false},
    {Time24{10,30,5}, Time24{10,30,5}, true},
    {Time24{4.0,8,5}, Time24{4,8,5}, true},
    {Time24{0,5,5}, Time24{0,5,4}, false},
  }

  for _, c := range cases {
    got := EqualsTime24(c.in1,c.in2)
    if got != c.want {
      t.Errorf("compared b/w (%d) and (%d) == %t, want %t", c.in1, c.in2, got, c.want)
    }
  }
//  fmt.Print(EqualsTime24(one,two))
}

func TestLessThanTime24(t *testing.T) {
  cases := []struct {
    in1, in2 Time24
    want bool
  }{
    {Time24{2,8,5}, Time24{1,8,5}, false},  //a.hour>b.hour
    {Time24{0,8,5}, Time24{1,8,5}, true},   //a.hour<b.hour
    {Time24{0,8,5}, Time24{0,8,5}, false},  //a=b
    {Time24{0,5,5}, Time24{0,5,4}, false},  //a.sec>b.sec
    {Time24{5,7,10}, Time24{5,8,5}, true},  //a.min<b.min
  }


  for _, c := range cases {
    got := LessthanTime24(c.in1,c.in2)
    if got != c.want {
      fmt.Println(c.in1.minute==c.in2.minute)
      t.Errorf("compared b/w (%d) and (%d) == %t, want %t", c.in1, c.in2, got, c.want)
    }
  }
}

func TestString(t *testing.T) {
  in := Time24{2,8,5}
  pad := in.String()

  s:= strings.Split(pad, ":")
  if len(s)> 3 {
    t.Errorf("method string output is Not properly formatted")
  }

for i:=0; i<3;i++ {
    _, err:= strconv.Atoi(s[i])
    if err!= nil {
      t.Errorf("method string does not contain numeric values")
    }
  }
}


func TestValidTime24(t *testing.T) {
  cases := []struct {
    in Time24
    want bool
  }{
    {Time24{100,8,5}, false},
    {Time24{0,8,5}, true},
    {Time24{0,120,0},false},
    {Time24{13,47,60}, false},
  }
  for _, c := range cases {
    got := c.in.ValidTime24()
    if got != c.want {
      fmt.Println(c.in)
      t.Errorf("Time validity is == %t, want %t", got, c.want)
    }
  }
}

func testminTime24(t *testing.T) {
/*  ex := []Time24{
    {1,1,1},
    {2,3,4},
    {8,7,9},
    {0,0,1},

  }
  f, err := MinTime24(ex)
  if err!=nil {
    fmt.Println(err)
  }
  if f.ValidTime24() == false {
    t.Errorf("Returned slice is not a valid time")
  }

  exEmpty := []Time24{}
  _, err1 := MinTime24(exEmpty)
  if err!=nil {
    fmt.Println(err1)
  }

}

func TestLinearSearch(t *testing.T) {
  casesInt := []interface {
    in int
    lis []int
    want int
  }{
    {5, {4,2,5}, 3},

  }
  for _, c := range casesInt {
    got,err := LinearSearch(c.in, c.lis)
    if got != c.want {
      fmt.Println(err)
    }
  }*/
}
