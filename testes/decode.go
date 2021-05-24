package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	/* tests cases
	48656c6c6f20476f7068657221
	c8e1331ec7765cfb408dcf9698f24609
	d6f8debdebe090d1d90d72457d36170b77ed7a07c4658dd7a09069aabe9a09ab1501470db0d7ec6123fbdf5952d5e5d4f148bf6cf598fc237b78e12150810bc0
	*/

	const s = "gabriel"
	decoded, err := fromHexToByte(s)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Printf("%s\n", decoded)

}

func fromHexToByte(s string) ([]byte, error) {
	bytes, err := hex.DecodeString(s) // decoding to ASCII hex
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
