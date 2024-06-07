package main

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head==nil || head.Next==nil{
			return head		
	}

	Length:=1
	findSize:=head

	for findSize.Next!=nil{
		Length++
		findSize=findSize.Next
	}

	k=k%Length
	if k==0 {
		return head
	}

	reversed:=reverse(head)
	size:=1
	temp:=reversed

	for size!=k{
		temp=temp.Next
		size++
	}

	if temp==nil {
		return reversed
	}

	secondHalf:=temp.Next
	temp.Next=nil

	answer1:=reverse(reversed)
	answer2:=reverse(secondHalf)

	temp2:=answer1
	for temp2.Next!=nil{
		temp2=temp2.Next
	}
	temp2.Next=answer2

	return answer1
}

func reverse(head *ListNode) *ListNode  {
	current:=head
	var next *ListNode
	var prev *ListNode = nil

	for current!=nil{
		next=current.Next
		current.Next=prev
		prev=current
		current=next
	}
	return prev
}