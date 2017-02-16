package leetcode;

public class Item53 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a={1,2};
		System.out.println("result:"+maxSubArray2(a));

	}
	 /**
	 * @param nums
	 * @return
	 * 连续性！！！！！
	 */
	public static int maxSubArray(int[] nums) {
		 int sum=Integer.MIN_VALUE;
		 System.out.println("sum init:"+sum);
		 int len=nums.length;
		 for(int i=0;i<len;i++){
			 int tsum=0;
			 for(int j=i;j<len;j++){
				 System.out.println("tsum:"+tsum+"-sum:"+sum+"-nums:"+nums[i]+"-i:"+i);
				 tsum=tsum+nums[j];
				 if(tsum>sum){
					 sum=tsum;
				 }
			 }
		 }
		 return sum;	        
	 }
	public static int maxSubArray2(int[] nums) {
		 int sum=2147483647;
		 sum=-sum -1;
		 int len=nums.length;
		 int[][] data=new int[len][len];
		 for(int i=0;i<len;i++){
			 data[i][i]=nums[i];
			 if(data[i][i]>sum){
				 sum=data[i][i];
			 }
		 }
		 for(int i=1;i<len;i++){
			 for(int j=0;j<len&&j+i<len;j++){
				 System.out.println("[i]"+i+"[j]"+j);
				 data[j][j+i]=data[j][j+i-1]+nums[j+i];
				 if(data[j][j+i]>sum){
					 sum=data[j][j+i];
				 }
			 }
		 }
		 System.out.println(data[0][0]+" "+data[1][1]+" "+data[0][1]+" "+data[1][0]);
		 for(int i=0;i<len;i++){
			 for(int j=0;i<len;i++){
				 System.out.print("*data:"+data[i][j]+" i:"+i+"-j:"+j);
			 }
			 System.out.println("-----");
		 }
		 System.out.println(data.length);
		 return sum;	        
	 }
	public static int maxSubArray3(int[] nums) {
		int max = Integer.MIN_VALUE, sum = 0;
	    for (int i = 0; i < nums.length; i++) {
	        if (sum < 0) 
	            sum = nums[i];
	        else 
	            sum += nums[i];
	        if (sum > max)
	            max = sum;
	    }
	    return max;       
	 }

}
