package merkledag

import "hash"

func Add(store KVStore, node Node, h hash.Hash) []byte {
	// TODO 将分片写入到KVStore中，并返回Merkle Root
	switch node.Type() {
	case FILE:
		fileNode := node.(File)
		store.Put([]byte("content"), fileNode.Bytes())
		h.Write(fileNode.Bytes())
		return h.Sum(nil)
	case DIR:
		dirNode := node.(Dir)
		iterator := dirNode.It()
		for iterator.Next() {
			childNode := iterator.Node()
			childMerkleRoot := Add(store, childNode, h)
			h.Write(childMerkleRoot)
		}
		return h.Sum(nil)
	default:
	return nil
}
