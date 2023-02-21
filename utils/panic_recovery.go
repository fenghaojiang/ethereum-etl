package utils

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/fenghaojiang/ethereum-etl/common/log"
	"go.uber.org/zap"
)

const frameToSkip = 3

func ExecuteWithRecover(action string, fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("execute action panic", zap.String("action name", action), zap.Any("stack info", stackInfo(frameToSkip)))
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return fn()
}

func stackInfo(skip int) string {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
