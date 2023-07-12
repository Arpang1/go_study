# encoding/json

`encoding/json` 包提供了 JSON 数据的编码和解码。

#### Marshal
    
    # Marshal函数将一个数据结构转换为JSON格式的数据
    json.Marshal(v interface{}) ([]byte, error)

#### Unmarshal
    
    # Unmarshal函数将JSON格式的数据转换为一个Go数据结构
    json.Unmarshal(data []byte, v interface{}) error