// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"sync"
// )

// func makeThumbnails(filenames []string) {
// 	for _, f := range filenames {
// 		if _, err := thumbnail.ImageFile(f); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

// func makeThumbnails2(filenames []string) {
// 	for _, f := range filenames {
// 		go thumbnail.ImageFile(f)
// 	}
// }

// func makeThumbnails3(filenames []string) {
// 	ch := make(chan error)
// 	for _, f := range filenames {
// 		go func(f string) {
// 			_, err := thumbnail.ImageFile(f)
// 			ch <- err
// 		}(f)
// 	}
// 	for range filenames {
// 		<-ch
// 	}
// }

// func makeThumbnails4(filenames []string) error {
// 	ch := make(chan error)
// 	for _, f := range filenames {
// 		go func(f string) {
// 			_, err := thumbnail.ImageFile(f)
// 			ch <- err
// 		}(f)
// 	}
// 	for range filenames {
// 		if err := <-ch; err != nil {
// 			return err // NOTE: incorrect: goroutine leak!
// 		}
// 	}
// 	return nil
// }

// func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
// 	type item struct {
// 		thumbfile string
// 		err       error
// 	}
// 	ch := make(chan item, len(filenames))
// 	for _, f := range filenames {
// 		go func(f string) {
// 			var it item
// 			it.thumbfile, it.err = thumbnail.ImageFile(f)
// 			ch <- it
// 		}(f)
// 	}
// 	for range filenames {
// 		it := <-ch
// 		if it.err != nil {
// 			return nil, it.err
// 		}
// 		thumbfiles = append(thumbfiles, it.thumbfile)
// 	}
// 	return thumbfiles, nil
// }

// func makeThumbnails6(filenames <-chan string) int64 {
// 	sizes := make(chan int64)
// 	var wg sync.WaitGroup
// 	for f := range filenames {
// 		wg.Add(1)
// 		go func(f string) {
// 			defer wg.Done()
// 			thumb, err := thumbnail.ImageFile(f)
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}
// 			info, _ := os.Stat(thumb) // OK to ignore error
// 			sizes <- info.Size()
// 		}(f)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(sizes)
// 	}()
// 	var total int64
// 	for size := range sizes {
// 		total += size
// 	}
// 	return total
// }

// func main() {
// 	// Create a channel for passing filenames
// 	filenames := make(chan string)

// 	// Start a goroutine to send filenames
// 	go func() {
// 		// Note: Close the channel when done sending
// 		defer close(filenames)

// 		// Mock some filenames
// 		files := []string{
// 			"images/photo1.jpg",
// 			"images/photo2.jpg",
// 			"images/photo3.jpg",
// 			"images/photo4.jpg",
// 		}

// 		// Send filenames to channel
// 		for _, filename := range files {
// 			filenames <- filename
// 		}
// 	}()

// 	// Call makeThumbnails6 and get result
// 	totalSize := makeThumbnails6(filenames)

// 	// Print result
// 	fmt.Printf("Total size of thumbnails: %d bytes\n", totalSize)
// }
