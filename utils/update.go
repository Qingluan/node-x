package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"gitee.com/dark.H/gs"
)

func Uprade(file, oldFile string) {
	fpath := gs.Str(file)
	if fpath.IsExists() {
		testStr := gs.Str("chmod +x " + fpath.Str() + " && " + fpath.Str() + " -h").Exec()
		if testStr.In(`-update_me_by_updater_this_is_a_tag_has_no_meaning`) {
			gs.Str("test ok").Color("g").Println("Upgrade")
			if err := gs.Str(oldFile).Rm(); err != nil {
				gs.Str(oldFile + " rm failed !").Color("r").Println("Upgrade")
			}
			time.Sleep(1 * time.Second)
			if err := gs.Str(file).Mv(gs.Str(oldFile).AbsPath().Str()); err != nil {
				gs.Str(file + " -> " + oldFile + " mv failed !").Color("r").Println("Upgrade")
			}
			args := []string{gs.Str(oldFile).AbsPath().Str()}
			gs.List[string](args).Join(" ").Println("Upgrade:run")
			Daemon(args)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		} else {
			gs.Str("not upgrade self:" + testStr).Color("r").Println("Upgrade")

		}
	} else {
		gs.Str(fpath.Str() + " not found !!").Println("Upgrade")
	}
}

func BuildGoDev() {
	gs.Str(`mkdir -p  /tmp/repo_update/GoR ; cd /tmp/repo_update/GoR && wget -c 'https://go.dev/dl/go1.22.3.linux-amd64.tar.gz' && tar -zxf go1.22.3.linux-amd64.tar.gz ;cp -a /tmp/repo_update/GoR/go /usr/share/go ; export PATH="$PATH:/usr/share/go/bin"; go version;`).Exec()
}

func ErrLog(err error) {
	fp, err_ := os.OpenFile("/tmp/upgrade.err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err_ == nil {
		defer fp.Close()
	}
	fp.Write([]byte(err.Error() + "\n"))

}

func InfoLog(record string) {
	fp, err_ := os.OpenFile("/tmp/upgrade.err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err_ == nil {
		defer fp.Close()
	}
	fp.Write([]byte(record + "\n"))

}

func Daemon(args []string) {
	// defer os.Remove(LOG_FILE)
	LOG_FILE := "/tmp/node-x.log"
	// gs.Str("ps aux | grep linux")
	ppid := os.Getppid()
	gs.Str(fmt.Sprint(ppid)).Color("g").Println("Daemon PPID")
	if os.Getppid() != 0 {
		createLogFile := func(fileName string) (fd *os.File, err error) {
			dir := path.Dir(fileName)
			if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
				if err = os.MkdirAll(dir, 0755); err != nil {
					return
				}
			}

			if fd, err = os.Create(fileName); err != nil {
				ErrLog(err)
				return
			}
			return
		}
		if LOG_FILE != "" {
			logFd, err := createLogFile(LOG_FILE)
			if err != nil {
				ErrLog(err)
				// InfoLog
				logFd, err = createLogFile(LOG_FILE + ".bak.log")
				if err != nil {
					ErrLog(err)
					return
				}
				// return
			}
			defer logFd.Close()
			InfoLog("ready to start")
			cmdName := args[0]
			cmdRun, _ := exec.LookPath(cmdName)
			InfoLog("found cmd  to start:" + cmdRun)
			newProc, err := os.StartProcess(cmdRun, args, &os.ProcAttr{
				Files: []*os.File{logFd, logFd, logFd},
			})

			InfoLog("create  to start")
			if err != nil {
				ErrLog(err)
				// log.Fatal("daemon error:", err)
				return
			}
			log.Printf("Start-Deamon: run in daemon success, pid: %v\nlog : %s", newProc.Pid, LOG_FILE)

		} else {
			cmdName := args[0]
			cmdRun, _ := exec.LookPath(cmdName)
			newProc, err := os.StartProcess(cmdRun, args, &os.ProcAttr{
				Files: []*os.File{nil, nil, nil},
			})

			if err != nil {
				ErrLog(err)
				// log.Fatal("daemon error:", err)
				return
			}
			log.Printf("Start-Deamon: run in daemon success, pid: %v\n", newProc.Pid)

		}

		return
	}
}

func DaemonLog(args []string, logpath string) {
	// defer os.Remove(LOG_FILE)
	LOG_FILE := logpath
	// gs.Str("ps aux | grep linux")
	ppid := os.Getppid()
	gs.Str(fmt.Sprint(ppid)).Color("g").Println("Daemon PPID")
	if os.Getppid() != 0 {
		createLogFile := func(fileName string) (fd *os.File, err error) {
			dir := path.Dir(fileName)
			if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
				if err = os.MkdirAll(dir, 0755); err != nil {
					return
				}
			}

			if fd, err = os.Create(fileName); err != nil {
				ErrLog(err)
				return
			}
			return
		}
		if LOG_FILE != "" {
			logFd, err := createLogFile(LOG_FILE)
			if err != nil {
				ErrLog(err)
				// InfoLog
				logFd, err = createLogFile(LOG_FILE + ".bak.log")
				if err != nil {
					ErrLog(err)
					return
				}
				// return
			}
			defer logFd.Close()
			InfoLog("ready to start")
			cmdName := args[0]
			cmdRun, _ := exec.LookPath(cmdName)
			InfoLog("found cmd  to start:" + cmdRun)
			newProc, err := os.StartProcess(cmdRun, args, &os.ProcAttr{
				Files: []*os.File{logFd, logFd, logFd},
			})

			InfoLog("create  to start")
			if err != nil {
				ErrLog(err)
				// log.Fatal("daemon error:", err)
				return
			}
			log.Printf("Start-Deamon: run in daemon success, pid: %v\nlog : %s", newProc.Pid, LOG_FILE)

		} else {
			cmdName := args[0]
			cmdRun, _ := exec.LookPath(cmdName)
			newProc, err := os.StartProcess(cmdRun, args, &os.ProcAttr{
				Files: []*os.File{nil, nil, nil},
			})

			if err != nil {
				ErrLog(err)
				// log.Fatal("daemon error:", err)
				return
			}
			log.Printf("Start-Deamon: run in daemon success, pid: %v\n", newProc.Pid)

		}

		return
	}
}
