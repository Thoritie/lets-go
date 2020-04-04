# Golang 101 🐹
##  Let's study GO together 

ref: https://www.dropbox.com/sh/is3hwdqa1dpsb99/AACb9QvxPEUo1Z-tD0Nz2KNZa?dl=0

### History 
1. 2007 Google เจอว่าต้อง Complie Code หลักล้าน บรรทัด กว่าจะ Test ได้รอกันยาวว
2. Robert Griesemer, Rob Pike และ Thompson 🧑🏻‍💻ก่อให้เกิด "GO" ขึ้นมา ซึ่งความดีงามของ GO มีดังนี้
    - Fast compilation
    - Less cumbersome code
    - Unused memory freed automatically (garbage collection)
    - Easy-to-write software software that does several operations simultaneously (concurrency)
    - Good support for prosessors with multiple cores


---- 

### Let's get start
เริ่มกันที่ มาดู Overview file structure เค้่าหน้าตาเป็นไงก่อน

```go
package main
```
📦 ตรงเป็นการบอกว่า code ทุกส่วนที่นั้นเป็นส่วนนึงของ Package ที่ชื่อ "main"
ธรรมชาติของ GO จะต้องประกาศ package ไว้เพื่อบอกให้ compiler รู้ว่าจะ Execute ตรง file นี้


```go
import "fmt"
```
📚เรา Import Library มาใช้เพิ่มเติมได้ ตรงนี้บอกว่าเราจะใช้อะไรบางอย่าง จาก "fmt" package

```go
func main() {
    fmt.Println("Thor say, Greeting GO ")
}
```
💻 ตัวอย่าง function main ตรงนี้เมื่อเราสร้าง function ที่ชื่อ main ไว้ นี่จะเป็นสิ่งแรกที่ถูก Run
fmt.Println = เราดึงเอา function "Println" จาก Package fmt มาใช้
ซึ่ง function นี้ จะแสดงผลว่า "Thor say, Greeting GO"

----

### Export name
เนื่องจาก GO มีความเป็น **Package** ดังนั้น ถ้าอยู่คนละ Package กันการเรียกใช้ function จะต้องเป็นตัวใหญ่
แบบนี้ เราเรียก function `Pi` จาก Package `math`
```go
fmt.Println(math.Pi)
```

--- 

### Function
ทีนี้เราก็มาดูวิธีการ Declare function

```go
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(12, 14))
}
```

ตรง function add เราจะมี Parameter X และ Y ที่มี Type เป็น Int และค่าที่จะ Return ก็มี Type เป็น Int เช่นกัน

ส่วน Main ก็คือส่วน เรียก function add ที่เราเขียนไว้มาใช้งาน


...

MORE TBD..