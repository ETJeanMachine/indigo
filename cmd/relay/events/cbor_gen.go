// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package events

import (
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *EventHeader) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.MsgType (string) (string)
	if len("t") > 1000000 {
		return xerrors.Errorf("Value in field \"t\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("t"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("t")); err != nil {
		return err
	}

	if len(t.MsgType) > 1000000 {
		return xerrors.Errorf("Value in field t.MsgType was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.MsgType))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.MsgType)); err != nil {
		return err
	}

	// t.Op (int64) (int64)
	if len("op") > 1000000 {
		return xerrors.Errorf("Value in field \"op\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("op"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("op")); err != nil {
		return err
	}

	if t.Op >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Op)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Op-1)); err != nil {
			return err
		}
	}

	return nil
}

func (t *EventHeader) UnmarshalCBOR(r io.Reader) (err error) {
	*t = EventHeader{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("EventHeader: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 2)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 1000000)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.MsgType (string) (string)
		case "t":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.MsgType = string(sval)
			}
			// t.Op (int64) (int64)
		case "op":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.Op = int64(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
func (t *ErrorFrame) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.Error (string) (string)
	if len("error") > 1000000 {
		return xerrors.Errorf("Value in field \"error\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("error"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("error")); err != nil {
		return err
	}

	if len(t.Error) > 1000000 {
		return xerrors.Errorf("Value in field t.Error was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Error))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Error)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("message") > 1000000 {
		return xerrors.Errorf("Value in field \"message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("message"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("message")); err != nil {
		return err
	}

	if len(t.Message) > 1000000 {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *ErrorFrame) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ErrorFrame{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ErrorFrame: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 7)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 1000000)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.Error (string) (string)
		case "error":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Error = string(sval)
			}
			// t.Message (string) (string)
		case "message":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
