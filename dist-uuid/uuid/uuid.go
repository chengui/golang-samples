package uuid

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"

	flake "github.com/sony/sonyflake"
	redis "github.com/go-redis/redis/v8"
)

type Options struct {
	EpochTime string
	LocalFile string
	RedisAddr string
	RedisPass string
	RedisDB   int
}

type UUID struct {
	Assigner *Assigner
	Flaker   *flake.Sonyflake
}

func NewUUID(opts *Options) *UUID {
	assinger := NewAssigner(opts)
	startTime, _ := time.Parse("2006-01-02", opts.EpochTime)
	settings := flake.Settings{
		StartTime:      startTime,
		MachineID:      assinger.GetMachineID,
		CheckMachineID: assinger.CheckMachinID,
	}
	flaker := flake.NewSonyflake(settings)
	return &UUID{
		Flaker:   flaker,
		Assigner: assinger,
	}
}

func (u *UUID) Generate() uint64 {
	ID, _ := u.Flaker.NextID()
	return ID
}

type Assigner struct {
	RedisCli  *redis.Client
	LocalFile string
}

func NewAssigner(opts *Options) *Assigner {
	redisOpt := &redis.Options{
		Addr:     opts.RedisAddr,
		DB:       opts.RedisDB,
		Password: opts.RedisPass,
	}
	redisCli := redis.NewClient(redisOpt)
	return &Assigner{
		RedisCli:  redisCli,
		LocalFile: opts.LocalFile,
	}
}

func (a *Assigner) GetMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID, err = a.ReadFromFile()
	if err != nil {
		machineID, err = a.GenerateMachineID()
		if err != nil {
			return 0, err
		}
	}
	return machineID, err
}

func (a *Assigner) CheckMachinID(machineID uint16) bool {
	res, err := a.AddToRedisSet(machineID)
	if err != nil || res == 0 {
		return false
	}
	err = a.WriteToFile(machineID)
	if err != nil {
		return false
	}
	return true
}

func (a *Assigner) ReadFromFile() (uint16, error) {
	data, err := ioutil.ReadFile(a.LocalFile)
	if err != nil {
		return 0, err
	}
	idStr := strings.TrimSpace(string(data[:]))
	machineID, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(machineID), nil
}

func (a *Assigner) WriteToFile(machineID uint16) error {
	idStr := strconv.FormatInt(int64(machineID), 10)
	err := ioutil.WriteFile(a.LocalFile, []byte(idStr), 0666)
	return err
}

func (a *Assigner) AddToRedisSet(machineID uint16) (int64, error) {
	if a.RedisCli == nil {
		return 0, fmt.Errorf("redis not connected")
	}
	ctx := context.TODO()
	idStr := strconv.FormatInt(int64(machineID), 10)
	val, err := a.RedisCli.SAdd(ctx, "uuid-nodes", idStr).Result()
	fmt.Println(val, err)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (a *Assigner) GenerateMachineID() (uint16, error) {
	ip, err := PrivateIPv4()
	if err != nil {
		return 0, err
	}
	return GenHashCode(ip), nil
}
