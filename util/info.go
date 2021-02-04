package util

import (
	"os/user"

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
	if err != nil {
		panic(err)
	}
	token := windows.Token(0)
	member, err := token.IsMember(sid)
	if err != nil {
		panic(err)
	}

	return member
}

//IsUserAdmin checks if the current user is an administrator.
//If the process is impersonating a user, it will return that value.
func IsUserAdmin() bool {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	ids, err := u.GroupIds()
	for _, id := range ids {
		if id == "S-1-5-32-544" {
			return true
		}
	}
	return false
}
