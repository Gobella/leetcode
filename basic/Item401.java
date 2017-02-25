package leetcode;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

//A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).
//Each LED represents a zero or one, with the least significant bit on the right.
//For example, the above binary watch reads "3:25".
//
//Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.
//
//Example:
//
//Input: n = 1
//Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]ßß 

public class Item401 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Item401 i=new Item401();
		List<String> ls=i.readBinaryWatch(0);
		print(ls);
	}
	
	public List<String> readBinaryWatch(int num) {
		List<String> result=new ArrayList<String>();
		backtrack(result,new ArrayList<Integer>(),num,0);
        return result;
    }
	
	public void backtrack(List<String> result,List<Integer> pos,int count,int start){
		if(count==0){
			String str=getTime(pos);
			if(str!="")
				result.add(str);
			return;
		}	
		for(int i=start;i<10&&count>0;i++){
			pos.add(i);
			backtrack(result,pos,count-1,i+1);
			pos.remove(pos.size()-1);
		}
	}
	public String getTime(List<Integer> pos){
		int hour=0,min=0;
		for(int i=0;i<pos.size();i++){
			int tp=pos.get(i);
			if(tp<4){
				hour+=Math.pow(2,tp);
			}else{
				min+=Math.pow(2,tp-4);
			}
		}
		if(hour>11||min>59){
			return "";
		}
		if (min<10)
			return (String) (hour+":0"+min);
		else
			return (String)(hour+":"+min);
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
