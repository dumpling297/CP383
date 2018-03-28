package main
import (
        "encoding/json"
        //"io/ioutil"
        "strings"
        //"reflect"
        "os"
        "fmt"
)

//global array of tokens
var arrWrite []interface{}
var count int

func main() {

  arrWrite = append(arrWrite, "{")
  JsonToken()
  HTMLFormat()

}

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func JsonToken() {

  //get file and read it; check the size
  jFile := os.Args[1]
  content, err1 := os.Open(jFile)
  CheckError(err1)
  s, err2 := content.Stat()
  CheckError(err2)
  size := s.Size()
  //array to hold tokens

  //check for empty file; assign "error"
  if(size == 2){
    fmt.Printf("empty file")
  }

  //output the file
  back := make([]byte, size)
  _, err3 := content.Read(back)
  CheckError(err3)
  //fmt.Printf("Display JSON file: %d bytes: %s\n", n1, string(back))

   //extract tokens with json interface
   var f interface{}
   err4 := json.Unmarshal(back, &f)
   CheckError(err4)
    mJson := f.(map[string]interface{})
    //i := len(mJson)
    //j :=1
    SearchArray(mJson)
    //for loop for cases
    /*for key, values := range mJson {
      //look through values of keys -- case types

      arrWrite = append(arrWrite,key)
      arrWrite = append(arrWrite, ":")

      switch first := values.(type) {
        //array cases
        case []interface{}:
          //check what type of value the array has
            check := values.([]interface{})[0]
            switch check.(type){
            case float64:
              arrWrite = append(arrWrite,"[")
                for _, ven := range values.([]interface{}) {
                  arrWrite = append(arrWrite,ven)
                }
                arrWrite = append(arrWrite,"]")
            case string:
              arrWrite = append(arrWrite,"[")
              for _, ven := range values.([]interface{}) {
                arrWrite = append(arrWrite,ven)
              }
              arrWrite = append(arrWrite,"]")
            case bool:
              arrWrite = append(arrWrite,"[")
              for _, ven := range values.([]interface{}) {
                arrWrite = append(arrWrite,ven)
              }
              arrWrite = append(arrWrite,"]")
            case nil:
              arrWrite = append(arrWrite,"[")
                arrWrite = append(arrWrite,"")
              arrWrite = append(arrWrite,"]")
            default:
              arrK := values.([]interface{})[0].(map[string]interface{})
              arrWrite = append(arrWrite,"[")
              arrWrite = append(arrWrite,"{")
              SearchArray(arrK)
              arrWrite = append(arrWrite,"}")
              arrWrite = append(arrWrite,"]")
          }
        //regular cases
        case string:
          arrWrite = append(arrWrite,first)
          //fmt.Println(key, "is string", first)
          //check for escape characters
        case nil:
          arrWrite = append(arrWrite,first)
         //fmt.Println(key, "is null")
         //just return empty string
        case bool:
          arrWrite = append(arrWrite,first)
        //fmt.Println(key, "is bool", first)
        //return bool val
        case float64:
          arrWrite = append(arrWrite,first)
        //fmt.Println(key, "is int", first)
        //return the val
        default:
          arrWrite = append(arrWrite,first)
        //fmt.Println(key, "cannot handle this type")
        }
      if j<i {
        arrWrite = append(arrWrite,",")
        j+=1
      }

  }
  arrWrite = append(arrWrite,"}")
  //fmt.Println("Displaying Array of tokens:\n", arrWrite)*/
    arrWrite = append(arrWrite, "}")
  content.Close()
}

//recursive function, searches through the array
//makes sure to search any arrays within the array too
func SearchArray(m map[string]interface{}) {
  //while loops for arrays within
  arr := make(map[string]interface{})
  i := len(m)
  j :=1
  for key2, val2 := range m {
    arrWrite = append(arrWrite,key2)
    arrWrite = append(arrWrite,":")
      switch again := val2.(type) {
        case []interface{}:
          a := val2.([]interface{})
          l := len(a)-1
          check := val2.([]interface{})[0]
          switch check.(type){
          case float64:
              arrWrite = append(arrWrite,"[")
              for r, ven := range val2.([]interface{}) {
                arrWrite = append(arrWrite,ven)
                InsertComma(r,l)
              }
              arrWrite = append(arrWrite,"]")
              InsertComma(j,i)
              j+=1
          case string:
              arrWrite = append(arrWrite,"[")
              for r, ven := range val2.([]interface{}) {
                arrWrite = append(arrWrite,ven)
                InsertComma(r,l)
              }
              arrWrite = append(arrWrite,"]")
              InsertComma(j,i)
              j+=1
          case bool:
              arrWrite = append(arrWrite,"[")
              for r, ven := range val2.([]interface{}) {
                arrWrite = append(arrWrite,ven)
                InsertComma(r,l)
              }
              arrWrite = append(arrWrite,"]")
              InsertComma(j,i)
              j+=1
          case nil:
              arrWrite = append(arrWrite,"[")
              arrWrite = append(arrWrite,"")
              arrWrite = append(arrWrite,"]")
              InsertComma(j,i)
              j+=1
          default:
            arr = val2.([]interface{})[0].(map[string]interface{})
            arrWrite = append(arrWrite,"[")
            arrWrite = append(arrWrite,"{")
            SearchArray(arr)
            arrWrite = append(arrWrite,"}")
            arrWrite = append(arrWrite,"]")
            InsertComma(j,i)
            j+=1
          }
        //regular cases
        case string:
            //fmt.Println(key2, "is string", again)
            arrWrite = append(arrWrite,again)
            InsertComma(j,i)
            j+=1
          //check for escape characters
        case nil:
          arrWrite = append(arrWrite,again)
          InsertComma(j,i)
          j+=1
          //fmt.Println(key2, "is null")
         //just return empty string
        case bool:
          arrWrite = append(arrWrite,again)
          InsertComma(j,i)
          j+=1
          //fmt.Println(key2, "is bool", again)
        //return bool val
        case float64:
          arrWrite = append(arrWrite,again)
          InsertComma(j,i)
          j+=1
          //fmt.Println(key2, "is float64", again)
        //return the val
        default:
          //fmt.Println(key2, "is of a type I don't know how to handle")
        }
    }
}

func InsertComma(x int, y int) {
  if x < y {
    arrWrite = append(arrWrite,",")
  }
}

func HTMLFormat() {
  //beginning of html docs
  fmt.Println("<!Doctype html>\n")
  fmt.Println("<head>\n","JSON<br>","</head>")
  fmt.Println("<body>")

  //check through token array & colour/format
  lenArr := len(arrWrite)
  i:=0
  count = 1
  var Index interface{}
  for i<lenArr {
    Index = arrWrite[i];
    switch w := Index.(type) {
    case string:
      switch {
      case Index == "{":
        fmt.Println("<span style='color:LightSalmon'>{</span><br>")
        AddSpace(count);
        count++
      case Index == "[":
        fmt.Println("<span style='color:DarkTurquoise'>[</span>")
      case Index == "}":
        //AddSpace(count)
        fmt.Println("<br><span style='color:LightSalmon'>}</span>")
      case Index == "]":
        fmt.Println("<span style='color:DarkTurquoise'>]</span>")
      case Index == ":":
        fmt.Println("<span style='color:Teal'> : </span>")
      case Index == ",":
        fmt.Println("<span style='color:Red'>,</span><br>&emsp;")
      default:
        //all other strings
        //needs a StringChecker() for html proper output
        CheckString(Index)
      }
    case float64:
      fmt.Println("<span style='color:OrangeRed'>",w,"</span><br>")
    default:
      fmt.Println("<span style='color:LawnGreen'>",Index,"</span>;")
  }
  i+=1
 }

fmt.Println("</body>\n","</html>\n")
}


func CheckString(v interface{}) {
 v1 := v.(string)
 rep := strings.Replace(v1, "<", "&lt;", -1)
 rep = strings.Replace(v1, ">", "&gt;", -1)
 rep = strings.Replace(v1, "&", "&amp;", -1)
 rep = strings.Replace(v1, "'", "&apos;", -1)
 rep = strings.Replace(v1,`\n`, "<span style='color:Orange'></span>",-1)


 //AddSpace(count)
 fmt.Printf("<span style='color:PaleVioletRed'>&quot;%s&quot;</span> ", rep)
}

func AddSpace(i int) {
  j :=0;
  for j<i {
    fmt.Printf("&emsp;")
    j+=1
  }
}
