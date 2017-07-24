package client

import (
	"fmt"
	"path"
	"strings"
)

const CloudDataPath = "/tmp/cloud/"

func GetCentOSCmds(conf *Config) []string {
	s := conf.Server
	extra := ""
	if conf.UserData != "" {
		s = CloudDataPath
	} else if !strings.HasSuffix(s, "/") {
		s = s + "/"
	}
	cmds := []string{
		"sudo yum update",
		"sudo yum -y install cloud-init curl wget",
		fmt.Sprintf("sudo cloud-init %s init -l", extra),
		fmt.Sprintf("sudo cloud-init %s modules -m config", extra),
		fmt.Sprintf("sudo cloud-init %s modules -m final", extra),
	}
	if conf.UseCloudDataSource == false {
		//cmds = append([]string{fmt.Sprintf("sudo echo 'apt_preserve_sources_list: true\ncloud_init_modules: [write-files, update_etc_hosts, users-groups]\ncloud_final_modules: [scripts-vendor, scripts-per-once, scripts-per-boot, scripts-per-instance, scripts-user]\nusers: []\ndatasource_list: [NoCloud]\ndatasource: \n  NoCloud: \n    seedfrom: %s' > /etc/cloud/cloud.cfg.d/95_nocloud.cfg", s)}, cmds...)
	}
	if conf.DontCleanCloudInitStatus == false {
		cmds = append([]string{"sudo rm -rf /var/lib/cloud/instance/*"}, cmds...)
	}
	return cmds
}
func GetUbuntuCmds(conf *Config) []string {
	s := conf.Server
	extra := ""
	if conf.UserData != "" {
		s = CloudDataPath
	} else if !strings.HasSuffix(s, "/") {
		s = s + "/"
	}
	cmds := []string{
		"sudo apt-get update",
		"sudo apt-get -y install cloud-init curl wget",
	}

	if conf.UseCloudDataSource == false {
		cmds = append(cmds, fmt.Sprintf("sudo sh -c \"echo 'apt_preserve_sources_list: true\ncloud_init_modules: [write-files, update_etc_hosts, users-groups]\ncloud_final_modules: [scripts-vendor, scripts-per-once, scripts-per-boot, scripts-per-instance, scripts-user]\nusers: []\ndatasource_list: [NoCloud]\ndatasource: \n  NoCloud: \n    seedfrom: %s' > /etc/cloud/cloud.cfg.d/96_nocloud.cfg\"", s))
	}

	cmds = append(cmds,
		fmt.Sprintf("sudo cloud-init %s init -l", extra),
		fmt.Sprintf("sudo cloud-init %s modules -m config", extra),
		fmt.Sprintf("sudo cloud-init %s modules -m final", extra),
	)

	if conf.DontCleanCloudInitStatus == false {
		cmds = append([]string{"sudo rm -rf /var/lib/cloud/instance/*"}, cmds...)
	}
	return cmds
}

func GetCoreOSCmds(conf *Config) []string {
	filePath := path.Join(CloudDataPath, "user-data")
	if conf.UserData != "" {
		return []string{
			fmt.Sprintf("sudo coreos-cloudinit --from-file=%s", filePath),
		}
	}

	return []string{
		fmt.Sprintf("sudo coreos-cloudinit --from-url=%s", conf.Server),
	}
}
