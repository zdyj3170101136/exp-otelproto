package grpc_stream

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client traceprotobuf.StreamTracerClient
	stream traceprotobuf.StreamTracer_SendBatchClient
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewStreamTracerClient(conn)

	// Establish stream to server.
	c.stream, err = c.client.SendBatch(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Export(batch core.SpanBatch) {
	// Send the batch via stream.
	c.stream.Send(batch.(*traceprotobuf.SpanBatch))

	// Wait for response from server. This is full synchronous operation,
	// we do not send batches concurrently.
	_, err := c.stream.Recv()

	if err != nil {
		log.Fatal("Error from server when expecting batch response")
	}
}
