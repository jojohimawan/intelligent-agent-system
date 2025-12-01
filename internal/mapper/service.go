package mapper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/jojohimawan/intelligent-agent-system/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

type MappingEntry struct {
	VSSPath string
	Unit    string
	Type    string
}

type SignalConfig struct {
	Protocol   string `bson:"protocol" json:"protocol"`
	SignalName string `bson:"signal_name" json:"signal_name"`
	VSS        struct {
		Path string `bson:"path" json:"path"`
		Unit string `bson:"unit" json:"unit"`
		Type string `bson:"type" json:"type"`
	} `bson:"vss" json:"vss"`
}

type Service struct {
	mu       sync.RWMutex
	cache    map[string]MappingEntry
	filePath string
	store    *MongoStore
}

func NewService(store *MongoStore, filePath string) *Service {
	return &Service{
		cache:    make(map[string]MappingEntry),
		store:    store,
		filePath: filePath,
	}
}

func (s *Service) Init(ctx context.Context) error {
	configs, err := s.fetchFromMongo(ctx)
	if err == nil {
		log.Printf("[INFO]Mapper: successfully fetch %d rules from MongoDB", len(configs))

		s.updateCache(configs)

		go func() {
			if err := s.saveToDisk(configs); err != nil {
				log.Printf("[WARN]Mapper: failed to backup to disk: %v", err)
			} else {
				log.Printf("[INFO]Mapper: backup successfully saved to disk")
			}
		}()
		return nil
	}

	log.Printf("[INFO]Mapper: MongoDB unreachable (%v). Trying offline fallback...", err)
	if err := s.loadFromDisk(); err != nil {
		return fmt.Errorf("[ERR]Mapper: Couldn't load mappings from network or disk: %w", err)
	}

	log.Printf("[INFO]Mapper: Loaded %d rules from local disk cache", len(s.cache))
	return nil
}

func (s *Service) Translate(in *models.DecodedSignal) (*models.VSSPoint, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	key := in.Source + ":" + in.Param
	entry, exists := s.cache[key]
	if !exists {
		return nil, false
	}

	return &models.VSSPoint{
		Path:  entry.VSSPath,
		Value: in.Value,
		Unit:  entry.Unit,
	}, true
}

func (s *Service) updateCache(configs []SignalConfig) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache = make(map[string]MappingEntry)
	for _, cfg := range configs {
		key := cfg.Protocol + ":" + cfg.SignalName
		s.cache[key] = MappingEntry{
			VSSPath: cfg.VSS.Path,
			Unit:    cfg.VSS.Unit,
			Type:    cfg.VSS.Type,
		}
	}
}

func (s *Service) fetchFromMongo(ctx context.Context) ([]SignalConfig, error) {
	if s.store == nil {
		return nil, fmt.Errorf("[ERR]Mapper: Mongo store is nil")
	}

	cursor, err := s.store.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []SignalConfig
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Service) saveToDisk(configs []SignalConfig) error {
	data, err := json.MarshalIndent(configs, "", "")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.filePath)
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp(dir, "mappings-*.tmp")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		return err
	}

	if err := tmpFile.Sync(); err != nil {
		return err
	}

	if err := tmpFile.Close(); err != nil {
		return err
	}

	return os.Rename(tmpFile.Name(), s.filePath)
}

func (s *Service) loadFromDisk() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var configs []SignalConfig
	if err := json.NewDecoder(file).Decode(&configs); err != nil {
		return err
	}

	s.updateCache(configs)
	return nil
}
