package p1 // 包名和文件夹名不需要相同，但是建议用相同的名称

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func P1_test1(name string) string {
	sb := strings.Builder{}
	sb.WriteString("hello [")
	sb.WriteString(name)
	sb.WriteString("] from p1 test1")

	sl := make([]string, 4)
	sl = append(sl, sb.String())
	sl = append(sl, "::join1 [", name, "]")
	rt := strings.Join(sl, "")
	fmt.Println(rt)

	fmt.Println(sb.String())
	sl2 := []string{sb.String(), "::join2 ["}
	sl2 = append(sl2, name, "]")
	rt = strings.Join(sl2, "")

	fmt.Println(cap(sl2))

	return rt
}

type P1_Struct1 struct {
	Public_var1  string
	private_var2 string
}

var P1_Struct1Instance = P1_Struct1{
	Public_var1:  "var1 is public",
	private_var2: "var2 is private",
}

func P1_TestStruct(instance P1_Struct1) {
	// 大写开头的struct成员变量在包内部和包外部都是可见的
	fmt.Printf("public var1 in P1_Struct1 is: %v\n", instance.Public_var1)
	// 小写开头的struct成员变量在包内部是可见的，但是在包外部不可见
	instance.private_var2 += " but changed"
	fmt.Printf("private var2 in P1_Struct1 is: %v\n", instance.private_var2)
}

func P1_test2() {
	var a bool = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

func P1_TestRegex() {
	pt1, str1 := "(?i)(?s).*a coder.*", "I'm a coder"
	ok, _ := regexp.MatchString(pt1, str1)
	fmt.Printf("%q match %q ? %v\n", str1, pt1, ok)

	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("meet error in recover:", e)
		}
	}()

	pt2 := "(?s)(.*)@hpe.*"
	str2 := "lei.wang29@hpe.com"
	reg1 := regexp.MustCompile(pt2)
	fmt.Printf("%q match %q ? %v\n", str2, pt2, reg1.MatchString(str2))
	matchs := reg1.FindStringSubmatch(str2)
	if matchs != nil && len(matchs) > 0 {
		log.Println("find subMatchs:", matchs)
		for i := 0; i < len(matchs); i++ {
			log.Printf("\tgroup [%v] is: %s\n", i, matchs[i])
		}
	}
}
