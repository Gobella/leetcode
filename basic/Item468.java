package leetcode;

public class Item468 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}
	
	public String validIPAddress(String IP) {
		if(IP.contains(":")){
			return isIPv6(IP);
		}else if(IP.contains(".")){
			return isIPv4(IP);
		}
		return "Neither";
		
        
    }
	public String isIPv4(String ip){
		String neither="Neither";
		String IPv4="IPv4";
		String[] strs=ip.split(".");
		if(strs.length!=4)
			return neither;
		for(int i=0;i<4;i++){
			if(strs[i].length()>1&&strs[i].charAt(0)=='0')
				return neither;
			int t=Integer.parseInt(strs[i]);
			if(t>255)
				return neither;
		}
		return "";
	}
	public String isIPv6(String ip){
		String IPv6="IPv6";
		return "";
	}

}
