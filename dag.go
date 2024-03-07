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
func Add(store KVStore, node Node, h hash.Hash) []byte {
    switch node.Type() {
    case FILE:
        return storeFILE(store, node, h)
    case DIR:
        return storeDIR(store, node, h)
    default:
        // 未知的节点类型
        return nil
    }
}

func storeFILE(store KVStore, node Node, h hash.Hash) []byte {
    data := node.(File).Bytes()
    obj := Object{
        Data: data,
    }
    hash := h.Sum(data)
    store.Put(hash, obj)
    return hash
}

func storeDIR(store KVStore, node Node, h hash.Hash) []byte {
    dir := node.(Dir)
    var links []Link
    for it := dir.It(); it.Next(); {
        hash := Add(store, it.Node(), h)
        link := Link{
            Name: it.Node().Name(),
            Hash: hash,
            Size: int(it.Node().Size()),
        }
        links = append(links, link)
    }
    obj := Object{
        Links: links,
    }
    h.Write(obj.Bytes())
    hash := h.Sum(hash)
    store.Put(hash, obj)
    return hash
}
