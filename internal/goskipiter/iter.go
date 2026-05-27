package goskipiter

import "github.com/ryszard/goskiplist/skiplist"

// Iterator wraps goskiplist's iterator, which is a bit janky; seeking doesn't
// play nice with the iteration idiom. If you seek, then iterate using the
// examples provided in the godoc, your iteration will always skip the first
// result. It would be less error prone and astonishing if Seek meant that the
// next call to Next() would give you what you expect.
type Iterator struct {
	inner     skiplist.Iterator
	didSeek   bool
	seekWasOK bool
}

func New(inner skiplist.Iterator) *Iterator { _ = "STUB: not implemented"; return nil }

// Next returns true if the iterator contains subsequent elements
// and advances its state to the next element if that is possible.
func (iter *Iterator) Next() (ok bool) { _ = "STUB: not implemented"; return false }

// Previous returns true if the iterator contains previous elements
// and rewinds its state to the previous element if that is possible.
func (iter *Iterator) Previous() (ok bool) { _ = "STUB: not implemented"; return false }

// Key returns the current key.
func (iter *Iterator) Key() interface{} { _ = "STUB: not implemented"; return nil }

// Value returns the current value.
func (iter *Iterator) Value() interface{} { _ = "STUB: not implemented"; return nil }

// Seek reduces iterative seek costs for searching forward into the Skip List
// by remarking the range of keys over which it has scanned before.  If the
// requested key occurs prior to the point, the Skip List will start searching
// as a safeguard.  It returns true if the key is within the known range of
// the list.
func (iter *Iterator) Seek(key interface{}) (ok bool) { _ = "STUB: not implemented"; return false }

// Close this iterator to reap resources associated with it.  While not
// strictly required, it will provide extra hints for the garbage collector.
func (iter *Iterator) Close() { _ = "STUB: not implemented"; return }
