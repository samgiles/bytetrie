# Byte Trie

An efficient, iterative byte trie. With retrieval in `O(1)` worst case time
(constant time with respect to the search field, `O(n)` with respect to the
length of the search key).  The tradeoff is space complexity each node
allocates space for a 255 array for fast lookup of child nodes at each level.
In practice this is unlikely to be an issue unless your search fields are very
large and are likely to not contain any common prefixes.

## Usage

Simple usage

```GO
import ( "github.com/samgiles/bytetrie" )
node := new(Node)

node.Insert([]byte("search-key"))
var hasValue = node.HasValue([]byte("search-key"))
```

Using an acceptor:

This could be used to check each byte in a stream against the search field.

```GO
import ( "github.com/samgiles/bytetrie" )
node := new(Node)

var value = []byte("search-key")
node.Insert([]byte("search-key"))

var currentNode = node

for _, b  := range value {
	currentNode, accepted = currentNode.Accepts(b)

	if !accepted {
		fmt.Printf("Your search value is not in the search set")
	}
}
```

## Build and Test

Build:
```SHELL
go build
```

Test:
```SHELL
go test
go test -bench .
```


# License

MIT - Samuel Giles 2015
