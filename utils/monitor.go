package utils

import (
	"fmt"
	"strconv"
	"time"

	"gitee.com/dark.H/gs"
)

func WatchIfExists(tg *TgAuth, pid_or_process_key gs.Str) {
	if pid, err := strconv.Atoi(pid_or_process_key.String()); err == nil {
		go func(pid int) {
			tick := time.NewTicker(time.Second * 1)
		label1:
			for {
				select {
				case <-tick.C:
					// 检查进程是否存在
					out := gs.Str("ps aux  | awk '{ print $2}' | grep -v grep | grep " + strconv.Itoa(pid)).Exec().Trim()
					if out.String() == "" {
						break label1
					}
				}
			}
			tg.Say(fmt.Sprint(pid) + "进程已退出")
		}(pid)
	} else {
		go func(pid int) {
			tick := time.NewTicker(time.Second * 1)
		label:
			for {
				select {
				case <-tick.C:
					// 检查进程是否存在
					out := gs.Str("ps aux  | grep -v grep | grep " + pid_or_process_key.String()).Exec().Trim()
					if out.String() == "" {
						break label
					}
				}
			}
			tg.Say(pid_or_process_key.String() + "进程已退出")
		}(pid)
	}
}
