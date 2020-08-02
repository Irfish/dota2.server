package gin

type Gin struct {
	Address string
}
func (p *Gin)Addr() string {
	return p.Address
}
func Run() {
	//server := new(Gin)
	//g.Run(server)
}
