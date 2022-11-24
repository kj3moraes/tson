# tson
`tson` is JSON viewer and editor written in Go. This repository is forked from the original repository given here - [skanehira/tson](https://github.com/skanehira/tson). This tool displays JSON as a tree and you can search and edit key or values.

![](https://i.imgur.com/9Z6qOY4.gif)

## Support OS
- Mac
- Linux

## Installation
```sh
$ git clone https://github.com/skanehira/tson
$ cd tson && go install
```

## Usage
```sh
# from file
$ tson < test.json

# from pipe
$ curl -X POST http://gorilla/likes/regist | tson

# from url(only can use http get mthod)
$ tson -url http://gorilla/likes/json
```

### Use `tson` as a library in your application
You can use tson in your application as following.

```go
package main

import (
	"fmt"

	tson "github.com/skanehira/tson/lib"
)

func main() {
	j := []byte(`{"name":"gorilla"}`)

	// tson.Edit([]byte) will return []byte, error
	res, err := tson.Edit(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}
```

## Keybinding
### JSON tree

| key    | description                    |
|--------|--------------------------------|
| j      | move down                      |
| k      | move up                        |
| g      | move to the top                |
| G      | move to the bottom             |
| ctrl-f | page up                        |
| ctrl-b | page down                      |
| h      | hide current node              |
| H      | collaspe value nodes           |
| l      | expand current node            |
| L      | expand all nodes               |
| r      | read from file                 |
| s      | save to file                   |
| a      | add new node                   |
| A      | add new value                  |
| d      | clear children nodes           |
| e      | edit json with $EDITOR         |
| q      | quit tson                      |
| Enter  | edit node                      |
| / or f | search nodes                   |
| ?      | show helps                     |
| space  | expand/collaspe children nodes |
| ctrl-j | move to next parent node       |
| ctrk-k | move to next previous node     |
| ctrl-c | quit tson                      |

### help
| key    | description        |
|--------|--------------------|
| j      | move down          |
| k      | move up            |
| g      | move to the top    |
| G      | move to the bottom |
| ctrl-f | page up            |
| ctrl-b | page down          |
| q      | close help         |

## About editing nodes
When editing a node value, the JSON value type is determined based on the value.
For example, after inputed `10.5` and saving the JSON to a file, it will be output as a float type `10.5`.
If the value sorround with `"`, it will be output as string type always.
The following is a list of conversion rules.

| input value        | json type |
|--------------------|-----------|
| `gorilla`          | string    |
| `10.5`             | float     |
| `5`                | int       |
| `true` or `false`  | boolean   |
| `null`             | null      |
| `"10"` or `"true"` | string    |

## About adding new node
You can use `a` to add new node with raw json string.

For expample, you have following tree.

```
{array} <- your cursor in there
├──a
├──b
└──c
```

If you input `{"name":"gorilla"}` and press add button,
then you will get new tree as following.

```
{array} <- your cursor in there
├──a
├──b
├──c
└──{object}
   └──name
      └──gorilla
```

Also, You can use `A` to add new value to current node.

For example, you have following tree.

```
{object} <- your cursor in there
└──name
   └──gorilla
```

If you input `{"age": 26}` and press add button,
then you will get new tree as following.

```
{object} <- your cursor in there
├──name
│  └──gorilla
└──age
   └──26
```

# Author
skanehira
