# Go 1.11 より Go Modules が導入されて、GOPATH 外にプロジェクトを作成可能に。

- Special Thanks【[Create projects independent of $GOPATH using Go Modules](https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)】
- Go@1.11以降を使っていれば、GOPATH を気にせず好きなところにプロジェクト作成 👍

## Go Modules なし(Go@1.11以前)

1. You have to create all your projects inside a single folder where $GOPATH is defined.
2. Versioning Go packages were not supported. It doesn’t allow you to specify a particular version for a Go package like you do in package.json. Also, you couldn’t use two different versions of the same package.

## Go Modules あり(Go@1.11以降)

1. Can create project anywhere you like outside $GOPATH.
2. Packages versioning works by using `go.mod` file.

# プロジェクト作成 - ビルド - 実行まで。

## プロジェクト作成

- `go mod init`コマンドで`go.mod`ファイル自動生成。

```
go mod init github.com/{username}/{project/module name}}
```

## Java は JAR ファイル。Go は実行可能ファイル。

- `go build` -> 実行可能ファイル作成
- `go run` -> 実行可能ファイル作成(`go build`)＋実行 ※**実行ファイルが残らない**
- `go.mod`があるディレクトリで`go build`もしくは`go build .`を実行 -> `go-learning.exe`を生成.
- `go.mod`があるディレクトリで`go run .`を実行 -> main パッケージ.main 関数のあるファイルを起点にビルド＆実行してくれる。
- main パッケージの main 関数から処理が実行される。【エントリポイント】

https://qiita.com/t-yama-3/items/1b6e7e816aa07884378e

# Grammer

## 型 Basic Type

整数

- **整数リテラルで暗黙的に定義したら一律"int"扱いとなる。**

```
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
// uint -> unsigned integer
// - unsigned -> 符号なし
// - int8  -> 8bit(2^8 = 256) -> -128 ~ 127
// - uint8 -> 8bit(2^8 = 256) -> 0 ~ 255
// int と uint = CPUのbit対応数で決まる(32bitか64bit)

byte // alias for uint8
```

浮動小数点数

- **小数リテラルで暗黙的に定義したら一律"float64"扱いとなる。**

```
float32 float64
```

複素数

```
complex64 complex128
```

文字

```
rune(ルーン) // 1つのUnicode文字を格納する

var x rune = 'A' // 65
var y rune = '🙂' // 128072
```

## 型変換 Type Conversion

```go
// int → float64
var x int = 10
var y float64 = float64(x) // 10
// float64 → int (切り捨て/Truncation
var f float64 = 1.6
var i int = int(f) // 1
```

## 宣言 Declaration

- `const` `var` `type` `func`
- A declaration binds a **non-blank identifier** to a `const/var/type/func`.

```go
// まとめて宣言
const (
	size int64 = 1024
	eof        = -1
)
// 型指定なし
const a, b, c = 3, 4, "foo"
// 型が全て同じなら最後に定義するだけでOK
const u, v float32 = 0, 3
```

### Blank Identifier "\_"

-It serves as an anonymous placeholder instead of a regular (non-blank) identifier.

```go
func dummy() (int,int) {
    val1 := 10
    val2 := 12
    return val1,val2
}

func main(){
    rVal, _ := dummy()
    fmt.Println(rVal)
}
```

### iota

- 定数に連番の"定数"を定義するのに便利。
- 値はコンパイル時に決まる。
- 参考：[5 分で完全理解する Go の iota](https://speakerdeck.com/uji/5fen-dewan-quan-li-jie-surugofalseiota)

```go
// Without iota
const (
	Start      Status = 0
	OnGoing    Status = 1
	Finished   Status = 2
)

// With iota ... 変数のIndex が代入される
const (
	Start      Status = iota // 0
	OnGoing    Status = iota // 1
	Finished   Status = iota // 2
)
// 最初のConstSpecのみ定義でもOK
const (
	Start      Status = iota // 0
	OnGoing                  // 1
	Finished                 // 2
)
```

### nil

- In Go, `nil` is the zero value for pointers, interfaces, maps, slices, channels and function types, **representing an uninitialized value**.
- `error` is an interface so that it makes sense to check `err` is nil or not.
- 参考：[What does nil mean in golang?](https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang)

```go
func LoadEnvFile() {
    // Errorがなければ、err変数(ErrorIF)は初期化されておらず
    // nilのままとなるから「err != nil」のチェックを入れる。
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("LOAD .env file")
	}
}
```

## Scope

1. universe block -> [predeclared identifier](https://go.dev/ref/spec#Predeclared_identifiers)
2. package block -> const/vars outside any function
3. function body -> const/vars inside any function

```go
// package level ->
var Db *gorm.DB
var err error

func SetupDb() {
    // function body -> 関数実行時のみ
	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")

```

## Private Public

- Java であれば `public/protected/private`で判別。
- Go では `頭文字が大文字か小文字か` で判別。

```go
// Can access from external module
func SaveTweet (t *Tweet) {
    executeInsertSql(t);
}
// Cannot access from external module
func executeInsertSql(t *Tweet) {
    db.query("INSERT INTO tweet VALUES ...")
}
```

```go
package database

// 小文字 -> databaseパッケージ以外からAccess不可能
var conn *gorm.DB = nil

// 大文字 -> databaseパッケージ以外からAccess可能
func GetConnection() {
    if (conn === nil) {
        conn = connect()
    }
    return conn
}
```

## struct と Pointer 変数

```go
func main() {
    p1 := Person{"daisuke", 26}
    // & で 変数のメモリアドレス を取得
    // "ポインタ変数" の型は「const pointerVar *string」と表現する。
    var p2 *Person = &p1
    // or
    p2 := &p2


    fmt.Println("person info", p1) // person info {daisuke 26}
    fmt.Println("memory address", &p2) // memory address 0xc000006028
    // *ポインタ変数 で中身を参照
    fmt.Println("person info", *p2) // person info {daisuke 26}
}
```

### struct に紐づける形でメソッド定義。

```go
type Person struct {
    Name string
    Age int
}

// 型が ポインタ ではない場合、新しくstructが生成される
func (p Person) Introduce(msg string) {
    // 呼び出し元のstructとは異なるメモリ
    memory := &p
    fmt.Println(&memory)
}

// 型が ポインタ -である場合、呼び出し元のstructのポインタが渡される。
func (p *Person) Introduce(msg string) {
    p.Age = p.Age + "歳"
    fmt.Println("NAME: ", p.Name, "AGE: ", p.Age, msg)
}
```

### フィールド名 = 型名 の場合、省略可能

```go
type Product struct {
    Id string
    Name string
    Company  // Field名 = 型名 なので型定義不要
}

type Company struct {
    Id string
    Name string
}

func main() {
    product1 := Product{"p001", "iPhone", Company{"c001", "Apple"}}
```

### NewXxxx - struct 生成用の関数を作るのが定番らしい

```go
type Person struct {
    Name string
    Age int
}

func NewPerson(id string, age int) *Person {
	return &Person{id, age}
}

func main() {
    // Pointerが返ってきている
    p1 := NewPerson("daisuke", 26)
    // 値を参照
    fmt.Println(*p1) // {daisuke, 26}
```

## Package と Module

Package とは

- 各`----.go` の 1 行目に定義する名前空間。
- 同じ package がついているソースファイル間は、**関数や変数などが共有される = Import 不要で参照可能。**

Module とは

- 複数の Package の集合体。

```
・main.go
/database
　・connection.go
  ・repository.go
```

```go
// connection.go
package database

var Conn *gorm.DB

func Setup() {
    Conn, err := connect()
    if err != nil {
        logger.error("Failed to connect to database.")
    }
}

// repository.go
package database

func FindProductById(id int) {
    // 同じdatabaseパッケージ内のConnをImportせずに参照可能
    return Conn.query("SELECT * FROM product WHERE id = ", id)
}
```

## 関数(function) と メソッド(method)

- A function can be used by `func` keyword.
- A method is just a function **with a receiver argument.**

引数の型が同じ場合、最後に定義するだけで OK。

```go
func add(x, y int) int {
    return x + y
}
```

戻り値は２つ以上返せる。

```go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

戻り値の変数定義 を事前に。

```go
func SplitNames(nameStr string) (res []string) {
    res = strings.Split(nameStr, ",")
    return
}
```

`...int`で可変長引数となる。関数側では**Slice になっている。**<br>
**可変長引数に Slice を渡す場合は`slice...`とする必要がある。**

```go
func hello(nums ...int) {
    for _, val := range nums {
        fmt.Println(val)
    }
}

func main() {
    hello(1,2,3,4,5)

    slice1 := []int{1,2,3,4,5}
    hello(slice1...)
}
```

## 無名関数 をその場で実行する。

```go
func main() {
    func() {
        fmt.Println("無名関数")
    }() //ここの()が必須
}
```

## Defer

- `defer`の対象は**関数の呼び出し、`defer num := 2`とかは違う。**
- `defer`した関数は その関数が return する`前`に呼ばれる。

```go
func main() {
    name := "takakuwa"
    // defer(遅延)させた関数に渡した引数 name は即時評価
    defer fmt.Printf("My name is %s", name)
    name = "makito"

    // "My name is takakuwa"
```

- LIFO なので 最後に defer した関数から実行される。

```go
func main() {
    for i := 1 ; i < 10 ; i++ {
        defer fmt.Println(i)
    }
    // 9,8,7,6,5,4,3,2,1
```

## Defer で Panic を対処/Recover する。

- panic が発生すると後続の処理を行わず panic を return する...を繰り返す。
- **return する前に実行される defer 関数**を利用して panic を制御する。
- 参考：[Go 言語の defer を正しく理解する | How defer in Golang works
  Go
  初心者](https://qiita.com/Ishidall/items/8dd663de5755a15e84f2#panic-defer-recover)

```go
// hello関数よりpanicがreturnされ、main関数が異常終了する。
func hello() {
    fmt.Println("★1")
    panic("some error")
    fmt.Println("★2")
}

func main() {
    hello()
}
```

```go
func hello() {
    // panic が returnされる前に実行される
    defer func() {
        // recover() -> panicの引数
        if r := recover(); r != nil {
            fmt.Printf("Error message: %s", r) // PANIC ERROR MESSAGE
        }
    }()

    fmt.Println("★1")
    panic("PANIC ERROR MESSAGE")
    fmt.Println("★2")
}

func main() {
    hello()
}
```

## Defer でファイルクローズ

```go
func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    file.Write([]byte("HELLO"))
}
```

## Closure

```go
func storeName() func(name string) []string {
    var nameStore []string
    return func(name string) []string {
        if (name != "") {
            nameStore = append(nameStore, name)
        }
        return nameStore
    }
}

func main() {
    xstoreName := storeName()
    xstoreName("takakuwa")
    xstoreName("makito")
    xstoreName("fumi")

    for _, val := range xstoreName("") {
        fmt.Println(val) // takakuwa, makito, fumi
    }
```

## interface と interface{}

### interface

- struct/構造体 が interface で定義している関数を実装していたら、**暗黙的に implements 扱いとなる。**
- 以下, error インタフェースに沿ってカスタムエラー struct を定義した例

```go
// errorインタフェース
// type error interface {
// 	Error() string
// }

type InternalServerError struct {
    message string
}

func (err *InternalServerError) Error() string {
    return fmt.Sprintf("500 Error: %s", err.message)
}

type BadRequestError struct {
    message string
}

func (err *BadRequestError) Error() string {
    return fmt.Sprintf("400 Error: %s", err.message)
}

func validateRequest() (err error) {
    // Validationエラーがあったとして
    return &BadRequestError{"ItemA is invalid value."}
}

func main() {
    if err := validateRequest(); err != nil {
        fmt.Println(err.Error()) // 400 Error: ItemA is invalid value.
    }
```

### interface{}

- 型の１つ。`int string bool`の仲間。
- {}の部分まで含めて型名。
- **どんな型でも格納 OK。演算はできない。**
- 利用例: JSON 読取(文字/interface{}) -変換-> 特定の struct へ変換

### 変数の型確認方法(①reflect.TypeOf ②fmt.Printf("%T", var))
```go
// reflect.TypeOf
func main() {
	sec := 1 * time.Second
	fmt.Println(reflect.TypeOf(sec)) // time.Duration
	fmt.Printf("%T", sec)            // time.Duration
```

### Type Assertion / 型参照

- 第 1 引数:値　第 2 引数:参照結果(true/false)

```go
// 型参照結果(第２引数)を受け取らない
var num interface{} = 1
// 第2引数を受け取ってないためpanicが起きる。
str := num.(string) // panic: interface conversion: interface {} is int, not string

// 型参照結果(第２引数)を受け取る
var num interface{} = 1
str, isString := num.(string)
fmt.Println(str, isString)     // 空文字 false
```

### Type Switch / 型スイッチ

```go
func main() {
    var x interface{} = 1
    switch x.(type) {
        case int :
            fmt.Println("type is int")
        case string :
            fmt.Println("type is string")
        case bool :
            fmt.Println("type is boolean")
        default :
            fmt.Println("UNEXPECTED TYPE")
    }
}
```

## ブロック

### if

基本形

```go
func main() {
    num := 500
    // if cond
    if num < 100 {
        fmt.Println("num less than 100")
    } else if num >= 100 & num < 500 {
        fmt.Println("num is between 100 to 499")
    } else {
        fmt.Println("num is over thatn 500")
    }
```

if 文 with initialization

```go
func calcBMI(height float32, weight float32) (bmi float32) {
    bmi =  weight / (height * height)
    return
}

func main() {
    var height float32 = 1.8
    var weight float32 = 74.5
    // if initialization ; cond
    if bmi := calcBMI(height,weight); bmi > 25 {
        fmt.Println("肥満")
    } else if bmi < 25 {
        fmt.Println("NOT肥満")
    }
```

### for

基本形

```go
for i := 0; i < 10; i++ {
    fmt.Println("HELLO WORLD", i)
}
```

The init(1) and post(3) statements are optional.
for 1 ; 2 ; 3 { }

```go
func main() {
    sum := 1
    for ; sum < 1000 ; {
        sum += sum
    }
}
```

### 拡張 for 文（range）

```go
func main() {
    nums := []string{"A","B","C","D","E"}
    for idx, val := range nums {
        fmt.Printf("index:%d, value:%s \n", idx, val)
    }
```

### while 文(存在しないので for 文で実現)

```go
func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
}
```

### switch

基本形

```go
switch n {
    case 3:
        fmt.Println("n is 3")
    case 2:
        fmt.Println("n is 2")
    case 1:
        fmt.Println("n is 1")
    default:
        fmt.Println("n is other than 1,2,3")
```

複数の値の定義も可能

```go
n := 5

switch n {
    case 4, 5, 6:
        fmt.Println("n:", n)
    case 1, 2, 3:
        fmt.Println("n:", n)
}
```

型スイッチ

```go
var x interface{} = 1
switch x.(type) {
    case int :
        fmt.Println("type is int")
    case string :
        fmt.Println("type is string")
    case bool :
        fmt.Println("type is boolean")
    default :
        fmt.Println("UNEXPECTED TYPE")
}
```

## Array Slice

- Array: サイズ固定, **NOT ポインタ**
- Slice: サイズ動的, **配列を指定するポインタ**
- 参考：[Golang で Slice(スライス)について説明して、使う方法に関しても説明します](https://dev-yakuza.posstree.com/golang/slice/)

```go
// Array
var nums1 [4]int  //要素数を指定したら配列 [0 0 0 0]
// Slice
var nums2 []int   //要素数が未指定 -> Slice ★nil★

// 初期値ありのslice
nums := []string{"A","B","C","D","E"}
```

```go
// sliceはポインタなので 関数を渡っても同じインスタンスが使われる
func hello(slice []int) {
    slice[0] = 999
}

func main() {
    numbers := []int{1,2,3,4,5}
    hello(numbers)
    for _, val := range numbers {
    fmt.Println(val) // [999,2,3,4,5]
    }
}
```

make 関数　-> 初期値 0 無しで 要素数指定して slice 作成可能<br>
len 関数　　-> slice が持つ要素数<br>
cap 関数　　-> slice の総サイズ。要素数が Over しそうになれば自動で倍増する。<br>
append 関数 -> 要素が追加された配列を返す。第 2 引数は可変長引数。

```go
func main() {
    // 第1引数-sliceの型  第2引数-slice要素数(拡張可能)
    names := make([]int, 5) // [0,0,0,0,0]
    fmt.Println(len(names), cap(names)) // 5, 5
    names = append(names, 1) // [0,0,0,0,0,1]
    fmt.Println(len(names), cap(names)) // 6, 10
    names = append(names, 2,3,4) // [0,0,0,0,0,1,2,3,4]
    fmt.Println(len(names), cap(names)) // 9, 10
    for _, name := range names {
        fmt.Println(name)
    }
```

### Slicing ... 配列より Slice 作成

```go
func main() {
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:2]
```

### Slice1 をコピーして Slice2 を作成する（異なるメモリアドレスで）

- Slicing(`slice[:]`)だと同じメモリアドレスを共有してしまう。

```go
// 同じメモリアドレスを共有してる
func main() {
    slice1 := []int{1,2,3,4,5}
    slice2 := slice1[:]

    slice2[1] = 100

    fmt.Println(slice1) // [100,2,3,4,5]
    fmt.Println(slice2) // [100,2,3,4,5]
```

- 新しいインスタンス(異なるメモリアドレス)を作成する場合は、<br>①append 関数で slice1 をもとに slice2 を作成する。<br>②make 関数で作った slice2 へ copy 関数で slice1 要素を追加する。<br>③ make 関数で slice2 を作成して そこにループで slice1 要素を追加する。

①append 関数で slice1 をもとに slice2 を作成する。

```go
// appendの第2引数は可変長引数なので sliceのまま渡せないから ...
func main() {
    slice1 := []int{1,2,3,4,5}
    slice2 := append([]int{}, slice1...)

    slice1[0] = 999

    fmt.Println(slice1) // [999,2,3,4,5]
    fmt.Println(slice2) // [1,2,3,4,5]
```

②make 関数で作った slice2 へ copy 関数で slice1 要素を追加する。

```go
func main() {
    slice1 := []int{1,2,3,4,5}
    slice2 := make([]int, len(slice1))
    copy(slice2, slice1) // copy(dist, src)

    slice1[0] = 999

    fmt.Println(slice1)
    fmt.Println(slice2)
}
```

③ make 関数で slice2 を作成して そこにループで slice1 要素を追加する。

```go
// 異なるメモリアドレス
func main() {
    slice1 := []int{1,2,3,4,5}
    slice2 := make([]int, len(slice1))

    for idx, val := range slice1 {
        slice2[idx] = val
    }

    slice2[1] = 100

    fmt.Println(slice1) // [1,2,3,4,5]
    fmt.Println(slice2) // [100,2,3,4,5]
```

### Slice を引数に渡す

## Map

- `make(map[type]type)` で空の Map を作成する。
- 初期値ありの場合は `map[type]type{}`

```go
func main() {
    m = make(map[string]string)
    m["KEY1"] = "VALUE1"
    m["KEY2"] = "VALUE2"
    m["KEY3"] = "VALUE3"
    fmt.Println(m["KEY1"])
    fmt.Println(m["KEY2"])
    fmt.Println(m["KEY3"])
```

初期値あり

```go
func main() {
    var m = map[string]string{
        "KEY1":"VALUE1",
        "KEY2":"VALUE2",
        "KEY3":"VALUE3",
    }
    m["KEY4"] = "VALUE4"
    m["KEY5"] = "VALUE5"
    m["KEY6"] = "VALUE6"

    // 順序はランダム
    for key, value := range m {
    fmt.Printf("key: %s, value: %s \n", key, value)
    }
```

順序をソートしたい -> Key 用の Slice 作成 -> これをソートして Loop する。

```go
func main() {
    var m = map[string]string{
        "KEY1":"VALUE1",
        "KEY2":"VALUE2",
    "KEY3":"VALUE3",
    }
    m["KEY4"] = "VALUE4"
    m["KEY5"] = "VALUE5"
    m["KEY6"] = "VALUE6"

    // Keyだけ取り出す
    var sortedKeys []string
    for key, _ := range m {
        sortedKeys = append(sortedKeys, key)
    }
    sort.Strings(sortedKeys)

    // SortされたKeyでLoopしてMapの中身を取り出す
    for _, key := range sortedKeys {
        fmt.Println(m[key])
    }
```

要素削除

```go
func main() {
    nameMap := map[int]string{1:"takakuwa", 2:"makito", 3:"fumi"}
    fmt.Println(len(nameMap)) // 3
    delete(nameMap, 1) // 副作用
    fmt.Println(len(nameMap)) // 2
```

要素有無チェック

```go
func main() {
    nameMap := make(map[int]string)
    nameMap[1] = "takakuwa"
    nameMap[2] = "makito"

    if val, ok := nameMap[3]; ok {
        fmt.Println("Here is fumi.", val)
    } else {
        fmt.Println("NO fumi here.")
    }
```

## エラーハンドリング

- 参考：[Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/intro)

# 並列処理 と 並行処理

- 参考：[並行処理と並列処理](https://zenn.dev/hsaki/books/golang-concurrency/viewer/term)
- 並行/Concurrency：複数の処理を**独立に**実行できる構成のこと
- 並列/Parallelism：複数の処理を**同時に**実行すること
- **並列は並行を包含している**ので、「**並行処理は、並列に実行されるかもしれない。**」とも言える。

## Go で 複数の処理を同時に実行できるようにする。

- goroutine を使うことで、main スレッド(routine)とは別スレッド/routine を`go`文を使うことで作成できる。
- 各 routine は、チャネルを介すことで値の受渡しをすることができる。
- main ルーチンが終わるとその他のルーチンを強制終了されるため、`sync.WaitGroup`を使う等して**子ルーチンが終了するまで main ルーチンをブロックする**という制御を入れてあげる。
- ✔Go で並行処理を実現するには、main ルーチンより複数のルーチンを意図的に作成してそこで処理を並行に実行させることになる。
- API の場合「**リクエストが来るごとに、新しいゴールーチンを go 文で立てる**」といった動きになる。

# 標準ライブラリ standard library

## [math/rand](https://pkg.go.dev/math/rand)
- 乱数を生成するためのライブラリ。
- 生成される乱数が一定 OR 毎回異なる値となるように**Seed値**を設定する。

### 使い方（DefaultRandインスタンスを使う場合）
1. `rand.Seed`関数 -> Seed値設定
2. 好きな乱数生成関数を使う。(`rand.Int()`関数, `rand.Intn(n)`関数, `rand.Float64()`関数)
```go
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int())      // 7726282257792290944
	fmt.Println(rand.Intn(1000)) // 750
```

### 使い方（自らRandインスタンスを生成する場合）
1. [rand.NewSource(seed int)](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L51)関数でSourceインスタンス生成。
2. [rand.New(src Source)](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L78)関数でRandインスタンス生成。
3. 2で生成したRandインスタンスで各乱数生成関数を使う。
```go
func NewRandSource() *rand.Rand {
    src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)
    return rnd
}

func main() {
	rnd := NewRandSource()
	fmt.Println(rnd.Int())      // 3247200792509923393
	fmt.Println(rnd.Intn(1000)) // 283
```

### 早見表
|種類|関数|
|----|----|
|整数(64bit)|`rand.Int()`|
|整数(32bit)|`rand.Int31()`|
|整数(0~Nの間)|`rand.Intn()`|
|小数(64bit)|`rand.Float64()`|
|小数(32bit)|`rand.Float32()`|

### rand.Xxxxで使える`関数` と struct経由で使える`メソッド`
- ライブラリ内で、Defaultの乱数generator「[globalRandGenerator](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L312)」変数が定義されており、**`rand.Xxxx`関数ではこのDefaultの乱数generatorを利用している。**
```go
// globalRand() で globalRandGenerator を取得
func Int() int { return globalRand().Int() }
func Int31() int32 { return globalRand().Int31() }
```
- 対して、自らRandインスタンスを生成した場合は そのインスタンスに対して以下の`(r *Rand)付のメソッド`を呼び出すことになる。
```go
func (r *Rand) Int() int {
	u := uint(r.Int63())
	return int(u << 1 >> 1)
}
func (r *Rand) Int31() int32 { return int32(r.Int63() >> 32) }
```

### SourceはIFで、RandはStruct。
- [rand.NewSource(seed int)関数](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L51)では、SourceIFにMatchするStructインスタンスを生成。
- [rand.New(src Source)関数](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L78)では、RandのStructインスタンスを生成。
```go
type Source interface {
	Int63() int64
	Seed(seed int64)
}
// SourceIFの実装であるRand構造体を使っていく。
type Rand struct {
    //....
}
func (r *Rand) Int63() int64 {}
func (r *Rand) Seed(seed int64)
```

## [time](https://pkg.go.dev/time)
- Go1.9以前：`time.Now()` -> "wall clock"のみ取得
- Go1.9以降：`time.Now()` -> "wall clock"＋"monotonic clock"を取得。
- wall clock：時刻を示す。YYYY/MM/DD HH時MM分SS秒
- monotonic clock：時刻を計算する。
### 使い方
- `time.Now()`関数 -> Local Timezone付きで現在時刻取得。
- `LoadLocation(定義済TZ名 string)`関数 -> Location(TZ)取得
- **LoadLocation関数は外部依存があるため、ソースがランタイムに依存する** -> **決め打ちで`FixedZone(TZ名, offset)`関数を使う方が確実。**
- `time.Date(year,month,day,hour,min,sec,nsec,loc)`関数 -> Time生成
- `(t *time).Format(format string)`関数：【Time→Format文字列】
- **引数layoutに渡す時刻は「2006年1月2日15時4分5秒 アメリカ山地標準時MST(GMT-0700)」のものを使うことになっている。。**
- `time.Parse(format, targetStr string)`関数：【文字列→Time生成】
- `time.ParseInLocation(format, targetStr, loc)`関数：【TZ込文字列→
Time生成】
- 参考：[Goで時刻を扱うチートシート](https://zenn.dev/hsaki/articles/go-time-cheatsheet)
```go
func main() {
	// time.Now() -> Local Timezoneを設定してくれる
	japanTime := time.Now()
	fmt.Println(japanTime)          // 2023-09-23 12:17:12.6611893 +0900 JST m=+0.002200101
	fmt.Println(japanTime.String()) // 2023-09-23 12:17:12.6611893 +0900 JST m=+0.002200101
	// UTC時刻取得
	utcTime := japanTime.UTC()
	fmt.Println(utcTime)            // 2023-09-23 03:18:57.0224089 +0000 UTC
    // Unix時間取得
	fmt.Println(utcTime.Unix())     // 1695441393

	// Timezone指定
	// 1. 既存Timezone取得【LoadLocation関数】
	jst, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }
	// 2. 自分でTimezone作成【FixedZone関数】
	customJst := time.FixedZone("Custom Asia/Tokyo", 9*60*60)
	t1 := time.Date(2023, time.September, 10, 23, 0, 0, 0, jst)
	t2 := time.Date(2023, time.September, 10, 23, 0, 0, 0, customJst)
	// UTC時刻に時差が考慮された時刻が表示される
	fmt.Println(t1) // 2009-11-10 23:00:00 +0900 JST
	fmt.Println(t2) // 2009-11-10 23:00:00 +0900 Custom Asia/Tokyo

	// 3. Time -> Format　して文字列取得【Format関数】
	// 既存Format: https://pkg.go.dev/time#pkg-constants
	fmt.Println(t1.Format(time.RFC3339))          // 2023-09-10T23:00:00+09:00
	fmt.Println(t1.Format("2006-01-02"))          // 2023-09-10
	fmt.Println(t1.Format("2006-01-02 15:04:05")) // 2023-09-10 23:00:00

	// 4. 文字列 -> Time【Parse(format string, targetString: string)関数】
	// 4-1. Parse without timezone
	targetStr1 := "2022-03-23"
	format1 := "2006-01-02"
	t3, err := time.Parse(format1, targetStr1)
	if err != nil {
		panic(err)
	}
	fmt.Println(t3) // 2022-03-23 00:00:00 +0000 UTC
	// 4-1. Parse with timezone
	targetStr2 := "2022-03-23T07:00:00+09:00"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t4, err := time.ParseInLocation(time.RFC3339, targetStr2, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t4) // 2022-03-23 07:00:00 +0900 JST
```


# 開発環境

- Linter（goLint）
- 自動フォーマット
- 自動 IMPORT
