package index

type TermIndexVisitor interface {
	StartRound()

	// Visit 匹配到一个索引条目，由访问者确定是否是可接受的匹配项。
	//索引条目{@link TermIndexEntry#getItems()}一定包含1个或多个索引对象{@link TermIndexItem}
	//@param entry 当前索引条目。
	//@param pos 当前匹配位置
	//@return 是可接受的匹配项时返回true，否则返回false。对于可接受的匹配项会调用{@link #endVisit(TermIndexEntry)}，否则不会调用。
	Visit(entry TermIndexEntry, text string, pos int) bool

	// PositionAfterAcceptItem 如果visit时接受了某个索引项，该方法会返回接受索引项之后当前匹配的指针
	PositionAfterAcceptItem() int

	//结束索引条目的访问。
	//@param entry 当前索引条目。
	//@param pos 当前匹配位置
	endVisit(entry TermIndexEntry, text string, pos int)

	//结束一轮词条匹配。
	endRound()
}