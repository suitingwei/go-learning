package timeAttacks

import (
	"github.com/labstack/gommon/random"
	"runtime"
	"testing"
)

var pwd string

func TestGuessPwd(t *testing.T) {
	runtime.GOMAXPROCS(1)

	for i := 1; i < 1000; i++ {
		pwd = random.String(uint8(i), random.Lowercase)

		possiblePwd := GuessWord(printNothing)

		if possiblePwd == pwd {
			//最终结果
			t.Logf("最终猜测密码=%s,实际密码=%s\n", possiblePwd, pwd)
		} else {
			t.Fatalf("最终猜测密码=%s,实际密码=%s\n", possiblePwd, pwd)
		}
	}
}

func TestSingleGuessPwd(t *testing.T) {
	runtime.GOMAXPROCS(1)

	pwd = "ywcypgdjteqgseuafyvxqldmxjgfjxbohciuwljg"

	possiblePwd := GuessWord(printNothing)

	if possiblePwd == pwd {
		//最终结果
		t.Logf("最终猜测密码=%s,实际密码=%s\n", possiblePwd, pwd)
	} else {
		t.Fatalf("最终猜测密码=%s,实际密码=%s\n", possiblePwd, pwd)
	}
}
