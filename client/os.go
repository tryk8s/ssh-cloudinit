package client

import "fmt"

func GetUbuntuCmds(conf *Config) []string {
	return []string{
		"sudo apt-get update",
		"sudo apt-get -y install cloud-init",
		fmt.Sprintf("sudo echo 'cloud_init_modules: [write-files, update_etc_hosts, users-groups]\nusers: []\ndatasource_list: [NoCloud]\ndatasource: \n  NoCloud: \n    seedfrom: %s' > /etc/cloud/cloud.cfg.d/95_nocloud.cfg", conf.Server),
		"sudo cloud-init init",
		"sudo cloud-init modules -m config",
		"sudo cloud-init modules -m final",
	}
}
