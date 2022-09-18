package crypt

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Gob_encrypt(req interface{}) ([]byte, error) {
	var rplBytes bytes.Buffer
	enc := gob.NewEncoder(&rplBytes)

	err := enc.Encode(req)
	if err != nil {
		return nil, err
	}
	return rplBytes.Bytes(), nil
}
func Gob_decrypt(in []byte, resp interface{}) error {
	var buff = bytes.NewBuffer(in)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(resp)
	if err != nil {
		fmt.Printf("%s\n", "Decoding failed")
		return err
	}
	return nil
}
