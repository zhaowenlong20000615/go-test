-- 限流对象
local key = KEYS[1]
-- 窗口大小
local window = tonumber(ARGV[1])
-- 阈值
local threshold = tonumber(ARGV[2])
-- 当前时间
local now = tonumber(ARGV[3])
-- 窗口的起始时间
local min = now - window

-- 从有序集合中删除所有小于窗口起始时间的元素，
-- -inf：表示负无穷大，即有序集合中最小的分数值。
-- +inf：表示正无穷大，即有序集合中最大的分数值。
-- 这样可以确保有序集合中只包含当前时间窗口内的请求时间戳。
redis.call('ZREMRANGEBYSCORE', key, '-inf', min)
-- 计算有序集合中的元素数量（即当前时间窗口内的请求数量）
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf')
-- 判断请求数量是否超过阈值
if cnt >= threshold then
    -- 执行限流
    return "true"
else
    -- 将当前时间戳添加到有序集合中
    -- 把 score 和 member 都设置成 now
    redis.call('ZADD', key, now, now)
    -- 设置有序集合的过期时间
    redis.call('PEXPIRE', key, window)
    return "false"
end