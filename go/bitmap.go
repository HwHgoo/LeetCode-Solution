package leetcode

// only works for non-negtive numbers
type Bitmap struct {
	bytes []byte
}

type BitmapIter struct {
	m   *Bitmap
	cur int
	max int
}

func NewBitmap(bits int) *Bitmap {
	size := bits / 8
	if bits%8 != 0 {
		size += 1
	}

	return &Bitmap{bytes: make([]byte, size)}
}

func (bm *Bitmap) Set(n int) {
	// bm.bytes[n/8] |= (1 << (n % 8))
	bm.bytes[n>>3] |= (1 << (n & 7))
}

func (bm *Bitmap) Get(n int) bool {
	// return bm.bytes[n/8]&(1<<(n%8)) != 0
	return bm.bytes[n>>3]&(1<<(n&7)) != 0
}

func (bm *Bitmap) Iterator() *BitmapIter {
	return &BitmapIter{m: bm, cur: -1, max: len(bm.bytes)*8 - 1}
}

func (iter *BitmapIter) Next() (int, bool) {
	if iter.cur >= iter.max {
		return -1, false
	}
	iter.cur++
	return iter.cur, iter.m.Get(iter.cur)
}

func (iter *BitmapIter) NextUnset() int {
	if iter.cur == -1 {
		iter.cur++
	}

	for iter.cur < iter.max && iter.m.Get(iter.cur) {
		iter.cur++
	}

	return iter.cur
}

func (iter *BitmapIter) NextSet() int {
	if iter.cur == -1 {
		iter.cur++
	}

	for iter.cur < iter.max && !iter.m.Get(iter.cur) {
		iter.cur++
	}

	return iter.cur
}
