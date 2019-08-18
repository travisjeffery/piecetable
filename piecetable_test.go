package piecetable_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	writegood "github.com/travisjeffery/writegood"
)

func TestInsert_Empty(t *testing.T) {
	t.Skip()

	d := &writegood.PieceTable{}
	req := require.New(t)

	// initial case, empty doc
	d.Insert(0, []byte("hello"))
	t.Logf("added: %s\n", d.Add)
	req.Equal([]byte("hello"), d.Add)
	req.Equal([]*writegood.Piece{{
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Add,
	}}, d.Pieces)
	b, _ := d.Bytes()
	req.Equal([]byte("hello"), b)
	t.Logf("added: %s\n", d.Add)

	// add to start
	d.Insert(0, []byte("world"))
	req.Equal([]*writegood.Piece{{
		Start:  len("hello"),
		Length: len("world"),
		Type:   writegood.Add,
	}, {
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Add,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("worldhello"), b)

	// add to middle
	d.Insert(len("world"), []byte("bye"))
	t.Logf("added: %s\n", d.Add)
	req.Equal([]*writegood.Piece{{
		Start:  len("hello"),
		Length: len("world"),
		Type:   writegood.Add,
	}, {
		Start:  len("helloworld"),
		Length: len("bye"),
		Type:   writegood.Add,
	}, {
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Add,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("worldbyehello"), b)

	d.Insert(len("wor"), []byte("yes"))
	t.Logf("added: %s\n", d.Add)
	req.Equal([]*writegood.Piece{{
		Start:  len("hello"),
		Length: len("wor"),
		Type:   writegood.Add,
	}, {
		Start:  len("helloworldbye"),
		Length: len("yes"),
		Type:   writegood.Add,
	}, {
		Start:  len("hellowor"),
		Length: len("ld"),
		Type:   writegood.Add,
	}, {
		Start:  len("helloworld"),
		Length: len("bye"),
		Type:   writegood.Add,
	}, {
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Add,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("woryesldbyehello"), b)

	// add to end
	d.Insert(len("woryesldbyehello"), []byte("grief"))
	t.Logf("added: %s\n", d.Add)
	req.Equal([]*writegood.Piece{{
		Start:  len("hello"),
		Length: len("wor"),
		Type:   writegood.Add,
	}, {
		Start:  len("helloworldbye"),
		Length: len("yes"),
		Type:   writegood.Add,
	}, {
		Start:  len("hellowor"),
		Length: len("ld"),
		Type:   writegood.Add,
	}, {
		Start:  len("helloworld"),
		Length: len("bye"),
		Type:   writegood.Add,
	}, {
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Add,
	}, {
		Start:  len("woryesldbyehello"),
		Length: len("grief"),
		Type:   writegood.Add,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("woryesldbyehellogrief"), b)
}

func TestInsert_Existing(t *testing.T) {
	t.Skip()

	req := require.New(t)
	d := &writegood.PieceTable{
		Original: []byte("helloworld"),
		Pieces: []*writegood.Piece{{
			Start:  0,
			Length: len("helloworld"),
			Type:   writegood.Original,
		}},
	}
	d.Insert(len("hello"), []byte("earth"))
	t.Logf("original: %s, added: %s\n", d.Original, d.Add)
	req.Equal([]*writegood.Piece{{
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Original,
	}, {
		Start:  0,
		Length: len("earth"),
		Type:   writegood.Add,
	}, {
		Start:  len("hello"),
		Length: len("world"),
		Type:   writegood.Original,
	}}, d.Pieces)
	b, _ := d.Bytes()
	req.Equal([]byte("helloearthworld"), b)
}

func TestDelete(t *testing.T) {
	req := require.New(t)
	d := &writegood.PieceTable{
		Original: []byte("helloworld"),
		Pieces: []*writegood.Piece{{
			Start:  0,
			Length: len("helloworld"),
			Type:   writegood.Original,
		}},
	}
	d.Insert(len("hello"), []byte("earth"))
	b, _ := d.Bytes()
	req.Equal([]byte("helloearthworld"), b)

	// delete whole earth piece
	d.Delete(len("hello"), len("helloearth"))
	req.Equal([]*writegood.Piece{{
		Start:  0,
		Length: len("hello"),
		Type:   writegood.Original,
	}, {
		Start:  len("hello"),
		Length: len("world"),
		Type:   writegood.Original,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("helloworld"), b)

	// delete part of piece
	d.Delete(len("hel"), len("hello"))
	req.Equal([]*writegood.Piece{{
		Start:  0,
		Length: len("hel"),
		Type:   writegood.Original,
	}, {
		Start:  len("hello"),
		Length: len("world"),
		Type:   writegood.Original,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("helworld"), b)

	// delete middle of piece
	d.Delete(len("helwor"), len("helworl"))
	req.Equal([]*writegood.Piece{{
		Start:  0,
		Length: len("hel"),
		Type:   writegood.Original,
	}, {
		Start:  len("hello"),
		Length: len("wor"),
		Type:   writegood.Original,
	}, {
		Start:  len("helloworl"),
		Length: len("d"),
		Type:   writegood.Original,
	}}, d.Pieces)
	b, _ = d.Bytes()
	req.Equal([]byte("helword"), b)
}