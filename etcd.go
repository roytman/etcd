package etcd

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/fperf/fperf"
	"go.etcd.io/etcd/clientv3"
)

const validKeyValueBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	fperf.Register("etcd", New, "etcd benchmark")
}

// Op is the operation type issued to etcd
type Op string

// Operations
const (
	Put    Op = "put"
	Get    Op = "get"
	Range  Op = "range"
	Delete Op = "delete"
	Trx    Op = "trx"
)

type client struct {
	etcd    *clientv3.Client
	keySpace   *randSpace
	valueSpace   *randSpace
	op      Op
	ops  []clientv3.Op
}

// New creates a fperf client
func New(fs *fperf.FlagSet) fperf.Client {
	var keySize,valueSize,trxSize int
	var op Op
	fs.IntVar(&keySize, "key-size", 8, "length of the random key")
	fs.IntVar(&valueSize, "value-size", 8, "length of the random value")
	fs.IntVar(&trxSize, "trx-size", 128, "number of operations in the transaction")
	fs.Parse()
	args := fs.Args()
	if len(args) == 0 {
		op = Put
	} else {
		op = Op(args[0])
		if len(args) == 2 {
			trxSize, _ = strconv.Atoi(args[1])
		}
	}
	c := &client{
		keySpace:   newRandSpace(keySize),
		valueSpace:   newRandSpace(valueSize),
		op:      op,
	}
	if c.op == Trx {
		var ops  []clientv3.Op
		keys := map[string]string{}
		for i := 0; i< trxSize; i++ {
			key := c.keySpace.randString()
			if _, ok := keys[key]; ok {
				continue
			}
			keys[key] = key
			value := c.valueSpace.randString()
			ops = append(ops, clientv3.OpPut(key, value))
		}
		c.ops = ops
	}
	return c
}

// Dial to etcd
func (c *client) Dial(addr string) error {
	endpoints := strings.Split(addr, ",")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 2 * time.Second,
		MaxCallSendMsgSize: 16 * 1024 * 1024 * 1024,
	})
	if err != nil {
		return err
	}
	c.etcd = cli
	return nil
}

// Request etcd
func (c *client) Request() error {
	switch c.op {
	case Put:
		return doPut(c)
	case Get:
		return doGet(c)
	case Range:
		return doRange(c)
	case Delete:
		return doDelete(c)
	case Trx:
		return doTrx(c)
	}
	return fmt.Errorf("unknown op %s", c.op)
}

func doPut(c *client) error {
	key := c.keySpace.randString()
	value := c.valueSpace.randString()
	_, err := c.etcd.Put(context.Background(), key, value)
	return err
}
func doGet(c *client) error {
	_, err := c.etcd.Get(context.Background(), c.keySpace.randString())
	return err
}
func doDelete(c *client) error {
	_, err := c.etcd.Delete(context.Background(), c.keySpace.randString())
	return err
}
func doRange(c *client) error {
	start, end := c.keySpace.randRange()
	_, err := c.etcd.Get(context.Background(), start, clientv3.WithRange(end))
	return err
}

func doTrx(c *client) error {
	//fmt.Printf("size %d\n", len(ops))
	_, err := c.etcd.Txn(context.Background()).Then(c.ops ...).Commit()
	//fmt.Printf("Succeeded %v , size %d\n", resp.Succeeded, len(resp.Responses))
	return err
}

type randSpace struct {
	r      *rand.Rand
	nbytes int
}

func newRandSpace(nbytes int) *randSpace {
	return &randSpace{
		r:      rand.New(rand.NewSource(time.Now().Unix())),
		nbytes: nbytes,
	}
}

func (ks *randSpace) randString() string {
	p := make([]byte, ks.nbytes)
    for i := range p {
        p[i] = validKeyValueBytes[rand.Intn(len(validKeyValueBytes))]
    }
	return string(p)
}
func (ks *randSpace) randRange() (string, string) {
	start := []byte(ks.randString())
	if len(start) == 0 {
		return "", ""
	}
	// the max range is 256 according to the last bytes of start
	end := make([]byte, len(start))
	copy(end, start)
	end[len(end)-1] = 0xFF
	return string(start), string(end)
}
