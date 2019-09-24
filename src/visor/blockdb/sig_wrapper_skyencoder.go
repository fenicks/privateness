// Code generated by github.com/SkycoinProject/skyencoder. DO NOT EDIT.

package blockdb

import "github.com/SkycoinProject/skycoin/src/cipher/encoder"

// encodeSizeSigWrapper computes the size of an encoded object of type sigWrapper
func encodeSizeSigWrapper(obj *sigWrapper) uint64 {
	i0 := uint64(0)

	// obj.Sig
	i0 += 65

	return i0
}

// encodeSigWrapper encodes an object of type sigWrapper to a buffer allocated to the exact size
// required to encode the object.
func encodeSigWrapper(obj *sigWrapper) ([]byte, error) {
	n := encodeSizeSigWrapper(obj)
	buf := make([]byte, n)

	if err := encodeSigWrapperToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeSigWrapperToBuffer encodes an object of type sigWrapper to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeSigWrapperToBuffer(buf []byte, obj *sigWrapper) error {
	if uint64(len(buf)) < encodeSizeSigWrapper(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Sig
	e.CopyBytes(obj.Sig[:])

	return nil
}

// decodeSigWrapper decodes an object of type sigWrapper from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeSigWrapper(buf []byte, obj *sigWrapper) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Sig
		if len(d.Buffer) < len(obj.Sig) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Sig[:], d.Buffer[:len(obj.Sig)])
		d.Buffer = d.Buffer[len(obj.Sig):]
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeSigWrapperExact decodes an object of type sigWrapper from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeSigWrapperExact(buf []byte, obj *sigWrapper) error {
	if n, err := decodeSigWrapper(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
