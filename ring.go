package ring

import (
	"errors"
	"fmt"
	hash1 "github.com/OneOfOne/xxhash"
	"github.com/arriqaaq/rbt"
)

var (
	ERR_EMPTY_RING    = errors.New("empty ring")
	ERR_KEY_NOT_FOUND = errors.New("key not found")
)

type hasher interface {
	hash(string) int64
}

func newXXHash() hasher {
	return xxHash{}
}

// https://cyan4973.github.io/xxHash/
type xxHash struct {
}

func (x xxHash) hash(data string) int64 {
	h := hash1.New32()
	h.Write([]byte(data))
	r := h.Sum32()
	h.Reset()
	return int64(r)
}

type Ring struct {
	store    *rbt.Tree
	nodeMap  map[string]bool
	replicas int
	hashfn   hasher
}

func New() *Ring {
	r := &Ring{
		store:   rbt.NewTree(),
		nodeMap: make(map[string]bool),
		hashfn:  newXXHash(),
	}
	return r
}

func NewRing(nodes []string, replicas int) *Ring {
	r := &Ring{
		store:    rbt.NewTree(),
		nodeMap:  make(map[string]bool),
		replicas: replicas,
		hashfn:   newXXHash(),
	}
	for _, node := range nodes {
		r.nodeMap[node] = true
		hashKey := r.hashfn.hash(node)
		r.store.Insert(hashKey, node)
	}
	return r
}

func (r *Ring) Add(node string) {
	if _, ok := r.nodeMap[node]; ok {
		return
	}
	r.nodeMap[node] = true
	hashKey := r.hashfn.hash(node)
	r.store.Insert(hashKey, node)

	for i := 1; i <= r.replicas; i++ {
		vNodeKey := fmt.Sprintf("%s-%d", node, i)
		hashKey := r.hashfn.hash(vNodeKey)
		r.store.Insert(hashKey, vNodeKey)
	}
}

func (r *Ring) Remove(node string) {
	if _, ok := r.nodeMap[node]; !ok {
		return
	}
	hashKey := r.hashfn.hash(node)
	r.store.Delete(hashKey)
	for i := 1; i <= r.replicas; i++ {
		vNodeKey := fmt.Sprintf("%s-%d", node, i)
		hashKey := r.hashfn.hash(vNodeKey)
		r.store.Delete(hashKey)
	}
	delete(r.nodeMap, node)
}

func (r *Ring) Get(key string) (string, error) {
	if r.store.Size() == 0 {
		return "", errors.New("empty ring")
	}
	hashKey := r.hashfn.hash(key)
	n := r.store.Nearest(hashKey)
	var q *rbt.Node
	if hashKey > n.GetKey() {
		q = rbt.FindSuccessor(n)
	}
	fmt.Println("q: ", hashKey, n)
	if q != nil {
		return q.GetValue(), nil
	}
	return n.GetValue(), nil
}
