package main

import(
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main(){
	cmd :=exec.Command("sh")
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
		log.Fatal(err)
	}
	os.Exit(-1)

}