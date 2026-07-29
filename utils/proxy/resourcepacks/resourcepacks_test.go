package resourcepacks

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func TestOnResourcePackStackStopsWhenContextIsCancelled(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	handler := NewResourcePackHandler(ctx, nil)
	handler.clientDone = make(chan struct{})

	want := errors.New("connection closed")
	cancel(want)

	err := handler.OnResourcePackStack(&packet.ResourcePackStack{})
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

func TestOnResourcePackChunkDataStopsWhenContextIsCancelled(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	handler := NewResourcePackHandler(ctx, nil)
	handler.awaitingPack = &downloadingPack{
		ID:      uuid.New(),
		newFrag: make(chan *packet.ResourcePackChunkData),
	}

	want := errors.New("connection closed")
	cancel(want)

	err := handler.OnResourcePackChunkData(&packet.ResourcePackChunkData{})
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}
