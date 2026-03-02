package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func LoadDealer(path string) (dealer *Dealer, err error) {
	start := time.Now()
	defer func() {
		log.Printf("[loader] loading %q took %v, err: %v", path, time.Since(start), err)
	}()

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	var root Root
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&root); err != nil {
		// this is the panic that wwe recover from if we use SafeLoad
		// panic("meow")
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}

	d := &root.Dealer
	// added this here
	var errs []error
	if errs = d.Validate(); errs != nil {
		return nil, fmt.Errorf("validating data: %w", err)
	}
	return d, nil
}

func SaveReport(path string, content string) (nil, err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening file for writing: %w", err)
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	defer log.Println(err)
	return nil, err
}

func SafeLoad(path string) (d *Dealer, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("meow")

			log.Println(err)
		}
	}()
	return LoadDealer(path)
}

// Q: what if the fn has to return something?
func withTimer(name string, fn func()) {
	start := time.Now()
	defer func() {
		log.Printf("[function] %s took %v to execute", name, time.Since(start))
	}()
	fn()
}
