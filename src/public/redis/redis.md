
redis-cli -h  127.0.0.1 -p 6379



SUBSCRIBE testchan  订阅

PUBLISH testchan "hello"  推送


ReceiveMessage()

SUBSCRIBE options:pub:instrument:BTCUSDT


options_kline_


ZSliceCmd
集合
func NewZSliceCmd(args ...interface{}) *ZSliceCmd
func NewZSliceCmdResult(val []Z, err error) *ZSliceCmd
func (cmd *ZSliceCmd) Args() []interface{}
func (cmd *ZSliceCmd) Err() error
func (cmd *ZSliceCmd) Name() string
func (cmd *ZSliceCmd) Result() ([]Z, error)
func (cmd *ZSliceCmd) String() string
func (cmd *ZSliceCmd) Val() []Z

https://blog.csdn.net/comprel/article/details/96716582


Sorted Set 有序集合
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]


# REDIS
REDIS_ADDR=47.115.143.171:16379


ZRANGEBYSCORE options:kline:cache:BTCUSDT-20210121-35000-C:60 -inf +inf