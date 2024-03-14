package merkledag
import "hash"

type Link struct {
	Name string
	Hash []byte
	Size int
}
type Object struct {
	Links []Link
	Data  []byte
}
func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {
    var links []Link

    // 如果节点是一个目录，递归地添加其链接的节点
    if dir, ok := node.(Dir); ok {
        it := dir.It()
        for it.Next() {
            linkedNode := it.Node()
            linkedHash, err := Add(store, linkedNode, h)
            if err != nil {
                return nil, err
            }
            // 创建一个Link并添加到links中
            link := Link{
                Name: linkedNode.Name(),
                Hash: linkedHash,
                Size: int(linkedNode.Size()),
            }
            links = append(links, link)
        }
    }

    // 创建一个新的Object并将links添加到Object中
    obj := &Object{
        Links: links,
        Data:  data,
    }

    // 序列化Object的数据
    objData, err := serializeObject(obj)
    if err != nil {
        return nil, err
    }

    // 计算Object的哈希值
    h.Write(objData)
    hashValue := h.Sum(nil)

    // 将Object的数据存储在KVStore中
    if err := store.Put(hashValue, objData); err != nil {
        return nil, err
    }

    return hashValue, nil
}
func serializeNode(node Node) ([]byte, error) {  
	func serializeNode(node Node) ([]byte, error) {
        var data []byte
    
        // 将Size转换为字节切片
        sizeBytes := make([]byte, 8)
        binary.BigEndian.PutUint64(sizeBytes, node.Size())
        data = append(data, sizeBytes...)
    
        // 将Name转换为字节切片
        nameBytes := []byte(node.Name())
        data = append(data, nameBytes...)
    
        // 将Type转换为字节切片
        typeBytes := make([]byte, 4)
        binary.BigEndian.PutUint32(typeBytes, uint32(node.Type()))
        data = append(data, typeBytes...)
    
        return data, nil
    }
}  

