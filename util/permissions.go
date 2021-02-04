package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/sys/windows"
)

//RunningAsAdmin returns whether the current process has administrative privileges.
func RunningAsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	Handle(err)
	token := windows.Token(0)
	member, err := token.IsMember(sid)
	Handle(err)

	return member
}

//IsUserAdmin checks if the current user is an administrator.
//If the process is impersonating a user, it will return that value.
func IsUserAdmin() bool {
	u, err := user.Current()
	Handle(err)
	ids, err := u.GroupIds()
	Handle(err)
	for _, id := range ids {
		if id == "S-1-5-32-544" {
			return true
		}
	}
	return false
}

//IsWritable return whether a path or a file is writable.
func IsWritable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	var fn string
	var f *os.File
	if info.IsDir() {
		fmt.Println("Checking path")
		fn = filepath.Join(path, "check")
		f, err = os.Create(fn)
		f.Close()
		os.Remove(fn)
	} else {
		fmt.Println("Checking file")
		fn = path
		f, err = os.Open(fn)
		f.Close()
	}

	if err != nil {
		return false
	}

	return true
}
