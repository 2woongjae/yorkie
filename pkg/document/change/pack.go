package change

import (
	"github.com/hackerwins/rottie/pkg/document/checkpoint"
	"github.com/hackerwins/rottie/pkg/document/key"
)

type Pack struct {
	DocumentKey *key.Key
	Checkpoint  *checkpoint.Checkpoint
	Changes     []*Change
}

func NewPack(
	key *key.Key,
	cp *checkpoint.Checkpoint,
	changes []*Change,
) *Pack {
	return &Pack{
		DocumentKey: key,
		Checkpoint:  cp,
		Changes:     changes,
	}
}