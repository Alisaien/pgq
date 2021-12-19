package pgq

type SelectCmd struct {
	Cmd
	fromCls   *buffer
	selectCls *buffer
	whereCls  *buffer

	sql *buffer
}

func Select() *SelectCmd {
	return &SelectCmd{
		fromCls:   pool.Get().(*buffer),
		selectCls: pool.Get().(*buffer),
		whereCls:  pool.Get().(*buffer),
	}
}

func (c *SelectCmd) Distinct(expr string) {

}

func (c *SelectCmd) Select(expr string) *SelectCmd {
	if c.selectCls.len() > 0 {
		c.selectCls.writeSQL(",")
	}
	c.selectCls.writeSQL(expr)

	return c
}

func (c *SelectCmd) From(fromItem string) *SelectCmd {
	if c.selectCls.len() > 0 {
		c.selectCls.writeSQL(",")
	}
	c.selectCls.writeSQL(fromItem)

	return c
}

func (c *SelectCmd) Where(cond string) *SelectCmd {
	if c.whereCls.len() > 0 {
		c.whereCls.writeSQL(" AND " + cond)
	} else {
		c.whereCls.writeSQL("WHERE " + cond)
	}

	return c
}

func (c *SelectCmd) SQL() string {
	c.sql.writeString("SELECT ").write(c.selectCls)
	c.sql.writeString(" FROM ").write(c.fromCls)
	if c.whereCls.len() > 0 {
		c.sql.writeString(" WHERE ").write(c.whereCls)
	}

	return c.sql.string()
}
