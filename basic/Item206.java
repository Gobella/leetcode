package leetcode;

public class Item206 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		ListNode l=new ListNode(1);
		ListNode head=l;
		l.next=new ListNode(2);
		l=l.next;
		l.next=new ListNode(3);
		ListNode ls=reverseList(head);
		
		while(ls!=null){
			System.out.println(ls.val);
			ls=ls.next;
		}
	}
	public static ListNode reverseList(ListNode head) {
		if(head==null||head.next==null){
			return head;
		}
		ListNode temp=head.next.next;
		ListNode cusor=head;
		head.next.next=head;
		cusor=head.next;
		head.next=null;	
		head=cusor;
		while(temp!=null){
			head=temp;
			temp=temp.next;
			head.next=cusor;
			cusor=head;
		}		
		return head;       
    }

}
