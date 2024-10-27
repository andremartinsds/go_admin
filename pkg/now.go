package pkg

import "time"

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
	//return now.Format("2006-01-02 15:04:05")
}
