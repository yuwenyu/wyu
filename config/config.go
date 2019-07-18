package config

import "time"

const(
	ConfSession 	string = "sessions"
	ConfCommons 	string = "commons"
	ConfTemplates 	string = "templates"

	MSessionKeySID 			string = "wyu_ssk"
	MSessionKeyPrefix 		string = "wyu_ssk_prefix"
	MSessionRedisNetwork	string = "tcp"
	MSessionRedisKeyPairs	string = "wyu_secret"
	MSessionStorePath		string = "/"

	MSessionRedisPool	int = 10
	MSessionStoreMaxAge	time.Duration = 60
)

var(
	MapConfLists map[string]map[int]string = map[string]map[int]string{
		ConfSession:	map[int]string{
			0:"session",
		},
		ConfCommons:	map[int]string{
			0:"common_cfg",
			1:"common_log",
			2:"common_server",
		},
		ConfTemplates:	map[int]string{
			0:"template_statics",
			1:"template_root",
			2:"template_config",
		},
	}
	MapConfParam map[string]map[int]string = map[string]map[int]string{
		MapConfLists[ConfSession][0]:	map[int]string{
			0:"key_sid",
			1:"key_prefix",
			2:"redis_pool",
			3:"redis_network",
			4:"redis_address",
			5:"redis_password",
			6:"redis_keypairs",
			7:"store_max_age",
			8:"store_path",
		},
		MapConfLists[ConfCommons][0]:	map[int]string{
			0:"template_status",
			1:"template_static_status",
		},
		MapConfLists[ConfCommons][1]:	map[int]string{
			0:"log_status",
			1:"log_root",
			2:"log_fn_prefix",
		},
		MapConfLists[ConfCommons][2]:	map[int]string{
			0:"port",
		},
		MapConfLists[ConfTemplates][0]:	map[int]string{
			0:"static",
			1:"static_file",
		},
		MapConfLists[ConfTemplates][1]:	map[int]string{
			0:"directory",
			1:"directory_view",
			2:"resources",
		},
		MapConfLists[ConfTemplates][2]:	map[int]string{
			0:"static_url",
			1:"static_url_version",
		},
	}
)
