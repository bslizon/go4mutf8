# go4mutf8

Codec for Android or JNI MUTF-8(Modified UTF-8) by Go

Reference

- [Oracle Doc](https://docs.oracle.com/javase/1.5.0/docs/guide/jni/spec/types.html#wp16542)
- [/dex/src/main/java/com/android/dex/Mutf8.java](https://www.google.com/search?q=libcore%5Cdex%5Csrc%5Cmain%5Cjava%5Ccom%5Candroid%5Cdex%5CMutf8.java)
- [简书](https://www.jianshu.com/p/f604a4224098)

Example

```go
package main

import (
	"fmt"

	"github.com/bslizon/go4mutf8"
)

func main() {
	m, err := go4mutf8.Encode("Hello World 🐒")
	if err != nil {
		panic(err)
	}

	s, err := go4mutf8.Decode(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}

```