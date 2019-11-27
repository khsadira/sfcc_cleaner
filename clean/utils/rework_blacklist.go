package utils

import "github.com/khsadira/cleaner/blacklist"

func removeIndexStruct(s []DataStruct, index int) []DataStruct {
	return append(s[:index], s[index+1:]...)

}

func ReworkBlacklist(hosts *Global, datas []blacklist.Stock) {
	for _, bdata := range datas {
		for i, host := range hosts.Hosts {
			for j, site := range host.Sites {
				for k, opt := range site.Opts {
					for l, data := range opt.Data {
						if data.ID == bdata.Id {
							hosts.Hosts[i].Sites[j].Opts[k].Data = removeIndexStruct(hosts.Hosts[i].Sites[j].Opts[k].Data, l)
						}
					}
				}
			}
		}
	}
}
