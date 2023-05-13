package pointer_list

/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环

注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。
*/

/*
先判断是否有环
有环的话，快慢指针相遇的点一定在环内
然后让一个指针从头开始，一个指针从相遇点开始，两个指针都是一次走一步，
那么两个指针一定会在入环点相遇
*/

func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			s = head
			for f != s {
				f = f.Next
				s = s.Next
			}
			return f
		}
	}
	return nil
}
