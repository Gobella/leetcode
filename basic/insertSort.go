package basic

func  InsertSort(a []int,n int) []int{
	if n==0 || a==nil{
		return nil
	}else if n==1{
		return a
	}

	for i:=1;i<n;i++{
		temp:=a[i]
		j:=i-1
		for ;j>=0&&a[j]>temp;j--{
			a[j+1]=a[j]
		}
		a[j+1]=temp

	}
	return a
}

func SelectSort(a []int,n int) []int{
	if n==0 || a==nil{
		return nil
	}else if n==1{
		return a
	}

	for i:=0;i<n;i++{
		temp:=i
		j:=i+1
		for ;j<n;j++{
			if a[temp]>a[j]{
				temp=j
			}
		}
		tv:=a[i]
		a[i]=a[temp]
		a[temp]=tv
	}
	return a
}
