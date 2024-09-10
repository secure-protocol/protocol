package bloom

import (
	"context"
	_ "embed"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/spaolacci/murmur3"
	"math"
	"strconv"
)

var (

	//go:embed setscript.lua
	setLuaScript string
	setScript    = redis.NewScript(setLuaScript)

	//go:embed testscript.lua
	testLuaScript string
	testScript    = redis.NewScript(testLuaScript)
)

type Filter struct {
	client   *redis.Client
	size     uint
	hashIter uint
	mem      []uint
	key      string
}

// OptimalK calculates the optimal k value for creating a new Bloom filter
// maxn is the maximum anticipated number of elements
func OptimalK(m, maxN uint) uint {
	return uint(math.Ceil(float64(m) * math.Ln2 / float64(maxN)))
}

// OptimalM calculates the optimal m value for creating a new Bloom filter
// p is the desired false positive probability
// optimal m = ceiling( - n * ln(p) / ln(2)**2 )
func OptimalM(maxN uint, p float64) uint {
	return uint(math.Ceil(-float64(maxN) * math.Log(p) / (math.Ln2 * math.Ln2)))
}

func NewFilter(ctx context.Context, client *redis.Client, key string, maxElement uint, p float64) (*Filter, error) {
	m := OptimalM(maxElement, p)
	k := OptimalK(m, maxElement)

	l := client.ScriptLoad(ctx, setLuaScript)
	if err := l.Err(); err != nil {
		return nil, err
	}
	l = client.ScriptLoad(ctx, testLuaScript)
	if err := l.Err(); err != nil {
		return nil, err
	}

	return &Filter{
		client:   client,
		size:     m,
		hashIter: k,
		mem:      nil,
		key:      key,
	}, nil
}

func (f *Filter) Clear(ctx context.Context) error {
	del := f.client.Del(ctx, f.key)
	return del.Err()
}

func (f *Filter) Add(ctx context.Context, data []byte) error {
	locations := f.getLocations(data)
	cmd := setScript.EvalSha(ctx, f.client, []string{f.key}, locations)
	if errors.Is(cmd.Err(), redis.Nil) {
		return nil
	}
	return cmd.Err()
}

func (f *Filter) BatchAdd(ctx context.Context, dataList ...[]byte) error {
	locations := make([]string, 0, len(dataList)*int(f.hashIter))
	for _, bytes := range dataList {
		locations = append(locations, f.getLocations(bytes)...)
	}
	cmd := setScript.EvalSha(ctx, f.client, []string{f.key}, locations)
	if errors.Is(cmd.Err(), redis.Nil) {
		return nil
	}
	return cmd.Err()
}

func (f *Filter) Exist(ctx context.Context, data []byte) (bool, error) {
	locations := f.getLocations(data)
	cmd := testScript.EvalSha(ctx, f.client, []string{f.key}, locations)
	if errors.Is(cmd.Err(), redis.Nil) {
		return false, nil
	}
	return cmd.Bool()
}

func (f *Filter) getLocations(data []byte) []string {
	locations := make([]string, f.hashIter)
	for i := uint(0); i < f.hashIter; i++ {
		hashValue := hash64(append(data, byte(i)))
		locations[i] = strconv.FormatUint(uint64(hashValue%uint64(f.size)), 10)
	}

	return locations
}

func hash64(in []byte) uint64 {
	return murmur3.Sum64(in)
}
