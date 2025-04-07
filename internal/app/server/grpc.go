package grpcserver

import (
	"log"
	"sync"
	"time"

	candle "github.com/AkyurekDogan/exinity-task/internal/app/proto"
)

type CandleServiceServer struct {
	candle.UnimplementedCandleServiceServer
	Subscribers map[string][]chan *candle.Candle // symbol → list of subscriber channels
	mu          sync.Mutex
}

func NewCandleServiceServer() *CandleServiceServer {
	return &CandleServiceServer{
		Subscribers: make(map[string][]chan *candle.Candle),
	}
}

func (s *CandleServiceServer) SubscribeCandles(req *candle.SubscribeRequest, stream candle.CandleService_SubscribeCandlesServer) error {
	symbols := req.GetSymbols()
	if len(symbols) == 0 {
		symbols = []string{"BTCUSDT", "ETHUSDT", "PEPEUSDT"}
	}

	// Create channels and subscribe the client
	channels := make([]chan *candle.Candle, len(symbols))
	for i, sym := range symbols {
		ch := make(chan *candle.Candle, 100) // Buffered channel to prevent blocking
		channels[i] = ch

		s.mu.Lock()
		s.Subscribers[sym] = append(s.Subscribers[sym], ch)
		s.mu.Unlock()
	}

	// Stream candles to the client
	for {
		select {
		case <-stream.Context().Done():
			// Handle client disconnect
			return nil
		default:
			for _, ch := range channels {
				select {
				case c := <-ch:
					if err := stream.Send(c); err != nil {
						log.Println("Failed to send candle:", err)
						return err
					}
				default:
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *CandleServiceServer) BroadcastCandle(c *candle.Candle) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Broadcast to all subscribers of this symbol
	for _, ch := range s.Subscribers[c.Symbol] {
		// Non-blocking send to prevent slow clients from blocking
		select {
		case ch <- c:
		default:
			// If the client is not ready, skip to avoid blocking
			log.Printf("Dropping candle for %s due to slow client\n", c.Symbol)
		}
	}
}
