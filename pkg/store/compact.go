package store

func Compact() {
	kv := Kvstore
	// 切换到合并文件
	kv.CurPage = COMPACT_FILE_PAGE
	kv.SwitchFile()
}
