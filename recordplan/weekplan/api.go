package weekplan

import "errors"

var (
	// ErrNotFound 表示没有相关的数据
	ErrNotFound = errors.New("data not found")
)

// Run 启动后台检查
func Run(checkInterval, concurrency int) error {
	return _checker.init(checkInterval, concurrency)
}

// Reload 加载内存，在添加和修改数据库后调用
func Reload(id string) {
	_checker.Add(1)
	go _checker.loadRoutine(id)
}

// Remove 移除内存，在删除数据库后调用
func Remove(id string) {
	_checker.remove(id)
}

// BatcRemove 批量移除内存，在删除数据库后调用
func BatcRemove(ids []string) {
	_checker.batchRemove(ids)
}

// IsRecording 返回指定 id 的录像是否在录像时间段。
// 再查询出关联的 stream ，就可以知道 stream 是否需要录像。
// 如果 id 不存在，返回 ErrNotFound
func IsRecording(id string) (bool, error) {
	_checker.RLock()
	plan := _checker.weekplan[id]
	_checker.RUnlock()
	if plan == nil {
		return false, ErrNotFound
	}
	return plan.IsRecording, nil
}