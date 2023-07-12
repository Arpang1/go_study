# github.com/json-iterator/go

`github.com/json-iterator/go` 包提供了一个高性能，兼容标准库的JSON编码和解码的解决方案。

#### 定义一个全局变量

    var json = jsoniter.ConfigCompatibleWithStandardLibrary

#### Marshal
    
    # Marshal函数将一个数据结构转换为JSON格式的数据
    jsonit.Marshal(v interface{}) ([]byte, error)

#### Unmarshal
    
    # Unmarshal函数将JSON格式的数据转换为一个Go数据结构
    json.Unmarshal(data []byte, v interface{}) error
