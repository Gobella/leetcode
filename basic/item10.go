package basic

import "fmt"

/*

Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
*/
//.* state 3
//.^* state 1
//$* state 2
//$$ state 0
type status struct {
	state int
	char  string
}
func isMatch(s string, p string) bool {
	matchbytes:=[]byte(s)
	matchRule:=[]byte(p)

	maprule:=make(map[int]*status)

	step:=0
	for i:=0;i<len(matchRule);i++{
		temp:=string(matchRule[i])
		next:=""
		step++
		if i+1<len(matchRule){
			next=string(matchRule[i+1])
		}
		if temp=="."&&next!="*"{
			maprule[step]=&status{
				state:1,
			}
			continue
		}

		if temp=="."&&next=="*"{
			i++
			maprule[step]=&status{
				state:3,
			}
			continue
		}
		if next=="*"{
			i++
			maprule[step]=&status{
				state:2,
				char:temp,
			}
			continue
		}
		maprule[step]=&status{
			state:0,
			char:temp,
		}
	}
	for i:=1;i<step+1;i++{
		fmt.Println(i,maprule[i].state)
	}

	return match(1,maprule,matchbytes,step)
}

func match(stepMatch int,maprule map[int]*status,matchbytes []byte,allstep int) bool{
	for k:=0;k<len(matchbytes);k++{

		vstr:=string(matchbytes[k])
		steprule,ok:=maprule[stepMatch]
		if !ok{
			return false
		}
		if maprule[stepMatch].state==3{
			_,ok:=maprule[stepMatch+1]
			if !ok{
				return true
			}
			blen:=len(matchbytes)

			if k<blen&&match(stepMatch+1,maprule,matchbytes[k:blen],allstep){
				return true
			}
		}
		if steprule.state==0{
			fmt.Println("rule",0,steprule.char,vstr)
			stepMatch++
			if vstr!=steprule.char{
				return false
			}
			continue
		}
		if steprule.state==2{
			fmt.Println("rule",2,vstr,steprule.char)
			ti:=1
			for ;ti+stepMatch<allstep+1&&maprule[stepMatch+ti].state==2;ti++{}
			if ti+stepMatch-1!=allstep&&match(stepMatch+ti,maprule,matchbytes[k:len(matchbytes)],allstep){
				return true
			}

			if vstr!=steprule.char{
				k--
				fmt.Println("rule",2,vstr,k)
				stepMatch++
			}
		}
		if steprule.state==1{
			stepMatch++
		}
		fmt.Println("stepmatch",stepMatch,maprule[stepMatch])
	}
	steprule,ok:=maprule[stepMatch]
	if ok&&(steprule.state==2||steprule.state==3)&&stepMatch==allstep{
		return true
	}
	if ok&&stepMatch<allstep+1{
		//fmt.Println(ok,(steprule.state!=2||steprule.state!=3),stepMatch<step+1,steprule.state)
		return false
	}
	return true
}


func isMatch2(s string, p string) bool{
	sb:=[]byte(s)
	pb:=[]byte(p)
	return matchHelper(sb,pb,0,0)
}
func matchHelper(s []byte, p []byte,sindex int,pindex int) bool{
	if pindex==len(p){
		return sindex==len(s)
	}
	sb:=[]byte(s)
	pb:=[]byte(p)
	if pindex==len(p)-1||pb[pindex+1]!=byte('*'){
		if sindex==len(s)||(s[sindex]!=p[pindex]&&s[sindex]!=byte('.')){
			return false
		}else{
			return matchHelper(sb,pb,sindex+1,pindex+1)
		}
	}
	for ;sindex!=len(s)&&(p[pindex]==byte('.')||p[pindex]==s[sindex]);{
		if matchHelper(s,p,sindex,pindex+2){
			return true
		}
		sindex++
	}
	return matchHelper(s,p,sindex,pindex+2)

}