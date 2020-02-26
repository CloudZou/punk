package main

import "time"

var toolIndexs = []*Tool{
	{
		Name:      "punk",
		Alias:     "punk",
		BuildTime: time.Date(2019, 6, 21, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/CloudZou/AngryMiao-Stellar/tool/punk",
		Summary:   "punk工具集本体",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "punk",
		Hidden:    true,
	},
	{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2019, 5, 5, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
	{
		Name:      "genbts",
		Alias:     "punk-gen-bts",
		BuildTime: time.Date(2019, 10, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/CloudZou/AngryMiao-Stellar/tool/punk-gen-bts",
		Summary:   "缓存回源逻辑代码生成器",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "punk",
	},
	{
		Name:      "genredis",
		Alias:     "punk-gen-redis",
		BuildTime: time.Date(2019, 7, 23, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/CloudZou/AngryMiao-Stellar/tool/punk-gen-redis",
		Summary:   "redis缓存代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "punk",
	},
	{
		Name:      "wire",
		Alias:     "wire",
		BuildTime: time.Date(2019, 8, 20, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/google/wire/cmd/wire",
		Platform:  []string{"darwin", "linux", "windows"},
		Hidden:    true,
	},
	{
		Name:         "genproject",
		Alias:        "punk-gen-project",
		Install:      "go get -u github.com/CloudZou/AngryMiao-Stellar/tool/punk-gen-project",
		BuildTime:    time.Date(2019, 12, 22, 0, 0, 0, 0, time.Local),
		Platform:     []string{"darwin", "linux", "windows"},
		Hidden:       true,
		Requirements: []string{"wire"},
	},
	{
		Name:      "testcli",
		Alias:     "testcli",
		BuildTime: time.Date(2019, 9, 9, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/CloudZou/AngryMiao-Stellar/tool/testcli",
		Summary:   "测试代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "punk",
	},
}
