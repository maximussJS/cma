package services

import (
	"cma/packages/constants"
	"cma/packages/models"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

type State struct {
	fs *Filesystem
}

var state *State

var mutexes = map[constants.BlockchainType]*sync.Mutex{
	constants.Base:     &sync.Mutex{},
	constants.Ethereum: &sync.Mutex{},
}

func StateSingletonInstance() *State {
	if state != nil {
		return state
	}

	filesystem, err := NewFilesystem()

	if err != nil {
		log.Fatalf("cannot init filesystem %w", err)
	}

	state = &State{
		fs: filesystem,
	}

	return state
}

func getFileName(blockchain constants.BlockchainType) string {
	return filepath.Join(blockchain.String() + "_last_block.json")
}

func (s *State) getLock(blockchain constants.BlockchainType) (*sync.Mutex, error) {
	if lock, exists := mutexes[blockchain]; exists {
		return lock, nil
	}

	return nil, fmt.Errorf("State.getLock() cannot get lock for blockchain %s", blockchain)
}

func (s *State) Set(blockchain constants.BlockchainType, block *models.StateBlock) error {
	lock, err := s.getLock(blockchain)

	if err != nil {
		return fmt.Errorf("State.Set() get lock error %w", err)
	}

	lock.Lock()
	defer lock.Unlock()

	stateData, err := json.Marshal(block)
	if err != nil {
		return fmt.Errorf("State.Set() marhal error: %w", err)
	}

	filename := getFileName(blockchain)

	if err := s.fs.WriteString(filename, string(stateData)); err != nil {
		return fmt.Errorf("State.Set() write error: %w", err)
	}

	return nil
}

func (s *State) Get(blockchain constants.BlockchainType) (*models.StateBlock, error) {
	lock, err := s.getLock(blockchain)

	if err != nil {
		return nil, fmt.Errorf("State.Get() get lock error %w", err)
	}

	lock.Lock()

	defer lock.Unlock()

	filename := getFileName(blockchain)

	data, err := s.fs.ReadString(filename)

	if err != nil {
		return nil, fmt.Errorf("State.Get() read string error %w", err)
	}

	if data == "" {
		return nil, nil
	}

	var state models.StateBlock

	if err := json.Unmarshal([]byte(data), &state); err != nil {
		return nil, fmt.Errorf("State.Get() failed to unmarshal state: %w", err)
	}

	return &state, nil
}
