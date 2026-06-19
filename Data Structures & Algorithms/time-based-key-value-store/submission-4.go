type TimeValue struct {
    value     string
    timestamp int
}

type TimeMap struct {
	keyStore map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{
        keyStore: make(map[string][]TimeValue),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keyStore[key] = append(this.keyStore[key], TimeValue{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	l := 0
    r := len(this.keyStore[key]) - 1
    res := ""

    for l <= r {
        mid := (l + r) / 2
        currentTimestamp := this.keyStore[key][mid].timestamp
        currentValue := this.keyStore[key][mid].value

        if currentTimestamp == timestamp {
            return currentValue
        } else if currentTimestamp < timestamp {
            l = mid + 1
            res = currentValue
        } else {
            r = mid - 1
        }
    }

    return res
}
