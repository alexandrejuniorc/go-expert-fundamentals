package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		// handle cancelation
		fmt.Println("Hotel booking is canceled. Timeout reached")
		return
	//case <-time.After(5 * time.Second):
	case <-time.After(1 * time.Second):
		// book hotel
		fmt.Println("Hotel booked.")
	}
}
