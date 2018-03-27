package main

import (
    "time"
    "encoding/json"
)

func GetTimeByteArray() ([]byte, error) {
    return json.Marshal(time.Now())
}

