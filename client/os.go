package client

import "fmt"

func GetUbuntuCmds(conf *Config) []string {
	return []string{
		"apt-get update",
		"apt-get -y install cloud-init",
		fmt.Sprintf("echo 'cloud_init_modules: [write-files, update_etc_hosts, users-groups]\nusers: []\ndatasource_list: [NoCloud]\ndatasource: \n  NoCloud: \n    seedfrom: %s' > /etc/cloud/cloud.cfg.d/95_nocloud.cfg", conf.Server),
		"cloud-init init",
		"cloud-init modules -m config",
		"cloud-init modules -m final",
	}
}
