package main

import(
	"os/exec"
	"path"
	"fmt"
	"io/ioutil"
	"strconv"
	"syscall"
	"os"
)

const cgroupMemoryHierarchMount = "/sys/fs/cgroup/memory"

func main(){
	fmt.Println(os.Args[0])
	if os.Args[0] == "/proc/self/exe" {
		//容器进程
		fmt.Printf("current pid %d",syscall.Getpid())
		fmt.Printf("unormal")
		fmt.Println()
		cmd := exec.Command("sh","-c",`stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr= &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err:=cmd.Run();err != nil{
			fmt.Println(err)
		}
		os.Exit(1)
	}
	fmt.Printf("keep")
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 5001,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 5001,
				HostID:      0,
				Size:        1,
			},
		},
	}  
	// 在我的版本  go1.15.6 不能有这行，，1.7 可以
		// cmd.SysProcAttr.Credential = &syscall.Credential{Uid:uint32(1),Gid:uint32(1)}
	cmd.Stdin = os.Stdin 
	cmd.Stdout= os.Stdout 
	cmd.Stderr = os . Stderr

	if err := cmd.Run();err != nil{
		fmt.Println("ERROR",err)
		os.Exit(1)
	}else{
		fmt.Printf("%v",cmd.Process.Pid)
		//在系统默认创建挂载了memory subsystem 的hierarchy 上创建cgroup
		os.Mkdir(path.Join(cgroupMemoryHierarchMount,"testmemorylimit"),0755)
		//将容器进程加入这个cgroup
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchMount,"testmemorylimit","tasks"),[]byte(strconv.Itoa(cmd.Process.Pid)),0644)
		//限制cgroup进程使用
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchMount,"testmemorylimit","memory.limit_in_bytes"),[]byte("100m"),0644)

	}
	cmd.Process.Wait()
}
