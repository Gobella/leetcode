package leetcode;

public class Item459 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		System.out.println(repeatedSubstringPattern("abab"));
	}
	
	 public static boolean repeatedSubstringPattern(String str) {
		 if(str.length()==0){
			 return true;
		 }
	 	int bound=str.length()/2;
	 	for(int i=1;i<bound+1;i++){
	 		if(str.length()%i==0&&isdup(i,str))
	 			return true;
	 	}
        return false;
	 }
//	 递归
	 public static boolean nextMatch(int start,int count,String s){
		 if(start==s.length()){
			 return true;
		 }
		 for(int i=0;i<count;i++){
			 if(s.charAt(i)!=s.charAt(start+i)){
				 return false;
			 }
		 }
		 return nextMatch(start+count,count,s);
	 }
	 public static boolean isdup(int count,String s){
		 String str=s.substring(0,count);
		 for(int i=1;i<s.length()/count;i++){
			 if(!str.equals(s.substring(i*count,i*count+count)))
				 return false;
		 }
		 return true;
	 }

}
