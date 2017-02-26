package leetcode;

public class Item234 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		ListNode l=new ListNode(1);
		ListNode head=l;
		l.next=new ListNode(2);
//		l=l.next;
//		l.next=new ListNode(2);
//		l=l.next;
//		l.next=new ListNode(1);
//		l=l.next;
//		l.next=new ListNode(4);
//		l=l.next;
//		l.next=new ListNode(1);
		
//		for(int i=0;i<5;i++){
//			 l.next=new ListNode(i);
//			 l=l.next;			
//		}
//		for(int i=9;i>-1;i--){ 
//			 l.next=new ListNode(i);
//			 l=l.next;			
//		}
//		l.next=new ListNode(1);
		System.out.println(isPalindrome(head));
		

	}
	 public static boolean isPalindrome(ListNode head) {
		 if(head==null){
			 return true;
		 }
		 int middle=0,end=0;
		 ListNode cusor=head;
		 ListNode mid=head;
		 while(cusor!=null){
			 int temp=(0+end)/2;
			 for(int i=middle;i<temp;i++){
				 mid=mid.next;
				 middle++;
			 }	
			 cusor=cusor.next;
			 end++;
		 }
		 if(end%2==0){			 
			 mid=mid.next;
		 }
		 ListNode mh=reverse(mid);	
		 while(mh!=null){
			 if(mh.val!=head.val){
				 return false;
			 }
			 mh=mh.next;
			 head=head.next;
		 }
	     return true;   
	 }
	 
	 public static ListNode reverse(ListNode head){
		 
		 ListNode temp=null;
		 ListNode temp1=null;
		 if(head.next!=null){
			 temp=head.next.next;
			 head.next.next=head;
			 temp1=head.next;
			 head.next=null;			 
		 }
		 if(temp==null&&temp1!=null){
			 head=temp1;
		 }
		 while(temp!=null){
			 head=temp;
			 temp=temp.next;
			 head.next=temp1;
			 temp1=head;
		 }
		 return head;
	 }
//	 System.out.println("after:"+mid.val+"-"+mid.next.val);
//	 while(mh!=null){
//		 System.out.println(mh.val);
//		 mh=mh.next;
//	 }

}
