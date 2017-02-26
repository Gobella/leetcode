package leetcode;

import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

public class Item445 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		ListNode l=new ListNode(7);
		ListNode head=l;
		l.next=new ListNode(2);
		l=l.next;
		l.next=new ListNode(4);
		l=l.next;
		l.next=new ListNode(3);
		
		ListNode l1=new ListNode(5);
		ListNode head1=l1;
		l1.next=new ListNode(6);
		l1=l1.next;
		l1.next=new ListNode(4);

		addTwoNumbers(head,head1);
	}
	 public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
		 Stack<Integer> s1 = new Stack<Integer>();  
		 Stack<Integer> s2 = new Stack<Integer>();  
		 Stack<Integer> res = new Stack<Integer>();  
		
		 ListNode c1=l1;
		 ListNode c2=l2;
		 
		 while(c1!=null){
			 s1.push(c1.val);
			 c1=c1.next;
		 }
		 while(c2!=null){
			 s2.push(c2.val);
			 c2=c2.next;
		 }
		 int sum=0;
		 while(!s1.isEmpty()||!s2.isEmpty()||sum!=0){
			
			 int num1=0;
			 int num2=0;
			 if(!s1.isEmpty())
				 num1=s1.pop();
			 if(!s2.isEmpty())
				 num2=s2.pop();
			 sum=sum+num1+num2;
			 res.push(sum%10);
			 sum=sum/10;
		 }
		 ListNode resln=new ListNode(res.pop());
		 ListNode head =resln;
		 while(!res.isEmpty()){
			 resln.next=new  ListNode(res.pop());
			 resln=resln.next;
		 }		 
		 return head;        
	 }

}
