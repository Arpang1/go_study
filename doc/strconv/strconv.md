# strconv

`strconv` 包实现了将字符串转换为基础类型，并且反之亦然。

#### Itoa

    #将 `int` 转换为 `string`。例如，`Itoa(123)` 返回 `"123"`。
    strconv.Itoa(i int) string

#### Atoi

    #将 `string` 转换为 `int`。如果 `s` 不是一个有效的数字字符串，那么将返回一个错误。例如，`Atoi("123")` 返回 `123`。
    strconv.Atoi(s string) (int, error)

#### ParseBool

    #将 `string` 转换为 `bool`。接受 `"1"`, `"t"`, `"T"`, `"true"`, `"TRUE"`, `"True"`, `"0"`, `"f"`, `"F"`, `"false"`, `"FALSE"`, 以及 `"False"`。其它的所有字符串都会返回一个错误。
    strconv.ParseBool(str string) (bool, error)

#### FormatBool

    #将 `bool` 转换为 `string`。如果 `bool` 是 `true`，返回 `"true"`；如果 `bool` 是 `false`，返回 `"false"`。
    strconv.FormatBool(b bool) string

#### ParseFloat

    #将 `string` 转换为 `float64`。`bitSize` 定义了结果类型，可以是 32 或者 64。
    strconv.ParseFloat(s string, bitSize int) (float64, error)

#### FormatFloat

    #将 `float64` 转换为 `string`。
    #`fmt` 定义了格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（大指数为'e'，否则为'f'）、'G'（大指数为'E'，否则为'f'）。
    #`prec` 控制精度（排除指数部分的数字的数量）。对于 'e', 'E', 'f', 's'，它表示小数点后的数字数量。对于 'g' 和 'G'，它控制总的数字数量。如果 `prec` 为 -1，则使用最少数量的、但又必须的数字来表示 f。
    strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string

