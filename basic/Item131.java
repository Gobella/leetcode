package leetcode;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return all possible palindrome partitioning of s.
//
//For example, given s = "aab",
//Return
//
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
		
public class Item131 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String s="aab";
		List<List<String>> list=partition(s);
		System.out.println("len:"+list.size());
		Iterator<List<String>> itr =list.iterator();
		while(itr.hasNext()){
			List<String> tp=itr.next();
			print(tp);
		}
	}
	
	 public static List<List<String>> partition(String s) {
		 if(s.length()==0){
			 return null;
		 }
		 List<List<String>> result=new ArrayList<List<String>>();
		 findSubPlindromeString(s,new ArrayList<String>(),result,0);
		 return result;	        
	 }
	 //len is attribute of string s?
	 public static void findSubPlindromeString(String s,List<String> resultItem,List<List<String>> result,int start){
		 if(start==s.length()){
			 result.add(new ArrayList<>(resultItem));
			 return;
		 }		 
		 for(int i=start;i<s.length();i++){		 
				if(isPlindrome(s,start,i)){
					resultItem.add(s.substring(start, i+1));
					findSubPlindromeString(s,resultItem,result,i+1);
					resultItem.remove(resultItem.size()-1);
				}		
		 }
	 }
//	 保证s变量非空
	 public static boolean isPlindrome(String s,int start,int end){
		 char[] c=new char[3];
		 int   cd=c.length;
		 for(;start<=end;){
			 if(s.charAt(start++)!=s.charAt(end--)){
				 return false;
			 }
		 }
		 return true;
	 }
	 public static void print(List<String> list){
			Iterator<String> itr =list.iterator();
			while(itr.hasNext()){
			    String tp=itr.next();
				System.out.print(tp+"  ");
			}
			System.out.println("  ");
		}
}
