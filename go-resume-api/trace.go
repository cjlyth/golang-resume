package api

func (l *Scope) trace(s string) {
	defer l.un(l.en(s))
}

func (l *Scope) en(s string) string {
	l.Debugf("%s %s(...)\n", "~~> ENTER", s)
	return s
}
func (l *Scope) un(s string) {
	l.Debugf("%40s %s(...)\n", "EXIT", s)
}
