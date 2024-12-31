package main
import (
    "crypto/sha256"
    "fmt"
    "bytes"
)

// compute SHA-256 hash of given data and return the hash as a byte slice
func hash(data []byte) []byte {
    h := sha256.New()
    h.Write([]byte(data))
    hb := h.Sum(nil)
    
    return hb
}

// gets a slice of strings as arg and returns them as a slice of bytes after hashing each string
func createLeafNodes(datas []string) [][]byte {
    var leafNodes [][]byte
    
    for _, data := range datas {
        leafNodes = append(leafNodes, hash([]byte(data)))
    }
    
    return leafNodes
}

// appends the last element of a byte slice to itself if the length of input byte slice is of odd length
func makeEvenLen(arr [][]byte) [][]byte {
    if len(arr) % 2 == 1 {
        arr = append(arr, arr[len(arr) - 1])
    }
    return arr
}

// creates Merkle tree from a slice of data strings and returns a byte slice representing the Merkle tree with the root being the first element
func createMerkleTree(datas []string) [][]byte {
    leafNodes := createLeafNodes(datas)
    
    var allNodes [][]byte
    allNodes = leafNodes
    allNodes = makeEvenLen(allNodes)
    
    var curLevel [][]byte
    curLevel = allNodes
    
    for true {
        var nextLevel [][]byte
        for i := 0; i < len(curLevel); i += 2 {
            parentHash := hash(append(curLevel[i], curLevel[i+1]...))
            nextLevel = append(nextLevel, parentHash)
        }
        if(len(nextLevel)>1){
            nextLevel = makeEvenLen(nextLevel)
        }
        allNodes = append(nextLevel, allNodes...)
        
        if (len(nextLevel) > 1) {
            curLevel = makeEvenLen(nextLevel)
        } else if (len(nextLevel) == 1) {
            break
        }
    }
    return allNodes
}

// generates Merkle proof of a hash while also returning if the particular hash is found
func makeMerkleProof(hash []byte, merkleTree [][]byte) (bool,[][]byte) {
    var ind = -1
    for i,h := range merkleTree {
        if bytes.Equal(h, hash) {
            ind = i
            break
        }
    }
    if ind == -1 {
        return false, nil
    }
    
    var proof [][]byte
    proof = append(proof, merkleTree[ind])
    var sibInd int
    for ind > 0 {
        if(ind % 2 == 0){
            sibInd = ind - 1
        } else {
            sibInd = ind + 1
        }
        proof = append(proof, merkleTree[sibInd])
        
        ind = (ind - 1) / 2
    }
    return true, proof
}

func main() {
    var datas = []string{"a","b","c","d","e"}
    var mt = createMerkleTree(datas)
    for i,hash := range mt {
        fmt.Printf("%d: %x\n",i,hash)
    }
    
    _, pf := makeMerkleProof(hash([]byte("a")), mt)
    for _,hash := range pf {
        fmt.Printf("%x\n",hash)
    }
}