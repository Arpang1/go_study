# strings

`strings` 包实现了操作UTF-8编码字符串的简单函数。

#### Compare

    #比较两个字符串的字典顺序，如果 a == b，返回 0；如果 a < b，返回 -1；否则返回 +1。
    strings.Compare(a, b string) int

#### Contains

    #判断字符串 s 是否包含子串 substr。
    strings.Contains(s, substr string) bool

#### ContainsAny

    #判断字符串 s 是否包含 chars 中的任一字符。
    strings.ContainsAny(s, chars string) bool

#### ContainsRune

    #判断字符串 s 是否包含字符 r。
    strings.ContainsRune(s string, r rune) bool

#### Count

    #统计字符串 s 中 substr 出现的次数。
    strings.Count(s, substr string) int

#### EqualFold

    #判断两个字符串在忽略字母大小写的情况下是否相等。
    strings.EqualFold(s, t string) bool

#### Fields

    #将字符串 s 按空白（如空格、制表符、换行符等）分割成多个字符串，并返回这些字符串组成的切片。
    strings.Fields(s string) []string

#### FieldsFunc

    #将字符串 s 按照 f(c) 的返回值分割成多个字符串，并返回这些字符串组成的切片。
    strings.FieldsFunc(s string, f func(rune) bool) []string

#### HasPrefix

    #判断字符串 s 是否以 prefix 开头。
    strings.HasPrefix(s, prefix string) bool

#### HasSuffix

    #判断字符串 s 是否以 suffix 结尾。
    strings.HasSuffix(s, suffix string) bool

#### Index

    #返回字符串 s 中第一个 substr 的索引，如果没有匹配到则返回 -1。
    strings.Index(s, substr string) int

#### IndexAny

    #返回字符串 s 中第一个 chars 中的字符的索引，如果没有匹配到则返回 -1。
    strings.IndexAny(s, chars string) int

#### IndexByte

    #返回字符串 s 中第一个 c 的索引，如果没有匹配到则返回 -1。
    strings.IndexByte(s string, c byte) int

#### IndexFunc

    #返回字符串 s 中第一个满足 f(c) 的字符的索引，如果没有匹配到则返回 -1。
    strings.IndexFunc(s string, f func(rune) bool) int

#### IndexRune

    #返回字符串 s 中第一个 r 的索引，如果没有匹配到则返回 -1。
    strings.IndexRune(s string, r rune) int

#### Join

    #将一系列字符串连接为一个字符串，之间用 sep 分隔。
    strings.Join(a []string, sep string) string

#### LastIndex

    #返回字符串 s 中最后一个 substr 的索引，如果没有匹配到则返回 -1。
    strings.LastIndex(s, substr string) int

#### LastIndexAny

    #返回字符串 s 中最后一个 chars 中的字符的索引，如果没有匹配到则返回 -1。
    strings.LastIndexAny(s, chars string) int

#### LastIndexByte

    #返回字符串 s 中最后一个 c 的索引，如果没有匹配到则返回 -1。
    strings.LastIndexByte(s string, c byte) int

#### Repeat

    #返回将字符串 s 重复 count 次的新字符串。
    strings.Repeat(s string, count int) string

#### Replace
    
    #将字符串 s 中的前 n 个 old 替换为 new，并返回一个新的字符串，如果 n < 0 则替换所有的 old。
    strings.Replace(s, old, new string, n int) string

#### Split

    #将字符串 s 按 sep 分割，并返回一个字符串的切片。
    strings.Split(s, sep string) []string

#### SplitAfter
    
    #将字符串 s 按 sep 分割，并返回一个包含 sep 的字符串的切片。
    strings.SplitAfter(s, sep string) []string

#### ToLower

    #将字符串 s 转换为小写。
    strings.ToLower(s string) string

#### ToLowerSpecial
    
    #将字符串 s 转换为小写，使用特定的规则。
    strings.ToLowerSpecial(c unicode.SpecialCase, s string) string

#### ToUpper

    #将字符串 s 转换为大写。
    strings.ToUpper(s string) string

#### ToUpperSpecial
    
    #将字符串 s 转换为大写，使用特定的规则。
    strings.ToUpperSpecial(c unicode.SpecialCase, s string) string

#### Trim
    
    #将字符串 s 前后端所有包含在 cutset 中的字符都去掉。
    strings.Trim(s string, cutset string) string